package cart

type CartDto struct {
	IDs string `query:"ids"`
}

type CheckoutDto struct {
	Items []struct {
		ID       int
		Quantity int
	} `json:"items"`
}
