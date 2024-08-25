package interfaces

import (
	"os"

	"main/internal/model"
	"main/pkg/storage"
)

func getFile(path string) string {
	dat, err := os.ReadFile(path)
    if err != nil {
		// fmt.Printf("files error: %T", err)
        panic(err)
    }
	return string(dat)
}

var Abouts = []model.Interface_about{
	{
		InterfaceID: 1,
		Ver:         1,
		Body:        getFile("cmd/db/seed/interfaces/about.html"),
	},
} 

func About() {
	for _, row := range Abouts { storage.DB.Create(&row) }
}