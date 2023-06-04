package http

import (
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/presentation/http/controller"
	v1_user_delete "go-clean-api/cmd/presentation/http/controller/v1/users/delete"
	v1_user "go-clean-api/cmd/presentation/http/controller/v1/users/find-by-di"
	controllerv1userregister "go-clean-api/cmd/presentation/http/controller/v1/users/register"
	"go-clean-api/cmd/presentation/http/middlewares"
)

func loadControllers(container *container.Container) []controller.Controller {
	return []controller.Controller{
		controllerv1userregister.New(container),
		v1_user.New(container),
		v1_user_delete.New(container),
	}
}

func loadMiddlewaresGlobals() []controller.Middleware {
	return []controller.Middleware{
		middlewares.Global,
	}
}
