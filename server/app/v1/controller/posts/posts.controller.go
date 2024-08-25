package posts

import (
	"main/internal/model"
	"main/internal/types/dto"
	"main/pkg/engine/controller"
	"main/pkg/storage"
	view "main/server/app/v1/view/pages/posts"
)

func index(ctx *controller.Context, filter *Filters) error {
	var Posts []model.Posts
	var Types []model.Posts_types

	var Where = &model.Posts{Public: true}

	if filter != nil {
		Where.TypeID = filter.Type
	}

	storage.DB.Find(&Types)
	storage.DB.
		Order("Posts.created_at desc").
		Where(Where).
		Preload("Thumbnail").
		Find(&Posts)

	return ctx.Html(view.List(Posts, Types, ctx.QueryParam("type")))
}

func detail(ctx *controller.Context, ID *dto.ByID) error {
	var Posts model.Posts

	storage.DB.Where(&model.Posts{Public: true}).
		Preload("Thumbnail").
		Last(&Posts, ID)

	return ctx.Html(view.PostsDetails(Posts))
}
