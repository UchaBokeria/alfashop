package branches

import "main/internal/types/dto"

type CreateBranchesDto struct {
	Map 		 string 		`form:"Map"`
	Name 		 string 		`form:"Name"`
	DistrictID	 int 			`form:"DistrictID"`
	PhoneNumber  string 		`form:"PhoneNumber"`
}

type UpdateBranchesDto struct {
	ID              int      	`param:"id"`
	CreateBranchesDto
}

type DistrictsDto struct {
	dto.ByID
	Indicator 	string			`param:"indicator"`
}