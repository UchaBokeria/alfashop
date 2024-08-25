package company

import (
	"main/internal/model"
	"main/pkg/engine/controller"
	"main/pkg/storage"
	view "main/server/app/v1/view/pages/company"
)

func faq(ctx *controller.Context) error {
	var Faq []model.Faq
	storage.DB.Find(&Faq)
	return ctx.Html(view.Faq(Faq))
}

func index(ctx *controller.Context, param *CompanyDto) error {
	var Pages []model.Company
	storage.DB.Find(&Pages)

	var page model.Company
	storage.DB.Where("path = ?", param.Page).First(&page)
	return ctx.Html(view.Company(Pages, page))
}
