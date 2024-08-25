package contact

type ContactDto struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	Location string `json:"location"`
	LocationLink string `json:"locationLink"`
	LocationIframe string `json:"locationIframe"`
}
