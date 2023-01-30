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

	"github.com/goharbor/go-client/pkg/harbor"
	harborv1alpha1 "github.com/middlewaregruppen/harbor-operator/api/v1alpha1"
	"github.com/middlewaregruppen/harbor-operator/pkg/ensure"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ProjectReconciler reconciles a Project object
type ProjectReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	// clientSet is the Harbor ClientSet
	clientset *harbor.ClientSet
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

	e, err := ensure.NewProjectEnsurer(r.Client, r.clientset, req, proj)
	if err != nil {
		l.Error(err, "failed setting up projectensurer")
		return ctrl.Result{}, err
	}

	// Adds finalizers and update our resource with a status field
	//if e.proj.Status.Conditions == nil || len(e.proj.Status.Conditions) == 0 {
	l.Info("Setting up resource")
	if _, err = e.Setup(ctx); err != nil {
		return ctrl.Result{}, err
	}

	// Check if the Project is marked to be deleted, which is
	// indicated by the deletion timestamp on the resource.
	if proj.GetDeletionTimestamp() != nil {
		l.Info("Handling delete of resource")
		return e.HandleDelete(ctx)
	}

	// If project is not found in Harbor, then create
	if e.ProjectNotFound(ctx, proj.Name) {
		l.Info("Handling create of resource")
		return e.HandleCreate(ctx)
	}

	// If we got this far it means that project exists and we can update project in Harbor instead.,
	// Note that we always update, overwriting existing project with the state of that in the CRD.
	// We might want to consider only updating when there are changes.
	l.Info("Handling update of resource")
	if _, err = e.HandleUpdate(ctx); err != nil {
		return ctrl.Result{}, err
	}

	// The following implementation will update the status
	// TODO: We might want to consider putting this into the Ensurer interface
	meta.SetStatusCondition(&proj.Status.Conditions, metav1.Condition{Type: ensure.ProjectAvailable,
		Status: metav1.ConditionTrue, Reason: "Reconciling",
		Message: fmt.Sprintf("Harbor Project for custom resource (%s) created successfully", proj.Name)})

	if err := r.Status().Update(ctx, proj); err != nil {
		l.Error(err, "failed to update Project status")
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
