package interfaces

import (
	"fmt"
	"main/internal/model"
	"main/pkg/storage"
)

var Interfaces = []model.Interface{
	{
		Ver:  1,
		Name: "Configuration",
		Slug: "configuration",
		Presentation: `<h2 class="text-4xl text-center font-semibold font-nino text-primary">
            კეთილი იყოს თქვენი მობრძანება
            <b class="font-bold text-5xl font-poppins">RAVENOL</b>-ში
        </h2>

        <p class="w-[80%] ml-[10%] text-center text-primary py-10  font-arial">როგორც ავტომობილების პრემიუმ საპოხი მასალების და სითხეების მწარმოებელი და მიმწოდებელი, ჩვენ გთავაზობთ გადაწყვეტას ყველა საჭიროებისთვის RAVENOL-ის პროდუქტებით</p>
    
        <div class="absolute w-[100vw] ml-[-200px] bg-[#D7E3F4] px-[200px] py-[3rem] flex flex-wrap  justify-start items-start ">
            <div class="w-[60%] flex justify-start items-start">
                <video  class="w-full"
                        src="https://www.ravenol.de/storage/app/media/video/RAVENOL_Corporate_Film.de.mp4"
                        muted
                        controls="controls" 
                        autoplay="autoplay"/>
            </div>

            <div class="w-[40%] h-full px-5 flex flex-col gap-5 items-start justify-start">
                <h2 class="font-semibold text-3xl text-primary font-arial">ჩვენ გვიყვარს, რასაც ვაკეთებთ</h2>
                <p class="w-full text-primary ">
                ჩვენი განვითარების, წარმოების და გაყიდვების საქმიანობა მუდმივად ფართოვდება 1946 წლიდან. ზრდა მოითხოვს მოდერნიზაციას და სტრუქტურას ყველა დონეზე. ჩვენ ამას წარმატებით ვახორციელებთ და ყოველდღიურად ვსწავლობთ ახალ ნივთებს ჩვენი მაღალი ხარისხის სტანდარტების უზრუნველსაყოფად.
                </p>
            </div>
        </div>`,
	},
}

func Populate() {
	Language()

	for _, row := range Interfaces {
		storage.DB.Create(&row)

		storage.DB.
			Model(&row).
			Association("Products").
			Append(
				FindByID[model.Products](1),
				FindByID[model.Products](2),
				FindByID[model.Products](3),
				FindByID[model.Products](4),
				FindByID[model.Products](5),
				FindByID[model.Products](6),
				FindByID[model.Products](7),
			)

		storage.DB.
			Model(&row).
			Association("News").
			Append(
				FindByID[model.Posts](4),
				FindByID[model.Posts](3),
				FindByID[model.Posts](2),
				FindByID[model.Posts](1),
			)

		storage.DB.
			Model(&row).
			Association("Events").
			Append(
				FindByID[model.Posts](12),
				FindByID[model.Posts](13),
				FindByID[model.Posts](14),
				FindByID[model.Posts](15),
			)

		storage.DB.
			Model(&row).
			Association("NewVideos").
			Append(
				FindByID[model.Files](14),
				FindByID[model.Files](14),
			)

	}
}

func FindByID[T any](id int) *T {
	modeler := new(T)
	if err := storage.DB.Where("id = ?", id).First(modeler).Error; err != nil {
		// return err
		fmt.Print("Error finding existing assotiation ", modeler, " | with id -> ", id)
	}

	return modeler
}
