package faq

import (
	"fmt"
	"main/internal/model"
	"main/internal/types/dto"
	"main/pkg/engine/controller"
	"main/pkg/storage"
	"main/server/admin/view/pages/landing/tabs"
	"net/http"
	"strconv"
)

func index(ctx *controller.Context) error {
	
	var Faqs []model.Faq
	storage.DB.Order("created_at desc").Find(&Faqs)
	return ctx.Html(tabs.Faq(Faqs))
}


func create(ctx *controller.Context, Body *FaqDto) error {

	result := storage.DB.Create(&model.Faq{
		Name: Body.Name,
		Answer: Body.Answer,
		Question: Body.Question,
	})
						 
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Faqs []model.Faq
	storage.DB.Order("created_at desc").Find(&Faqs)
	return ctx.Html(tabs.Faq(Faqs))
}

func update(ctx *controller.Context, Body *FaqDto) error {
	var faq model.Faq
	ID, _ := strconv.Atoi(Body.ID)

	match := storage.DB.Last(&faq, uint(ID))
	if match.Error != nil {
		fmt.Print("No Faq as ", uint(ID))
		return ctx.String(http.StatusNotFound, "")
	}

	result := storage.DB.Model(faq).Updates(&model.Faq{
		Name: Body.Name,
		Answer: Body.Answer,
		Question: Body.Question,
	})

	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Faqs []model.Faq
	storage.DB.Order("created_at desc").Find(&Faqs)
	return ctx.Html(tabs.Faq(Faqs))
}

func remove(ctx *controller.Context, ID *dto.ByID) error {
	var Faq model.Faq

	storage.DB.First(&Faq, ID)
	result := storage.DB.Delete(&Faq)
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Faqs []model.Faq
	storage.DB.Order("created_at desc").Find(&Faqs)
	return ctx.Html(tabs.Faq(Faqs))
}
