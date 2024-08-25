package landing

import (
	"log"

	"github.com/a-h/templ"
	"gorm.io/gorm"

	"main/internal/model"
	"main/pkg/engine/controller"
	mailer "main/pkg/services/mail"
	"main/pkg/storage"
	view "main/server/app/v1/view/pages/landing"
)

func index(ctx *controller.Context) error {
	var Interface model.Interface

	storage.DB.
		Preload("SlideShow", func(db *gorm.DB) *gorm.DB {
			return db.Order("interface_slide_shows.index ASC").Preload("Pic")
		}).
		Preload("Inovations.Pic").
		Preload("Products.Pics").
		Preload("Products.Category").
		Preload("News.Thumbnail").
		Preload("Events.Thumbnail").
		Preload("NewVideos").
		Preload("Contact").
		Preload("About").
		Preload("SocialMedia").
		Last(&Interface)

	return ctx.Html(view.Landing(Interface))
}

func subscribe(ctx *controller.Context) error {
	Result, _ := mailer.Send(mailer.Config{
		To: "ucha2bokeria@gmail.com",
		Subject: "New Subsribe",
		Body: "მადლობა გამოწერისთვის, იხილეთ სიახლეები ჩვენს ვებ გვერდზე და მიიღეთ ექსკლუზიური სიახლეები ელ ფოსტის საშუალებით",
	})

	if Result {
		log.Print("Sent")
	}
	return ctx.Html(templ.NopComponent)
}