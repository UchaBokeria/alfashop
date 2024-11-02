package cart

type CartDto struct {
	IDs string `query:"ids"`
}

type CheckoutDto struct {
	Comment string `json:"comment"`
	Items   []struct {
		ID       int
		Quantity int
	} `json:"items"`
}
