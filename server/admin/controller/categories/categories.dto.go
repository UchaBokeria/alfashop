package categories

type CategoriesDto struct {
	ID          int      `param:"id"`
	ParentID    int      `form:"ParentID"`
	Name        string   `form:"Name"`
	Public      bool     `form:"Public"`
	Description string   `form:"Description"`
	Icon        [][]byte `form:"Icon"`
}
