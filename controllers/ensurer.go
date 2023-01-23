package controllers

import (
	"context"
	"errors"
	"fmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	harborv1alpha1 "github.com/middlewaregruppen/harbor-operator/api/v1alpha1"
)

type Ensurer interface {
	Finalize() error
	Ensure(context.Context, interface{}) error
}

type ProjectEnsurer struct {
}

func (e ProjectEnsurer) Ensure(ctx context.Context, i interface{}) error {
	p, ok := i.(harborv1alpha1.Project)
	if !ok {
		return errors.New(fmt.Sprintf("interface %v is not harborv1alpha1.Project"))
	}
	return nil
}

func NewProjectEnsurer(p *project.Client) ProjectEnsurer {

}
