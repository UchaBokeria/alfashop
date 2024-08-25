package contact

import (
	"main/internal/model"
	"main/pkg/engine/controller"
	"main/pkg/storage"
	"main/server/admin/view/pages/landing/tabs"
	"net/http"
)

func index(ctx *controller.Context) error {
	var Contact model.Interface_contact
	storage.DB.Last(&Contact)

	return ctx.Html(tabs.Contact(Contact))
}

func contact(ctx *controller.Context, Body *ContactDto) error {
	var Contact model.Interface_contact

	storage.DB.Last(&Contact)

	Contact.Email = Body.Email
	Contact.Phone = Body.Phone
	Contact.Location = Body.Location
	Contact.LocationLink = Body.LocationLink
	Contact.LocationIframe = Body.LocationIframe
	result := storage.DB.Save(&Contact)

	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, "N")
	}

	return ctx.Html(tabs.Contact(Contact))
}