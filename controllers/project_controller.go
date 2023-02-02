/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"

	harborv1alpha1 "github.com/middlewaregruppen/harbor-operator/api/v1alpha1"
	"github.com/middlewaregruppen/harbor-operator/pkg/util"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// Definitions to manage status conditions
const (
	// ProjectStatusAvailable represents the status of the Deployment reconciliation
	ProjectStatusAvailable = "Available"
	// ProjectStatusDegraded represents the status used when the custom resource is deleted and the finalizer operations are must to occur.
	ProjectStatusDegraded = "Degraded"
)

const projectFinalizers = "harbor.mdlwr.com/finalizer"

// ProjectReconciler reconciles a Project object
type ProjectReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	// clientSet is the Harbor ClientSet
	clientset *harbor.ClientSet
}

func (e ProjectReconciler) getProjectId(ctx context.Context, name string) (*int32, error) {
	projOK, err := e.clientset.V2().Project.ListProjects(ctx, &project.ListProjectsParams{Name: &name})
	if err != nil {
		return nil, err
	}
	if len(projOK.Payload) > 0 {
		return &projOK.Payload[0].ProjectID, nil
	}
	return nil, nil
}

//+kubebuilder:rbac:groups=harbor.mdlwr.com,resources=projects,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=harbor.mdlwr.com,resources=projects/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=harbor.mdlwr.com,resources=projects/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Project object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *ProjectReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	// Fetch Project - This ensures that the cluster has resources of type Project.
	// Stops reconciliation if not found, for example if the CRD's has not been applied
	proj := &harborv1alpha1.Project{}
	if err := r.Get(ctx, req.NamespacedName, proj); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Don't do anything if no harborservice is defined
	if len(proj.Spec.Harbor) == 0 {
		return ctrl.Result{}, nil
	}

	// Get the HarborService instance
	hs := &harborv1alpha1.HarborService{}
	if err := r.Get(ctx, types.NamespacedName{Namespace: proj.Namespace, Name: proj.Spec.Harbor}, hs); err != nil {
		return ctrl.Result{}, errors.New(fmt.Sprintf("couldn't find HarborService (%s/%s): %s", proj.Namespace, proj.Spec.Harbor, err))
	}

	// Get harbor clientset
	cs, err := GetClientSet(ctx, r.Client, hs)
	if err != nil {
		return ctrl.Result{}, err
	}
	r.clientset = cs

	// Check status field here
	// Let's just set the status as Unknown when no status are available
	if proj.Status.Conditions == nil || len(proj.Status.Conditions) == 0 {
		meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: ProjectStatusAvailable, Status: metav1.ConditionUnknown, Reason: "Reconciling", Message: "Starting reconciliation"})
		if err := r.Status().Update(ctx, proj); err != nil {
			return ctrl.Result{}, err
		}

		// Get the Project resource again so that we don't encounter any "the object has been modified"-errors
		if err := r.Get(ctx, req.NamespacedName, proj); err != nil {
			return ctrl.Result{}, err
		}
	}

	// Add finalizers that will be handled later during delete events
	if !controllerutil.ContainsFinalizer(proj, projectFinalizers) {
		if ok := controllerutil.AddFinalizer(proj, projectFinalizers); !ok {
			return ctrl.Result{Requeue: true}, nil
		}
		if err := r.Update(ctx, proj); err != nil {
			return ctrl.Result{}, err
		}
	}

	isPublic := true
	if proj.Spec.IsPrivate {
		isPublic = false
	}
	projectReq := &models.ProjectReq{
		ProjectName: proj.Name,
		Metadata: &models.ProjectMetadata{
			Public: util.BoolToString(isPublic),
		},
	}

	// Check if Project exists in Harbor by doing a HEAD requests. If not found, the project is created
	_, err = r.clientset.V2().Project.HeadProject(ctx, &project.HeadProjectParams{ProjectName: proj.ObjectMeta.Name})
	if err != nil {
		if _, notFound := err.(*project.HeadProjectNotFound); notFound {
			_, err := r.clientset.V2().Project.CreateProject(ctx, &project.CreateProjectParams{Project: projectReq})
			if err != nil {
				// Update the status field of the resource in case we get errors creating the project
				meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: ProjectStatusAvailable,
					Status: metav1.ConditionFalse, Reason: "Reconciling",
					Message: fmt.Sprintf("Failed to create Harbor project for the custom resource (%s): (%s)", proj.Name, err)})

				// Get the Project resource again so that we don't encounter any "the object has been modified"-errors
				if err = r.Get(ctx, req.NamespacedName, proj); err != nil {
					return ctrl.Result{}, err
				}

				if err := r.Status().Update(ctx, proj); err != nil {
					return ctrl.Result{}, err
				}
				return ctrl.Result{}, err
			}
			// Project was created, re-queue so that we can re-check
			return ctrl.Result{RequeueAfter: time.Minute}, nil
		}
		return ctrl.Result{}, err
	}

	// We need the project ID. Some Harbor versions don't support using names
	id, err := r.getProjectId(ctx, proj.Name)
	if err != nil {
		return ctrl.Result{}, err
	}
	if id == nil {
		l.Info("COuldn't find the id")
		return ctrl.Result{}, nil
	}
	idstr := fmt.Sprintf("%d", *id)

	// Check if resource is marked to be deleted
	if proj.GetDeletionTimestamp() != nil {
		// Perform finalizers before deleting resource from cluster
		if controllerutil.ContainsFinalizer(proj, projectFinalizers) {

			// Add Degraded status to begin the process of terminating resources
			meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: ProjectStatusDegraded,
				Status: metav1.ConditionUnknown, Reason: "Finalizing",
				Message: fmt.Sprintf("Performing finalizer operations for the custom resource: %s ", proj.Name)})

			// Update the resource with updated status
			if err := r.Status().Update(ctx, proj); err != nil {
				return ctrl.Result{}, err
			}

			// TODO: run finalizers here. Always delete resources that belong to this CRD before proceeding further
			// Delete project from Harbor after finalizers have been carried out
			_, err := r.clientset.V2().Project.DeleteProject(ctx, &project.DeleteProjectParams{ProjectNameOrID: idstr})
			if nil != err {

				// Update the status field of the resource in case we get errors creating the project
				meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: ProjectStatusAvailable,
					Status: metav1.ConditionFalse, Reason: "Reconciling",
					Message: fmt.Sprintf("Failed to delete Harbor project for the custom resource (%s): (%s)", proj.Name, err)})

				if err := r.Status().Update(ctx, proj); err != nil {
					return ctrl.Result{}, err
				}
				return ctrl.Result{}, err
			}

			// Get the Project resource again so that we don't encounter any "the object has been modified"-errors
			if err = r.Get(ctx, req.NamespacedName, proj); err != nil {
				return ctrl.Result{}, err
			}

			meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: ProjectStatusDegraded,
				Status: metav1.ConditionTrue, Reason: "Finalizing",
				Message: fmt.Sprintf("Finalizer operations for custom resource %s name were successfully accomplished", proj.Name)})

			if err := r.Status().Update(ctx, proj); err != nil {
				return ctrl.Result{}, err
			}

			// Remove finalizers and update status field of the resource
			if ok := controllerutil.RemoveFinalizer(proj, projectFinalizers); !ok {
				return ctrl.Result{Requeue: true}, nil
			}
			if err := r.Update(ctx, proj); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	// If we got this far it means that project exists and we can update project in Harbor instead.,
	// Note that we always update, overwriting existing project with the state of that in the CRD.
	// We might want to consider only updating when there are changes.
	_, err = r.clientset.V2().Project.UpdateProject(context.TODO(), &project.UpdateProjectParams{
		Project:         projectReq,
		ProjectNameOrID: idstr,
	})
	if err != nil {
		l.Error(err, "Couldn't update project")
		if err := r.Get(ctx, req.NamespacedName, proj); err != nil {
			return ctrl.Result{}, err
		}
		// Update the status field of the resource in case we get errors creating the project
		meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: ProjectStatusAvailable,
			Status: metav1.ConditionFalse, Reason: "Reconciling",
			Message: fmt.Sprintf("Failed to update Harbor project for the custom resource (%s): (%s)", proj.Name, err)})

		if err := r.Status().Update(ctx, proj); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, err
	}

	// Get the Project resource again so that we don't encounter any "the object has been modified"-errors
	if err = r.Get(ctx, req.NamespacedName, proj); err != nil {
		return ctrl.Result{}, err
	}
	// The following implementation will update the status
	// TODO: We might want to consider putting this into the Ensurer interface
	meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: ProjectStatusAvailable,
		Status: metav1.ConditionTrue, Reason: "Reconciling",
		Message: fmt.Sprintf("Harbor Project for custom resource (%s) created successfully", proj.Name)})

	if err := r.Status().Update(ctx, proj); err != nil {
		l.Error(err, "failed to update Project status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProjectReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&harborv1alpha1.Project{}).
		Complete(r)
}
