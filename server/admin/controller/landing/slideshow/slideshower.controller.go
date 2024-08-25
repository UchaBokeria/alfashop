package slideshow

import (
	"fmt"
	"main/internal/model"
	"main/internal/types/dto"
	"main/pkg/engine/controller"
	uploader "main/pkg/helpers"
	"main/pkg/storage"
	"main/server/admin/view/pages/landing/tabs"
	"net/http"
	"strconv"
)

func index(ctx *controller.Context) error {
	var Slides []model.Interface_slideShow
	storage.DB.Order("interface_slide_shows.index ASC").Preload("Pic").Find(&Slides)

	return ctx.Html(tabs.Slideshow(Slides))
}


func Slideshow(ctx *controller.Context, Body *SlideshowDto) error {
	ID, _ := strconv.Atoi(Body.ID)
	var slide model.Interface_slideShow

	match := storage.DB.Last(&slide, uint(ID))
	if match.Error != nil {
		fmt.Print("No slide as ", uint(ID))
		return ctx.String(http.StatusNotFound, "")
	}

	Parameters := model.Interface_slideShow{
		Name: Body.Name,
		Slogan: Body.Slogan,
		Desc: Body.Desc,
		Url: Body.Url,
	}

	file, err := ctx.FormFile("Pic")
	if file != nil && err == nil {	
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print(Upload.Message) 
			return ctx.String(http.StatusBadRequest, Upload.Message)
		}
		Parameters.PicID = &Upload.ID
	}

	result := storage.DB.Model(slide).Updates(&Parameters)
	if result.Error != nil { return ctx.String(http.StatusBadRequest, "") }

	var Slides []model.Interface_slideShow
	storage.DB.Order("interface_slide_shows.index ASC").Preload("Pic").Find(&Slides)

	return ctx.Html(tabs.Slideshow(Slides))
}

func SlideshowNew(ctx *controller.Context, Body *SlideshowDto) error {
	var lastSlide model.Interface_slideShow
	storage.DB.Last(&lastSlide)

	Parameters := model.Interface_slideShow{
		InterfaceID: 1,
		Name: Body.Name,
		Slogan: Body.Slogan,
		Desc: Body.Desc,
		Url: Body.Url,
		Index: (lastSlide.Index + 1),
		TypeID: 1,
	}

	file, err := ctx.FormFile("Pic")
	if file != nil && err == nil {	
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print("File did not upload -> " + Upload.Message) 
			// return ctx.String(http.StatusBadRequest, Upload.Message)
		} else {
			Parameters.PicID = &Upload.ID
		}
	}


	result := storage.DB.Create(&Parameters)
	if result.Error != nil { return ctx.String(http.StatusBadRequest, result.Error.Error()) }

	var Slides []model.Interface_slideShow
	storage.DB.Order("interface_slide_shows.index ASC").Preload("Pic").Find(&Slides)

	return ctx.Html(tabs.Slideshow(Slides))
}

func SlideRemove(ctx *controller.Context, ID *dto.ByID) error {
	var Slide model.Interface_slideShow

	storage.DB.First(&Slide, ID)
	storage.DB.Delete(&Slide)

	var Slides []model.Interface_slideShow
	storage.DB.Order("interface_slide_shows.index ASC").Preload("Pic").Find(&Slides)
	return ctx.Html(tabs.Slideshow(Slides))
}
