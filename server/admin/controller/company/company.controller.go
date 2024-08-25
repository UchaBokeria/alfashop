package company

import (
	"main/internal/model"
	"main/internal/types/dto"
	"main/pkg/engine/controller"
	"main/pkg/storage"
	"main/server/admin/repository"
	"main/server/admin/view/pages/company"
)

func index(ctx *controller.Context) error {
	var companyPages []model.Company
	storage.DB.Find(&companyPages)
	return ctx.Html(company.Page(companyPages))
}

func unique(ctx *controller.Context, ID *dto.ByID) error {
	var Company model.Company
	storage.DB.First(&Company, ID)
	return ctx.Html(company.Update(Company))
}

func newForm(ctx *controller.Context) error {
	return ctx.Html(company.Create())
}

func create(ctx *controller.Context, params *CompanyCreateDto) error {
	repository.Create[model.Company](&model.Company{
		Name: params.Name,
		Path: params.Path,
		Body: params.Body,
	})

	return index(ctx)
}

func update(ctx *controller.Context, params *CompanyUpdateDto) error {
	repository.Update[model.Company](params.ID, &model.Company{
		Name: params.Name,
		Path: params.Path,
		Body: params.Body,
	})

	return index(ctx)
}

func remove(ctx *controller.Context, ID *dto.ByID) error {
	repository.Delete[model.Company](ID)
	return index(ctx)
}
