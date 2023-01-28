package ensure

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
)

type Ensurer interface {
	Setup(context.Context) (ctrl.Result, error)
	HandleDelete(context.Context) (ctrl.Result, error)
	HandleCreate(context.Context) (ctrl.Result, error)
	HandleUpdate(context.Context) (ctrl.Result, error)
}
