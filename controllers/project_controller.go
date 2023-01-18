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
	//"crypto/tls"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	//"net/http"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"net/url"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	//oclient "github.com/go-openapi/runtime/client"
	"github.com/goharbor/go-client/pkg/harbor"
	//hclient "github.com/goharbor/go-client/pkg/sdk/v2.0/client"
	harborv1alpha1 "github.com/middlewaregruppen/harbor-operator/api/v1alpha1"
)

// ProjectReconciler reconciles a Project object
type ProjectReconciler struct {
	client.Client
	Scheme *runtime.Scheme
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

	l.Info("Parsing URL")
	u, err := url.Parse("https://harbor-portal:31443")
	if err != nil {
		l.Error(err, "Couldn't parse URL")
		return ctrl.Result{}, err
	}

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		InsecureSkipVerify: true,
	// 	},
	// }

	// auth := oclient.BasicAuth("amir", "amir")

	// TODO(user): your logic here
	// c := hclient.Config{
	// 	URL:       u,
	// 	Transport: tr,
	// 	AuthInfo:  auth,
	// }
	c := &harbor.ClientSetConfig{
		URL:      u.String(),
		Insecure: true,
		Username: "admin",
		Password: "Harbor12345",
	}

	l.Info("Setting up clientset")
	cs, err := harbor.NewClientSet(c)
	if err != nil {
		l.Error(err, "Couldn't create clientset")
		return ctrl.Result{}, err
	}

	v2client := cs.V2()
	l.Info(fmt.Sprintf("Getting Project client %+v", v2client))
	if v2client.Project != nil {
		projectOK, err := v2client.Project.ListProjects(ctx, &project.ListProjectsParams{})
		if err != nil {
			l.Error(err, "Couldn't list projects")
			return ctrl.Result{}, err
		}
		l.Info(fmt.Sprintf("Got %d projects", projectOK.XTotalCount))
		for _, p := range projectOK.Payload {
			l.Info(fmt.Sprintf("Name: %s", p.Name))
		}
	} else {
		l.Info("Couln't get OK")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProjectReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&harborv1alpha1.Project{}).
		Complete(r)
}
