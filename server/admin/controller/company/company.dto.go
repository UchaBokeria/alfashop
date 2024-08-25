package company

type CompanyCreateDto struct {
	ID              int      	`param:"id"`
	Name			string		`form:"name"`
	Path			string		`form:"path"`
	Body			string		`form:"body"`
}

type CompanyUpdateDto struct {
	ID              int      	`param:"id"`
	CompanyCreateDto
}
