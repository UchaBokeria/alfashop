package interfaces

import (
	"main/internal/model"
	"main/pkg/storage"
)

var Languages = []model.Languages{
	{
		InterfaceID: 1,
		IconID:      45,
		Name:        "ქართული",
		Slug:        "ge",
	},
	{
		InterfaceID: 1,
		IconID:      46,
		Name:        "English",
		Slug:        "en",
	},
	{
		InterfaceID: 1,
		IconID:      47,
		Name:        "Русский",
		Slug:        "ru",
	},
}

var LanguageDictionary = []model.LanguageDictionary{
	{
		LanguageID: 1,
		Key:        "NewsPageTitle",
		Value:      "სიახლეები",
	},
	{
		LanguageID: 2,
		Key:        "NewsPageTitle",
		Value:      "News",
	},
	{
		LanguageID: 3,
		Key:        "NewsPageTitle",
		Value:      "სიახლეები",
	},
}

func Language() {
	for _, row := range Languages {
		storage.DB.Create(&row)
	}
	for _, row := range LanguageDictionary {
		storage.DB.Create(&row)
	}
}
