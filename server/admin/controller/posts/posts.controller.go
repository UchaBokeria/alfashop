package posts

import (
	"fmt"
	"main/internal/model"
	"main/internal/types/dto"
	"main/pkg/engine/controller"
	uploader "main/pkg/helpers"

	"main/pkg/storage"
	"main/server/admin/view/pages/posts"
	"net/http"
	"strconv"
)

func index(ctx *controller.Context) error {
	var Postz []model.Posts
	storage.DB.Order("created_at desc").Preload("Thumbnail").Find(&Postz)
	return ctx.Html(posts.Page(Postz))
}

func update(ctx *controller.Context, Body *PostsDto) error {

	if err := ctx.Bind(&Body); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	var Post model.Posts
	ID, _ := strconv.Atoi(Body.ID)

	match := storage.DB.Last(&Post, uint(ID))
	if match.Error != nil {
		fmt.Print("No Faq as ", uint(ID))
		return ctx.String(http.StatusNotFound, "")
	}

	TypeID, _ := strconv.Atoi(Body.TypeID)
	Public := true
	if Body.Public == "false" { Public = false }
	
	Parameters := map[string]interface{}{
		"Title": Body.Title,
		"Body": Body.Body,
		"Public": Public,
		"Url": Body.Url,
		"TypeID": TypeID,
	}

	file, err := ctx.FormFile("Thumbnail")
	if file != nil && err == nil {
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print(Upload.Message) 
			// return ctx.String(http.StatusBadRequest, Upload.Message)
		} else {
			Parameters["ThumbnailID"] = Upload.ID
		}
		
	}

	result := storage.DB.Model(Post).Updates(&Parameters)
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Postz []model.Posts
	storage.DB.Order("created_at desc").Preload("Thumbnail").Find(&Postz)
	return ctx.Html(posts.Page(Postz))
}

func create(ctx *controller.Context, Body *PostsDto) error {

	TypeID, _ := strconv.Atoi(Body.TypeID)
	
	Public := true
	if Body.Public == "false" { Public = false }

	Parameters := model.Posts{
		Title: Body.Title,
		Body: Body.Body,
		Public: Public,
		Url: Body.Url,
		TypeID: TypeID,
	}

	file, err := ctx.FormFile("Thumbnail")
	if file != nil && err == nil {
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print(Upload.Message) 
			return ctx.String(http.StatusBadRequest, Upload.Message)
		}
		
		Parameters.ThumbnailID = Upload.ID
	} else {
		fmt.Print(err.Error()) 
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	result := storage.DB.Create(&Parameters)
						 
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Postz []model.Posts
	storage.DB.Order("created_at desc").Preload("Thumbnail").Find(&Postz)
	return ctx.Html(posts.Page(Postz))
}

func remove(ctx *controller.Context, ID *dto.ByID) error {
	var Post model.Posts

	storage.DB.First(&Post, ID)
	result := storage.DB.Delete(&Post)
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Postz []model.Posts
	storage.DB.Order("created_at desc").Preload("Thumbnail").Find(&Postz)
	return ctx.Html(posts.Page(Postz))
}
