package controllers

import (
	"context"
	"fmt"

	"github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	harborv1alpha1 "github.com/middlewaregruppen/harbor-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type Ensurer interface {
	Setup(context.Context) error
	Finalize(context.Context) error
	Ensure(context.Context) error
}

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

func (e ProjectEnsurer) Finalize(ctx context.Context) error {
	// Check if the Project is marked to be deleted, which is
	// indicated by the deletion timestamp on the resource.
	if e.proj.GetDeletionTimestamp() != nil {
		// Perform finalizers before deleting resource from cluster
		if controllerutil.ContainsFinalizer(e.proj, projectFinalizers) {

			// Add Degraded status to begin the process of terminating resources
			meta.SetStatusCondition(&e.proj.Status.Conditions, metav1.Condition{Type: typeDegraded,
				Status: metav1.ConditionUnknown, Reason: "Finalizing",
				Message: fmt.Sprintf("Performing finalizer operations for the custom resource: %s ", e.proj.Name)})

			// Update the resource with updated status
			if err := e.client.Status().Update(ctx, e.proj); err != nil {
				// l.Error(err, "failed to update Project status")
				// return ctrl.Result{}, err
				return err
			}

			// TODO: run finalizers here

			// Delete project from Harbor after finalizers have been carried out
			_, err := e.clientset.V2().Project.DeleteProject(ctx, &project.DeleteProjectParams{ProjectNameOrID: e.proj.ObjectMeta.Name})
			if nil != err {
				//l.Error(err, "failed to delete project from Harbor")
				// Update the status field of the resource in case we get errors creating the project
				meta.SetStatusCondition(&e.proj.Status.Conditions, metav1.Condition{Type: typeAvailable,
					Status: metav1.ConditionFalse, Reason: "Reconciling",
					Message: fmt.Sprintf("Failed to delete Harbor project for the custom resource (%s): (%s)", e.proj.Name, err)})

				if err := e.client.Status().Update(ctx, e.proj); err != nil {
					// l.Error(err, "failed to update Project status")
					// return ctrl.Result{}, err
					return err
				}
				//return ctrl.Result{}, err
				return err
			}

			// Get the Project resource again so that we don't encounter any "the object has been modified"-errors
			if err = e.client.Get(ctx, e.req.NamespacedName, e.proj); err != nil {
				//l.Error(err, "failed to re-fetch project")
				//return ctrl.Result{}, err
				return err
			}

			meta.SetStatusCondition(&e.proj.Status.Conditions, metav1.Condition{Type: typeDegraded,
				Status: metav1.ConditionTrue, Reason: "Finalizing",
				Message: fmt.Sprintf("Finalizer operations for custom resource %s name were successfully accomplished", e.proj.Name)})

			if err := e.client.Status().Update(ctx, e.proj); err != nil {
				// l.Error(err, "failed to update Project status")
				// return ctrl.Result{}, err
				return err
			}

			// Remove finalizers and update status field of the resource
			if ok := controllerutil.RemoveFinalizer(e.proj, projectFinalizers); !ok {
				// l.Error(err, "failed to remove finalizer for Project")
				// return ctrl.Result{Requeue: true}, nil
				return nil
			}
			if err := e.client.Update(ctx, e.proj); err != nil {
				// l.Error(err, "failed to remove finalizer for Project")
				// return ctrl.Result{}, err
				return err
			}
		}
	}
	return nil
}

func (e ProjectEnsurer) Setup(ctx context.Context) error {
	//l := log.FromContext(ctx)
	// Check status field here
	// Let's just set the status as Unknown when no status are available
	if e.proj.Status.Conditions == nil || len(e.proj.Status.Conditions) == 0 {
		meta.SetStatusCondition(&e.proj.Status.Conditions, metav1.Condition{Type: typeAvailable, Status: metav1.ConditionUnknown, Reason: "Reconciling", Message: "Starting reconciliation"})
		if err := e.client.Status().Update(ctx, e.proj); err != nil {
			// l.Error(err, "failed to update Project status")
			// return ctrl.Result{}, err
			return err
		}

		// Get the Project resource again so that we don't encounter any "the object has been modified"-errors
		if err := e.client.Get(ctx, e.req.NamespacedName, e.proj); err != nil {
			// l.Error(err, "failed to re-fetch Project")
			// return ctrl.Result{}, err
			return err
		}
	}

	// Add finalizers that will be handled later during delete events
	if !controllerutil.ContainsFinalizer(e.proj, projectFinalizers) {
		if ok := controllerutil.AddFinalizer(e.proj, projectFinalizers); !ok {
			// l.Error(err, "failed to add finalizer into the custom resource")
			// return ctrl.Result{Requeue: true}, nil
			return nil
		}
		if err := e.client.Update(ctx, e.proj); err != nil {
			// l.Error(err, "failed to update custom resource to add finalizer")
			// return ctrl.Result{}, err
			return err
		}
	}

	if !controllerutil.ContainsFinalizer(e.proj, projectFinalizers) {
		if ok := controllerutil.AddFinalizer(e.proj, projectFinalizers); !ok {
			// l.Error(err, "failed to add finalizer into the custom resource")
			// return ctrl.Result{Requeue: true}, nil
			return nil
		}
		if err := e.client.Update(ctx, e.proj); err != nil {
			// l.Error(err, "failed to update custom resource to add finalizer")
			// return ctrl.Result{}, err
			return err
		}
	}

	return nil
}

func (e ProjectEnsurer) Ensure(ctx context.Context) error {
	isPublic := true
	if e.proj.Spec.IsPrivate {
		isPublic = false
	}
	projectReq := &models.ProjectReq{
		ProjectName: e.proj.ObjectMeta.Name,
		Metadata: &models.ProjectMetadata{
			Public: boolToString(isPublic),
		},
	}
	// Check if the actual Project in Harbor exists. If not create a new one
	if e.projectIsNotFound(ctx, e.proj.Name) {
		_, err := e.clientset.V2().Project.CreateProject(ctx, &project.CreateProjectParams{Project: projectReq})
		if err != nil {
			//l.Error(err, "Failed creating")
			// Update the status field of the resource in case we get errors creating the project
			meta.SetStatusCondition(&e.proj.Status.Conditions, metav1.Condition{Type: typeAvailable,
				Status: metav1.ConditionFalse, Reason: "Reconciling",
				Message: fmt.Sprintf("Failed to create Harbor project for the custom resource (%s): (%s)", e.proj.Name, err)})

			if err := e.client.Status().Update(ctx, e.proj); err != nil {
				// l.Error(err, "failed to update Project status")
				// return ctrl.Result{}, err
				return err
			}
			return err
		}
		// Project updated successfully. Requeue and check state again
		//return ctrl.Result{RequeueAfter: time.Minute}, nil
		//return nil
	}

	// Check if the actual project in Harbor exists. If so then update the existing one
	if !e.projectIsNotFound(ctx, e.proj.ObjectMeta.Name) {
		// _, err = r.clientset.V2().Project.GetProject(ctx, &project.GetProjectParams{ProjectNameOrID: proj.Name})
		// if err != nil {
		// 	l.Error(err, "couldn't get project")
		// 	return ctrl.Result{}, err
		// }
		_, err := e.clientset.V2().Project.UpdateProject(ctx, &project.UpdateProjectParams{
			Project:         projectReq,
			ProjectNameOrID: e.proj.ObjectMeta.Name,
		})
		if err != nil {
			// l.Error(err, "error updating project...")
			// Update the status field of the resource in case we get errors creating the project
			meta.SetStatusCondition(&e.proj.Status.Conditions, metav1.Condition{Type: typeAvailable,
				Status: metav1.ConditionFalse, Reason: "Reconciling",
				Message: fmt.Sprintf("Failed to update Harbor project for the custom resource (%s): (%s)", e.proj.Name, err)})

			if err := e.client.Status().Update(ctx, e.proj); err != nil {
				// l.Error(err, "failed to update Project status")
				// return ctrl.Result{}, err
				return err
			}
			// return ctrl.Result{}, err
			return err
		}
		// Project updated successfully. Requeue and check state again
		//return ctrl.Result{RequeueAfter: time.Minute}, nil
		//return nil
	}
	return nil
}

func (e ProjectEnsurer) projectIsNotFound(ctx context.Context, key string) bool {
	_, err := e.clientset.V2().Project.HeadProject(ctx, &project.HeadProjectParams{ProjectName: key})
	if err != nil {
		if _, notFound := err.(*project.HeadProjectNotFound); notFound {
			return true
		}
		return true
	}
	return false
}
