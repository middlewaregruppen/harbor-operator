package ensure

import (
	"context"
	"fmt"
	"time"

	"github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	harborv1alpha1 "github.com/middlewaregruppen/harbor-operator/api/v1alpha1"
	"github.com/middlewaregruppen/harbor-operator/pkg/util"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const projectFinalizers = "harbor.mdlwr.com/finalizer"

// Definitions to manage status conditions
const (
	// ProjectAvailable represents the status of the Deployment reconciliation
	ProjectAvailable = "Available"
	// ProjectDegraded represents the status used when the custom resource is deleted and the finalizer operations are must to occur.
	ProjectDegraded = "Degraded"
)

type ProjectEnsurer struct {
	client    client.Client
	clientset *harbor.ClientSet
	proj      *harborv1alpha1.Project
	req       ctrl.Request
}

func NewProjectEnsurer(c client.Client, clientset *harbor.ClientSet, req ctrl.Request, proj *harborv1alpha1.Project) (*ProjectEnsurer, error) {
	return &ProjectEnsurer{
		client:    c,
		proj:      proj,
		req:       req,
		clientset: clientset,
	}, nil
}

func (e ProjectEnsurer) HandleDelete(ctx context.Context) (ctrl.Result, error) {

	// Perform finalizers before deleting resource from cluster
	if controllerutil.ContainsFinalizer(e.proj, projectFinalizers) {

		// Add Degraded status to begin the process of terminating resources
		meta.SetStatusCondition(&e.proj.Status.Conditions, metav1.Condition{Type: ProjectDegraded,
			Status: metav1.ConditionUnknown, Reason: "Finalizing",
			Message: fmt.Sprintf("Performing finalizer operations for the custom resource: %s ", e.proj.Name)})

		// Update the resource with updated status
		if err := e.client.Status().Update(ctx, e.proj); err != nil {
			return ctrl.Result{}, err
		}

		// TODO: run finalizers here. Always delete resources that belong to this CRD before proceeding further
		// Delete project from Harbor after finalizers have been carried out
		_, err := e.clientset.V2().Project.DeleteProject(ctx, &project.DeleteProjectParams{ProjectNameOrID: e.proj.Name})
		if nil != err {

			// Update the status field of the resource in case we get errors creating the project
			meta.SetStatusCondition(&e.proj.Status.Conditions, metav1.Condition{Type: ProjectAvailable,
				Status: metav1.ConditionFalse, Reason: "Reconciling",
				Message: fmt.Sprintf("Failed to delete Harbor project for the custom resource (%s): (%s)", e.proj.Name, err)})

			if err := e.client.Status().Update(ctx, e.proj); err != nil {
				return ctrl.Result{}, err
			}
			return ctrl.Result{}, err
		}

		// Get the Project resource again so that we don't encounter any "the object has been modified"-errors
		if err = e.client.Get(ctx, e.req.NamespacedName, e.proj); err != nil {
			return ctrl.Result{}, err
		}

		meta.SetStatusCondition(&e.proj.Status.Conditions, metav1.Condition{Type: ProjectDegraded,
			Status: metav1.ConditionTrue, Reason: "Finalizing",
			Message: fmt.Sprintf("Finalizer operations for custom resource %s name were successfully accomplished", e.proj.Name)})

		if err := e.client.Status().Update(ctx, e.proj); err != nil {
			return ctrl.Result{}, err
		}

		// Remove finalizers and update status field of the resource
		if ok := controllerutil.RemoveFinalizer(e.proj, projectFinalizers); !ok {
			return ctrl.Result{Requeue: true}, nil
		}
		if err := e.client.Update(ctx, e.proj); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (e ProjectEnsurer) Setup(ctx context.Context) (ctrl.Result, error) {
	// Check status field here
	// Let's just set the status as Unknown when no status are available
	if e.proj.Status.Conditions == nil || len(e.proj.Status.Conditions) == 0 {
		meta.SetStatusCondition(&e.proj.Status.Conditions, metav1.Condition{Type: ProjectAvailable, Status: metav1.ConditionUnknown, Reason: "Reconciling", Message: "Starting reconciliation"})
		if err := e.client.Status().Update(ctx, e.proj); err != nil {
			return ctrl.Result{}, err
		}

		// Get the Project resource again so that we don't encounter any "the object has been modified"-errors
		if err := e.client.Get(ctx, e.req.NamespacedName, e.proj); err != nil {
			return ctrl.Result{}, err
		}
	}

	// Add finalizers that will be handled later during delete events
	if !controllerutil.ContainsFinalizer(e.proj, projectFinalizers) {
		if ok := controllerutil.AddFinalizer(e.proj, projectFinalizers); !ok {
			return ctrl.Result{Requeue: true}, nil
		}
		if err := e.client.Update(ctx, e.proj); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (e ProjectEnsurer) HandleCreate(ctx context.Context) (ctrl.Result, error) {
	if _, err := e.Setup(ctx); err != nil {
		return ctrl.Result{}, err
	}
	isPublic := true
	if e.proj.Spec.IsPrivate {
		isPublic = false
	}
	projectReq := &models.ProjectReq{
		ProjectName: e.proj.ObjectMeta.Name,
		Metadata: &models.ProjectMetadata{
			Public: util.BoolToString(isPublic),
		},
	}
	// Check if the actual Project in Harbor exists. If not create a new one
	_, err := e.clientset.V2().Project.CreateProject(ctx, &project.CreateProjectParams{Project: projectReq})
	if err != nil {
		// Update the status field of the resource in case we get errors creating the project
		meta.SetStatusCondition(&e.proj.Status.Conditions, metav1.Condition{Type: ProjectAvailable,
			Status: metav1.ConditionFalse, Reason: "Reconciling",
			Message: fmt.Sprintf("Failed to create Harbor project for the custom resource (%s): (%s)", e.proj.Name, err)})

		if err := e.client.Status().Update(ctx, e.proj); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, err
	}
	// Project updated successfully. Requeue and check state again
	return ctrl.Result{RequeueAfter: time.Minute}, nil
}

func (e ProjectEnsurer) HandleUpdate(ctx context.Context) (ctrl.Result, error) {
	isPublic := true
	if e.proj.Spec.IsPrivate {
		isPublic = false
	}
	projectReq := &models.ProjectReq{
		ProjectName: e.proj.Name,
		Metadata: &models.ProjectMetadata{
			Public: util.BoolToString(isPublic),
		},
	}
	_, err := e.clientset.V2().Project.UpdateProject(ctx, &project.UpdateProjectParams{
		Project:         projectReq,
		ProjectNameOrID: e.proj.Name,
	})
	if err != nil {
		// Update the status field of the resource in case we get errors creating the project
		meta.SetStatusCondition(&e.proj.Status.Conditions, metav1.Condition{Type: ProjectAvailable,
			Status: metav1.ConditionFalse, Reason: "Reconciling",
			Message: fmt.Sprintf("Failed to update Harbor project for the custom resource (%s): (%s)", e.proj.Name, err)})

		if err := e.client.Status().Update(ctx, e.proj); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, err
	}
	// Project updated successfully. Requeue and check state again
	return ctrl.Result{RequeueAfter: time.Minute}, nil
}

func (e ProjectEnsurer) ProjectNotFound(ctx context.Context, key string) bool {
	_, err := e.clientset.V2().Project.HeadProject(ctx, &project.HeadProjectParams{ProjectName: key})
	if err != nil {
		if _, notFound := err.(*project.HeadProjectNotFound); notFound {
			return true
		}
		return true
	}
	return false
}
