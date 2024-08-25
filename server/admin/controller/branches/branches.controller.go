package branches

import (
	"main/internal/model"
	"main/internal/types/dto"
	"main/pkg/engine/controller"
	"main/pkg/helpers"
	"main/pkg/storage"
	"main/server/admin/repository"
	"main/server/admin/view/pages/branches"
)

func index(ctx *controller.Context) error {
	var Branches []model.Branches
	storage.DB.Order("created_at desc").Preload("District.City").Find(&Branches)
	return ctx.Html(branches.Page(Branches))
}

func unique(ctx *controller.Context, ID *dto.ByID) error {
	var Branch model.Branches
	storage.DB.
		Preload("District.City").
		Preload("Shifts").
		First(&Branch, ID)

	var Shifts []model.Branch_shifts
	storage.DB.Find(&Shifts)

	var Cities []model.Cities
	storage.DB.Preload("District").Find(&Cities)

	var Districts []model.Districts
	storage.DB.Preload("District").Find(&Districts)

	ShiftsSelect, _ := helpers.ToSelect(Shifts, "ID", "Name")
	CitiesSelect, _ := helpers.ToSelect(Cities, "ID", "Display_name")
	DistrictsSelect, _ := helpers.ToSelect(Districts, "ID", "Display_name")

	return ctx.Html(branches.Update(Branch, CitiesSelect, ShiftsSelect, DistrictsSelect))
}

func create(ctx *controller.Context, params *CreateBranchesDto) error {
	repository.Create[model.Branches](&model.Branches{
		Map: params.Map,
		Name: params.Name,
		PhoneNumber: params.PhoneNumber,
		DistrictID: params.DistrictID,
	})

	return index(ctx)
}

func update(ctx *controller.Context, params *UpdateBranchesDto) error {
	repository.Update[model.Branches](params.ID, &model.Branches{
		Map: params.Map,
		Name: params.Name,
		PhoneNumber: params.PhoneNumber,
		DistrictID: params.DistrictID,
	})

	return index(ctx)
}

func remove(ctx *controller.Context, ID *dto.ByID) error {
	repository.Delete[model.Branches](ID)
	return index(ctx)
}

func properties(ctx *controller.Context) error {
	var Shifts []model.Branch_shifts
	storage.DB.Find(&Shifts)

	var Cities []model.Cities
	storage.DB.Preload("District").Find(&Cities)

	ShiftsSelect, _ := helpers.ToSelect(Shifts, "ID", "Name")
	CitiesSelect, _ := helpers.ToSelect(Cities, "ID", "Display_name")

	return ctx.Html(branches.Create(CitiesSelect, ShiftsSelect))
}

func districts(ctx *controller.Context, city *DistrictsDto) error{
	var Districts []model.Districts
	storage.DB.Where(&model.Districts{CityID: city.ID}).Preload("District").Find(&Districts)
	DistrictSelect, _ := helpers.ToSelect(Districts, "ID", "Display_name")

	if city.Indicator == "create" {
		return ctx.Html(branches.CreateDistricts(DistrictSelect))
	} else {
		return ctx.Html(branches.UpdateDistricts(DistrictSelect))
	}
}
