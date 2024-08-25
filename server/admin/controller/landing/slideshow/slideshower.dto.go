package slideshow

type SlideshowDto struct {
	ID    string `param:"id"`
	Name  string `form:"Name"`
	Slogan string `form:"Slogan"`
	Desc   string `form:"Desc"`
	Url   string `form:"Url"`
	Pic   []byte `form:"Pic"`
}
