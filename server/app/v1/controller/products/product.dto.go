package products

type FiltersQuery struct {
	Category 	 	int 		`query:"category"`
	Searcher        string      `query:"searcher"`
}

type ProductDetail struct {
	ID              int      	`param:"id"`
}