package posts

type PostsDto struct {
	ID         		string       `param:"id"`
	Title			string       `form:"Title"`
	Body 			string       `form:"Body"`
	Public 			string       `form:"Public"`
	Url 			string       `form:"Url"`
	Thumbnail		[]byte	     `form:"Thumbnail"`
	TypeID 			string       `form:"TypeID"`
}
