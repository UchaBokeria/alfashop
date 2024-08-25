package presentation

import (
	"main/internal/model"
	"main/pkg/engine/controller"
	"main/pkg/storage"
	"main/server/admin/view/pages/landing/tabs"
	"net/http"
)

func index(ctx *controller.Context) error {
	var Interface model.Interface
	storage.DB.Last(&Interface)
	return ctx.Html(tabs.Presentation(Interface.Presentation))
}

func presentation(ctx *controller.Context, Body *PresentationDto) error {
	var Interface model.Interface

	if match := storage.DB.Last(&Interface); match.Error != nil {
		return ctx.String(http.StatusNotFound, "")
	}

	if result := storage.DB.Model(Interface).Updates(&model.Interface{Presentation: Body.Content}); result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	return ctx.Html(tabs.Presentation(Interface.Presentation))
}
