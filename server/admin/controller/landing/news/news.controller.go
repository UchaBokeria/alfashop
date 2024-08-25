package news

import (
	"main/internal/model"
	"main/internal/types/dto"
	"main/pkg/engine/controller"
	"main/pkg/helpers"
	"main/pkg/storage"
	"main/server/admin/repository"
	"main/server/admin/view/pages/landing/tabs"
)

func index(ctx *controller.Context) error {
	var News []model.Posts
	var Interface model.Interface

	storage.DB.Order("created_at desc").
			   Preload("Thumbnail").
			   Find(&News)

	storage.DB.Preload("News.Thumbnail").
			   Last(&Interface)

	NewsSelect, _ := helpers.ToSelect(News, "ID", "Title")
	return ctx.Html(tabs.News(Interface.News, NewsSelect))
}

func create(ctx *controller.Context, params *NewsDto) error {
	storage.DB.Model(repository.FindByID[model.Interface](1)).
			   Association("News").
			   Append(repository.FindByID[model.Posts](params.NewsID))

	return index(ctx)
}

func remove(ctx *controller.Context, ID *dto.ByID) error {
	storage.DB.Model(repository.FindByID[model.Interface](1)).
			   Association("News").Delete(repository.FindByID[model.Posts](ID.ID))

	return index(ctx)
}
