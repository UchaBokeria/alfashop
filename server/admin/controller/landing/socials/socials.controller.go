package socials

import (
	"main/internal/model"
	"main/pkg/engine/controller"
	"main/pkg/storage"
	"main/server/admin/view/pages/landing/tabs"
	"net/http"
)

func index(ctx *controller.Context) error {
	var SocialMedia []model.Social_media
	storage.DB.Find(&SocialMedia)

	return ctx.Html(tabs.Socials(SocialMedia))
}

func Socials(ctx *controller.Context, Body *SocialDto) error {

	var SocialMediaFacebook model.Social_media
	resultFacebook := storage.DB.Model(SocialMediaFacebook).
						 Where(&model.Social_media{Name: "Facebook"}).
						 Updates(&model.Social_media{Url: Body.Facebook})
						 
	if resultFacebook.Error != nil { return ctx.String(http.StatusBadRequest, "Facebook") }

	var SocialMediaInstagram model.Social_media
	resultInstagram := storage.DB.Model(SocialMediaInstagram).
						 Where(&model.Social_media{Name: "Instagram"}).
						 Updates(&model.Social_media{Url: Body.Instagram})
						 
	if resultInstagram.Error != nil { return ctx.String(http.StatusBadRequest, "Instagram") }

	var SocialMediaTwitter model.Social_media
	resultTwitter := storage.DB.Model(SocialMediaTwitter).
						 Where(&model.Social_media{Name: "Twitter"}).
						 Updates(&model.Social_media{Url: Body.Twitter})
						 
	if resultTwitter.Error != nil { return ctx.String(http.StatusBadRequest, "Twitter") }

	var SocialMediaYouTube model.Social_media
	resultYouTube := storage.DB.Model(SocialMediaYouTube).
						 Where(&model.Social_media{Name: "YouTube"}).
						 Updates(&model.Social_media{Url: Body.YouTube})

	if resultYouTube.Error != nil { return ctx.String(http.StatusBadRequest, "YouTube") }

	var SocialMedia []model.Social_media
	storage.DB.Find(&SocialMedia)
	
	return ctx.Html(tabs.Socials(SocialMedia))
}
