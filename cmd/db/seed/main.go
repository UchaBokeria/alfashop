package main

import (
	"main/cmd/db/seed/branches"
	"main/cmd/db/seed/categories"
	"main/cmd/db/seed/chat"
	"main/cmd/db/seed/company"
	"main/cmd/db/seed/faq"
	"main/cmd/db/seed/files"
	"main/cmd/db/seed/interfaces"
	"main/cmd/db/seed/posts"
	"main/cmd/db/seed/products"
	"main/cmd/db/seed/users"
	"main/pkg/globals"
	"main/pkg/storage"
)

func main() {
	globals.SetupEnvironmentVariables()
	storage.Connect(storage.Default())

	company.Populate()
	users.Populate()

	files.Types()
	files.Populate()

	posts.Types()
	posts.Populate()

	branches.Cities()
	branches.Populate()
	branches.Shifts()

	faq.Populate()

	chat.Populate()

	categories.Populate()

	products.Populate()

	interfaces.Populate()
	interfaces.SocialMedia()
	interfaces.Slideshow()
	interfaces.Inovation()
	interfaces.Contact()
	interfaces.About()
	interfaces.Language()
}
