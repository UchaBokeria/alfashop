package products

type ProductsDto struct {
	ID               int      		`param:"id"`
	Name             string			`form:"Name"`
    // Slug             string			`form:"Slug"`
    Public           bool 			`form:"Public"`
    CategoryID       int			`form:"CategoryID"`
    Description      string			`form:"Description"`
    DescriptionHtml  string			`form:"DescriptionHtml"`
    // Pics             [][]byte	    `form:"Pics"` 
    // TechnicalSheet   []byte   		`form:"TechnicalSheet"`
}

type ProductPropertyDto struct {
    ProductsID      int      		`param:"ProductsID"`
    Characteristics string			`form:"Characteristics"`
    Unit            string			`form:"Unit"`
    Data            string			`form:"Data"`
    Examination     string			`form:"Examination"`
}
