package company

import (
	"main/internal/model"
	"main/pkg/storage"
	"os"
)


func getFile(path string) string {
	dat, err := os.ReadFile(path)
    if err != nil {
		// fmt.Printf("files error: %T", err)
        panic(err)
    }
	return string(dat)
}

var Seed = []model.Company{
	{
		Path: "about",
		Name: "ჩვენს შესახებ",
		Body: getFile("cmd/db/seed/interfaces/about.html"),
	},
	{
		Path: "terms",
		Name: "წესები და პირობები",
		Body: getFile("cmd/db/seed/interfaces/terms.txt"),
	},
}

func Populate() {
	for _, row := range Seed { storage.DB.Create(&row) }
}
