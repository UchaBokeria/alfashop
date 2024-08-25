package schemas

type Links struct {
	Name string
	Slug string
	Path string
}

var Routes = []Links{
	{Path: "/products?category=14", Name: "პროდუქტები", Slug: "products"},
	{Path: "/branches", Name: "კატალოგი", Slug: "branches"},
	{Path: "/posts", Name: "სიახლეები", Slug: "posts"},
	{Path: "/company/about", Name: "ჩვენ შესახებ", Slug: "about"},
	{Path: "/company/static/faq", Name: "დახმარება", Slug: "faq"},
}
