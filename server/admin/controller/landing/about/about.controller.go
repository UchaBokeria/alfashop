package about

import (
	"main/internal/model"
	"main/pkg/engine/controller"
	"main/pkg/storage"
	"main/server/admin/view/pages/landing/tabs"
	"net/http"
)

func index(ctx *controller.Context) error {
	var About model.Interface_about
	storage.DB.Last(&About)

	return ctx.Html(tabs.About(About.Body))
}

func contact(ctx *controller.Context, Body *AboutDto) error {
	var About model.Interface_about

	if match := storage.DB.Last(&About); match.Error != nil {
		return ctx.String(http.StatusNotFound, "")
	}

	if result := storage.DB.Model(About).Updates(&model.Interface_about{Body: Body.Content}); result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Abouts model.Interface_about
	storage.DB.Last(&Abouts)
	return ctx.Html(tabs.Presentation(Abouts.Body))
}