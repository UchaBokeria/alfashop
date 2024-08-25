package events

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
	var Events []model.Posts
	var Interface model.Interface

	storage.DB.Order("created_at desc").
			   Preload("Thumbnail").
			   Find(&Events, model.Posts{TypeID: 5})

	storage.DB.Preload("Events.Thumbnail").
			   Last(&Interface)

	EventsSelect, _ := helpers.ToSelect(Events, "ID", "Title")
	return ctx.Html(tabs.Events(Interface.Events, EventsSelect))
}

func create(ctx *controller.Context, params *EventsDto) error {
	storage.DB.Model(repository.FindByID[model.Interface](1)).
			   Association("Events").
			   Append(repository.FindByID[model.Posts](params.EventID))

	return index(ctx)
}

func remove(ctx *controller.Context, ID *dto.ByID) error {
	storage.DB.Model(repository.FindByID[model.Interface](1)).
			   Association("Events").Delete(repository.FindByID[model.Posts](ID.ID))

	return index(ctx)
}
