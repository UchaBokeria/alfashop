package newvideos

import (
	"fmt"
	"main/internal/model"
	"main/internal/types/dto"
	"main/pkg/engine/controller"
	"main/pkg/helpers"
	"main/pkg/storage"
	"main/server/admin/repository"
	"main/server/admin/view/pages/landing/tabs"
	"net/http"
)


func index(ctx *controller.Context) error {
	var Interface model.Interface
	storage.DB.Order("ID desc").Preload("NewVideos").Last(&Interface)
	return ctx.Html(tabs.NewVideos(Interface.NewVideos))
}

func create(ctx *controller.Context, params *NewVideoDto) error {
	file, err := ctx.FormFile("file")
	if file == nil || err != nil {	
		return ctx.String(
			http.StatusBadRequest, 
			"file in request is nil or could not be extracted : " + err.Error(),
		)
	}

	Upload := helpers.File(file)
	if !Upload.Success {
		fmt.Print("File did not upload -> " + Upload.Message) 
		return ctx.String(http.StatusBadRequest, Upload.Message)
	}

	storage.DB.Model(repository.FindByID[model.Interface](1)).
				Association("NewVideos").
				Append(repository.FindByID[model.Files](Upload.ID))

	return index(ctx)
}

func remove(ctx *controller.Context, ID *dto.ByID) error {
	storage.DB.Model(repository.FindByID[model.Interface](1)).
			   Association("NewVideos").Delete(repository.FindByID[model.Files](ID.ID))

	return index(ctx)
}
