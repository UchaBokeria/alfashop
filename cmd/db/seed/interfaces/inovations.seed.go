package interfaces

import (
	"main/internal/model"
	"main/pkg/storage"
)

var Inovations = []model.Interface_inovations{
	{
		InterfaceID: 1,
		Name:        "USVO",
		Slug:        "usvo",
		Title:       "USVO",
		Url:         "https://www.alfashop.de/de/unternehmen/innovationen/usvo",
		PicID:       GetPointedInt(7),
	},
	{
		InterfaceID: 1,
		Name:        "CleanSynto",
		Slug:        "cleansynto",
		Title:       "Clean Synto",
		Url:         "https://www.alfashop.de/de/unternehmen/innovationen/cleansynto",
		PicID:       GetPointedInt(8),
	},
	{
		InterfaceID: 1,
		Name:        "ATFprofessionalLine",
		Slug:        "atfprofessionalline",
		Title:       "ATF პროფესიონალური",
		Url:         "https://www.alfashop.de/de/unternehmen/innovationen/atf-professional-line",
		PicID:       GetPointedInt(9),
	},
	{
		InterfaceID: 1,
		Name:        "BoxInBag",
		Slug:        "boxinbag",
		Title:       "ჩანთა ყუთში",
		Url:         "https://www.alfashop.de/de/unternehmen/innovationen/bag-in-box",
		PicID:       GetPointedInt(10),
	},
}

func Inovation() {
	for _, row := range Inovations {
		storage.DB.Create(&row)
	}
}
