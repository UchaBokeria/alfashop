package categories

import (
	"fmt"
	"main/internal/model"
	"main/internal/types/dto"
	"main/pkg/engine/controller"
	"main/pkg/helpers"
	"main/pkg/storage"
	"main/server/admin/repository"
	"main/server/admin/view/pages/categories"
	"main/server/admin/view/shared"
	"net/http"

	"gorm.io/gorm"
)

func index(ctx *controller.Context) error {
	var CategoriesPages []model.Categories
	storage.DB.Where(&model.Categories{Level: 1}).
		Preload("Children", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Icon").Preload("Children", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Icon").Preload("Children", func(db *gorm.DB) *gorm.DB {
					return db.Preload("Icon")
				})
			})
		}).
		Preload("Icon").
		Find(&CategoriesPages)

	return ctx.Html(categories.Page(CategoriesPages))
}

func createForm(ctx *controller.Context, Param *dto.ByID) error {
	return ctx.Html(categories.Create(
		helpers.String(Param.ID),
		[]shared.SelectDataType{
			{Key: "true", Val: "აქტიური"},
			{Key: "false", Val: "დეაქტიური"},
		},
	))
}

func create(ctx *controller.Context, params *CategoriesDto) error {
	NewCategory := &model.Categories{
		Name:        params.Name,
		Public:      params.Public,
		Description: params.Description,
		Children:    []*model.Categories{},
		Level:       1,
	}

	if params.ParentID != 0 {
		Parent := model.Categories{}
		if err := storage.DB.First(&Parent, params.ParentID).Error; err != nil {
			return ctx.String(http.StatusInternalServerError, "Problem with parent -> "+err.Error())
		}
		NewCategory.Level = Parent.Level + 1
		NewCategory.ParentID = helpers.UintToPtInt(Parent.ID)
	}

	if file, err := ctx.FormFile("Icon"); file != nil && err == nil {
		if Upload := helpers.File(file); !Upload.Success {
			fmt.Print("File did not upload -> " + Upload.Message)
			return ctx.String(http.StatusInternalServerError, Upload.Message)
		} else {
			NewCategory.IconID = helpers.GetPointedInt(Upload.ID)
		}
	}

	if err := repository.Create[model.Categories](NewCategory); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return index(ctx)
}

func updateForm(ctx *controller.Context, Param *dto.ByID) error {
	var Category model.Categories

	if err := storage.DB.
		Preload("Children").
		Preload("Parent").
		Preload("Icon").
		First(&Category, Param.ID).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.Html(categories.Update(
		Category,
		[]shared.SelectDataType{
			{Key: "true", Val: "აქტიური"},
			{Key: "false", Val: "დეაქტიური"},
		},
	))
}

func update(ctx *controller.Context, params *CategoriesDto) error {
	UpdatedCategory := map[string]interface{}{
		"Name":        params.Name,
		"Public":      params.Public,
		"Description": params.Description,
		"Children":    []*model.Categories{},
	}

	if file, err := ctx.FormFile("Icon"); file != nil && err == nil {
		if Upload := helpers.File(file); !Upload.Success {
			fmt.Print("File did not upload -> " + Upload.Message)
			return ctx.String(http.StatusInternalServerError, Upload.Message)
		} else {
			UpdatedCategory["IconID"] = helpers.GetPointedInt(Upload.ID)
		}
	}

	helpers.Print(UpdatedCategory)
	if err := repository.Update[model.Categories](params.ID, UpdatedCategory); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return index(ctx)
}

func remove(ctx *controller.Context, ID *dto.ByID) error {
	if err := repository.Delete[model.Categories](ID); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return index(ctx)
}

func status(ctx *controller.Context, Visibility *dto.Visibility) error {
	if err := repository.Update[model.Categories](Visibility.ID, map[string]interface{}{
		"Public": Visibility.Status,
	}); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return index(ctx)
}
