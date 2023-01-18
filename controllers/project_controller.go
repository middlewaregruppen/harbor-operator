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
	"crypto/tls"
	"net/url"
	//"crypto/tls"
	//"fmt"
	oruntime "github.com/go-openapi/runtime"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	//"github.com/goharbor/go-client/pkg/harbor"
	oclient "github.com/go-openapi/runtime/client"
	hclient "github.com/goharbor/go-client/pkg/sdk/v2.0/client"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	harborv1alpha1 "github.com/middlewaregruppen/harbor-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
)

// ProjectReconciler reconciles a Project object
type ProjectReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	// clientSet is the Harbor ClientSet
	harborclient *HarborClient
}

type HarborClient struct {
	//Transport http.RoundTripper
	Transport oruntime.ClientTransport
	API       *hclient.HarborAPI
	URL       *url.URL
	AuthInfo  oruntime.ClientAuthInfoWriter
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

	proj := &harborv1alpha1.Project{}
	err := r.Get(ctx, req.NamespacedName, proj)
	if err != nil && errors.IsNotFound(err) {
		l.Error(err, "couldn't get project")
		return ctrl.Result{}, nil
	}

	//v2client := r.harborclient.API
	params := &project.HeadProjectParams{ProjectName: proj.ObjectMeta.Name}
	result, err := transport.Submit(&oruntime.ClientOperation{
		ID:                 "headProject",
		Method:             "HEAD",
		PathPattern:        "/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &project.HeadProjectReader{},
		AuthInfo:           r.harborclient.AuthInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})

	// ok, err := v2client.Project.HeadProject(ctx, &project.HeadProjectParams{ProjectName: proj.ObjectMeta.Name})
	// if err != nil {
	// 	l.Error(err, "couldn't head project")
	// }

	isPublic := true
	if proj.Spec.IsPrivate {
		isPublic = false
	}
	projectReq := &models.ProjectReq{
		ProjectName: proj.ObjectMeta.Name,
		Public:      &isPublic,
		//RegistryID: proj.Spec.Registry,
	}

	// Project exists
	if ok.IsCode(http.StatusOK) {
		updateParams := &project.UpdateProjectParams{
			Project:         projectReq,
			ProjectNameOrID: proj.ObjectMeta.Name,
		}
		_, err = v2client.Project.UpdateProject(ctx, updateParams)
		if err != nil {
			l.Error(err, "couldn't update project")
			return ctrl.Result{}, err
		}
	}

	// Project is not found
	if !ok.IsSuccess() {
		createParams := &project.CreateProjectParams{
			Project: projectReq,
		}
		_, err = v2client.Project.CreateProject(ctx, createParams)
		if err != nil {
			l.Error(err, "couldn't create project")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (h *HarborClient) ProjectNotFound(ctx context.Context, proj *harborv1alpha1.Project) (bool, error) {
	transport := oclient.New(h.URL.Hostname(), h.URL.EscapedPath(), []string{"http", "https"})

	params := &project.HeadProjectParams{ProjectName: proj.ObjectMeta.Name}
	result, err := transport.Submit(&oruntime.ClientOperation{
		ID:                 "headProject",
		Method:             "HEAD",
		PathPattern:        "/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &project.HeadProjectReader{},
		AuthInfo:           h.AuthInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*project.HeadProjectNotFound), nil

}

// SetupWithManager sets up the controller with the Manager.
func (r *ProjectReconciler) SetupWithManager(mgr ctrl.Manager, harborclient *HarborClient) error {
	r.harborclient = harborclient
	c := hclient.Config{
		URL:      harborclient.URL,
		AuthInfo: harborclient.AuthInfo,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	api := hclient.New(c)
	r.harborclient.API = api
	return ctrl.NewControllerManagedBy(mgr).
		For(&harborv1alpha1.Project{}).
		Complete(r)
}
