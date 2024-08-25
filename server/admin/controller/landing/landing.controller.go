package settings

import (
	"main/pkg/engine/controller"
	"main/server/admin/view/pages/landing"
)

type Param struct {
	Tab string `param:"tab"`
}

func index(ctx *controller.Context, param *Param) error {
	if param.Tab == "" {
		param.Tab = "contacts"
	}
	return ctx.Html(landing.Page(param.Tab))
}