package migration

import (
	"main/internal/model"
)

var Models = []interface{}{
	&model.Company{},
	&model.Users{},
	&model.Cities{},
	&model.Districts{},
	&model.Branches{},
	&model.Branch_shifts{},

	&model.Category_filters_option{},
	&model.Category_filters{},
	&model.Categories{},

	&model.Chat_status{},
	&model.Chat_type{},
	&model.Chat{},
	&model.Chat_letters{},

	&model.Faq{},

	&model.File_types{},
	&model.Files{},

	&model.Posts_types{},
	&model.Posts{},

	&model.Products{},

	&model.Interface{},
	&model.Interface_slideShow{},
	&model.Interface_inovations{},
	&model.Interface_contact{},
	&model.Interface_about{},
	&model.Social_media{},
	&model.Languages{},
	&model.LanguageDictionary{},
}
