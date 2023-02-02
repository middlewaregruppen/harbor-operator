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
	"net/url"
	"time"

	"github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/health"
	harborv1alpha1 "github.com/middlewaregruppen/harbor-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
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
	// HarborAvailable
	HarborStatusReady = "Ready"
	// HarborStatusNotReady
	HarborStatusNotReady = "NotReady"
)

// HarborServiceReconciler reconciles a HarborService object
type HarborServiceReconciler struct {
	client.Client
	Scheme    *runtime.Scheme
	clientset *harbor.ClientSet
}

func GetClientSet(ctx context.Context, c client.Client, h *harborv1alpha1.HarborService) (*harbor.ClientSet, error) {

	// Fetch credentials from secrets through the secretRef
	// TODO: Ignore not found errors here
	sec := &corev1.Secret{}
	if err := c.Get(ctx, types.NamespacedName{Namespace: h.Namespace, Name: h.Spec.SecretRef.Name}, sec); err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't find secret (%s/%s): %s", h.Namespace, h.Spec.SecretRef.Name, err))
	}

	var ustr string
	// Use URL field in the external backend
	if h.Spec.External != nil {
		ustr = fmt.Sprintf("%s://%s:%d", h.Spec.Scheme, h.Spec.External.Host, h.Spec.External.Port)
	}

	// Use service to create an URL
	if h.Spec.Internal != nil {
		ustr = fmt.Sprintf("%s://%s:%d", h.Spec.Scheme, h.Spec.Internal.Name, h.Spec.Internal.Port.Number)
	}

	u, err := url.Parse(ustr)
	if err != nil {
		return nil, err
	}

	// Setup Harbor clientset
	config := &harbor.ClientSetConfig{
		URL:      u.String(),
		Insecure: h.Spec.Insecure,
		Username: string(sec.Data["username"]),
		Password: string(sec.Data["password"]),
	}

	return harbor.NewClientSet(config)
}

//+kubebuilder:rbac:groups=harbor.mdlwr.com,resources=harborservices,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=harbor.mdlwr.com,resources=harborservices/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=harbor.mdlwr.com,resources=harborservices/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the HarborService object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *HarborServiceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	// Fetch HarborService - This ensures that the cluster has resources of type HarborService.
	// Stops reconciliation if not found, for example if the CRD's has not been applied
	hs := &harborv1alpha1.HarborService{}
	if err := r.Get(ctx, req.NamespacedName, hs); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Don't do anything if no backends are configured
	if hs.Spec.External == nil && hs.Spec.Internal == nil {
		return ctrl.Result{}, nil
	}

	// Get harbor clientset
	cs, err := GetClientSet(ctx, r.Client, hs)
	if err != nil {
		return ctrl.Result{}, err
	}
	r.clientset = cs

	// Check status field here
	// Let's just set the status as Unknown when no status are available
	if hs.Status.Conditions == nil || len(hs.Status.Conditions) == 0 {
		meta.SetStatusCondition(&hs.Status.Conditions, metav1.Condition{Type: HarborStatusNotReady, Status: metav1.ConditionUnknown, Reason: "Reconciling", Message: "Starting reconciliation"})
		if err := r.Status().Update(ctx, hs); err != nil {
			return ctrl.Result{}, err
		}

		// Get the HarborService resource again so that we don't encounter any "the object has been modified"-errors
		if err := r.Get(ctx, req.NamespacedName, hs); err != nil {
			return ctrl.Result{}, err
		}
	}

	// Add finalizers that will be handled later during delete events
	if !controllerutil.ContainsFinalizer(hs, projectFinalizers) {
		if ok := controllerutil.AddFinalizer(hs, projectFinalizers); !ok {
			return ctrl.Result{Requeue: true}, nil
		}
		if err := r.Update(ctx, hs); err != nil {
			return ctrl.Result{}, err
		}
	}

	// Check if resource is marked to be deleted
	if hs.GetDeletionTimestamp() != nil {
		if controllerutil.ContainsFinalizer(hs, projectFinalizers) {

			// Add Degraded status to begin the process of terminating resources
			meta.SetStatusCondition(&hs.Status.Conditions, metav1.Condition{Type: HarborStatusNotReady,
				Status: metav1.ConditionUnknown, Reason: "Finalizing",
				Message: fmt.Sprintf("Performing finalizer operations for the custom resource: %s ", hs.Name)})

			// Update the resource with updated status
			if err := r.Status().Update(ctx, hs); err != nil {
				return ctrl.Result{}, err
			}

			// TODO: run finalizers here. Always delete resources that belong to this CRD before proceeding further

			// Get the resource again so that we don't encounter any "the object has been modified"-errors
			if err := r.Get(ctx, req.NamespacedName, hs); err != nil {
				return ctrl.Result{}, err
			}

			meta.SetStatusCondition(&hs.Status.Conditions, metav1.Condition{Type: HarborStatusNotReady,
				Status: metav1.ConditionTrue, Reason: "Finalizing",
				Message: fmt.Sprintf("Finalizer operations for custom resource %s name were successfully accomplished", hs.Name)})

			if err := r.Status().Update(ctx, hs); err != nil {
				return ctrl.Result{}, err
			}

			// Remove finalizers and update status field of the resource
			if ok := controllerutil.RemoveFinalizer(hs, projectFinalizers); !ok {
				return ctrl.Result{Requeue: true}, nil
			}
			if err := r.Update(ctx, hs); err != nil {
				return ctrl.Result{}, err
			}
		}
	}

	// Check connectivity using the health API
	ok, err := r.clientset.V2().Health.GetHealth(ctx, health.NewGetHealthParams())
	if err != nil {
		l.Error(err, "failed checking health of harbor")
		// Update the status field of the resource in case we get errors checking the health
		meta.SetStatusCondition(&hs.Status.Conditions, metav1.Condition{Type: HarborStatusNotReady,
			Status: metav1.ConditionFalse, Reason: "Reconciling",
			Message: fmt.Sprintf("failed getting health of Harbor for the custom resource (%s): (%s)", hs.Name, err)})

		if err := r.Status().Update(ctx, hs); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, err
	}

	// Get the resource again so that we don't encounter any "the object has been modified"-errors
	if err = r.Get(ctx, req.NamespacedName, hs); err != nil {
		return ctrl.Result{}, err
	}
	// The following implementation will update the status
	meta.SetStatusCondition(&hs.Status.Conditions, metav1.Condition{Type: HarborStatusReady,
		Status: metav1.ConditionTrue, Reason: "Reconciling",
		Message: fmt.Sprintf("Got OK from Harbor service (%s) with status (%s)", hs.Name, ok.Payload.Status)})

	if err := r.Status().Update(ctx, hs); err != nil {
		return ctrl.Result{}, err
	}

	// Re-queue and check again after a minute
	return ctrl.Result{RequeueAfter: time.Minute}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *HarborServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	//r.clientset = clientset
	return ctrl.NewControllerManagedBy(mgr).
		For(&harborv1alpha1.HarborService{}).
		Complete(r)
}
