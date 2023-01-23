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
	"fmt"
	"time"

	//"crypto/tls"

	"github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"

	//"github.com/goharbor/go-client/pkg/harbor"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	harborv1alpha1 "github.com/middlewaregruppen/harbor-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const projectFinalizers = "harbor.mdlwr.com/finalizer"

// Definitions to manage status conditions
const (
	// typeAvailable represents the status of the Deployment reconciliation
	typeAvailable = "Available"
	// typeDegraded represents the status used when the custom resource is deleted and the finalizer operations are must to occur.
	typeDegraded = "Degraded"
)

// ProjectReconciler reconciles a Project object
type ProjectReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	// clientSet is the Harbor ClientSet
	clientset *harbor.ClientSet
}

func (r *ProjectReconciler) projectIsNotFound(ctx context.Context, key string) bool {
	_, err := r.clientset.V2().Project.HeadProject(ctx, &project.HeadProjectParams{ProjectName: key})
	if err != nil {
		if _, notFound := err.(*project.HeadProjectNotFound); notFound {
			return true
		}
	}
	return false
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
	err := r.Get(ctx, req.NamespacedName, proj)
	if err != nil {
		if errors.IsNotFound(err) {
			l.Info("Project resource not found. It has probably been deleted")
			return ctrl.Result{}, nil
		}
		l.Error(err, "failed to get Project")
		return ctrl.Result{}, nil
	}

	// Check status field here

	// Add finalizers that will be handled later during delete events
	if !controllerutil.ContainsFinalizer(proj, projectFinalizers) {
		if ok := controllerutil.AddFinalizer(proj, projectFinalizers); !ok {
			l.Error(err, "failed to add finalizer into the custom resource")
			return ctrl.Result{Requeue: true}, nil
		}
		if err = r.Update(ctx, proj); err != nil {
			l.Error(err, "failed to update custom resource to add finalizer")
			return ctrl.Result{}, err
		}
	}

	// Check if the Project is marked to be deleted, which is
	// indicated by the deletion timestamp on the resource.
	if proj.GetDeletionTimestamp() != nil {
		// Perform finalizers before deleting resource from cluster
		if controllerutil.ContainsFinalizer(proj, projectFinalizers) {

			// Add Degraded status to begin the process of terminating resources
			meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: typeDegraded,
				Status: metav1.ConditionUnknown, Reason: "Finalizing",
				Message: fmt.Sprintf("Performing finalizer operations for the custom resource: %s ", proj.Name)})

			// Update the resource with updated status
			if err := r.Status().Update(ctx, proj); err != nil {
				l.Error(err, "failed to update Project status")
				return ctrl.Result{}, err
			}

			// TODO: run finalizers here

			// Get the Project resource again so that we don't encounter any "the object has been modified"-errors
			if err = r.Get(ctx, req.NamespacedName, proj); err != nil {
				l.Error(err, "failed to re-fetch project")
				return ctrl.Result{}, err
			}

			meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: typeDegraded,
				Status: metav1.ConditionTrue, Reason: "Finalizing",
				Message: fmt.Sprintf("Finalizer operations for custom resource %s name were successfully accomplished", proj.Name)})

			if err := r.Status().Update(ctx, proj); err != nil {
				l.Error(err, "failed to update Project status")
				return ctrl.Result{}, err
			}

			// Remove finalizers and update status field of the resource
			if ok := controllerutil.RemoveFinalizer(proj, projectFinalizers); !ok {
				l.Error(err, "failed to remove finalizer for Project")
				return ctrl.Result{Requeue: true}, nil
			}
			if err := r.Update(ctx, proj); err != nil {
				l.Error(err, "failed to remove finalizer for Project")
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	// Define a new project request that we use to update or create new projects in Harbor
	isPublic := true
	if proj.Spec.IsPrivate {
		isPublic = false
	}
	projectReq := &models.ProjectReq{
		ProjectName: proj.ObjectMeta.Name,
		Metadata: &models.ProjectMetadata{
			Public: boolToString(isPublic),
		},
	}

	// Check if the actual Project in Harbor exists. If not create a new one
	if r.projectIsNotFound(ctx, proj.Name) {
		_, err = r.clientset.V2().Project.CreateProject(ctx, &project.CreateProjectParams{Project: projectReq})
		if err != nil {
			// Update the status field of the resource in case we get errors creating the project
			meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: typeAvailable,
				Status: metav1.ConditionFalse, Reason: "Reconciling",
				Message: fmt.Sprintf("Failed to create Harbor project for the custom resource (%s): (%s)", proj.Name, err)})

			if err := r.Status().Update(ctx, proj); err != nil {
				l.Error(err, "failed to update Project status")
				return ctrl.Result{}, err
			}
			return ctrl.Result{}, err
		}

		// Check if the actual project in Harbor exists. If so then update the existing one
		if !r.projectIsNotFound(ctx, proj.Name) {
			// _, err = r.clientset.V2().Project.GetProject(ctx, &project.GetProjectParams{ProjectNameOrID: proj.Name})
			// if err != nil {
			// 	l.Error(err, "couldn't get project")
			// 	return ctrl.Result{}, err
			// }
			_, err = r.clientset.V2().Project.UpdateProject(ctx, &project.UpdateProjectParams{
				Project:         projectReq,
				ProjectNameOrID: proj.ObjectMeta.Name,
			})
			if err != nil {
				// Update the status field of the resource in case we get errors creating the project
				meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: typeAvailable,
					Status: metav1.ConditionFalse, Reason: "Reconciling",
					Message: fmt.Sprintf("Failed to update Harbor project for the custom resource (%s): (%s)", proj.Name, err)})

				if err := r.Status().Update(ctx, proj); err != nil {
					l.Error(err, "failed to update Project status")
					return ctrl.Result{}, err
				}
				return ctrl.Result{}, err
			}
		}

		// Deployment created successfully
		// We will requeue the reconciliation so that we can ensure the state
		// and move forward for the next operations
		return ctrl.Result{RequeueAfter: time.Minute}, nil
	}

	// The following implementation will update the status
	meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: typeAvailable,
		Status: metav1.ConditionTrue, Reason: "Reconciling",
		Message: fmt.Sprintf("Harbor Project for custom resource (%s) created successfully", proj.Name)})

	if err := r.Status().Update(ctx, proj); err != nil {
		l.Error(err, "failed to update Memcached status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProjectReconciler) SetupWithManager(mgr ctrl.Manager, clientset *harbor.ClientSet) error {
	r.clientset = clientset
	return ctrl.NewControllerManagedBy(mgr).
		For(&harborv1alpha1.Project{}).
		Complete(r)
}

func ptrBool(b bool) *bool {
	return &b
}

func ptrInt64(b int64) *int64 {
	return &b
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
