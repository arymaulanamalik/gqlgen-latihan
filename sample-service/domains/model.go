package domains

type Cart struct {
	Items []CartItem `json:"items"`
}

type CartItem struct {
	Qty     string  `json:"qty"`
	Notes   string  `json:"notes"`
	Product Product `json:"product"`
}

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
