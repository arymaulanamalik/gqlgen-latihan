// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CartListResponse interface {
	IsCartListResponse()
}

type Cart struct {
	Items []*CartItem `json:"items"`
}

func (Cart) IsCartListResponse() {}

type CartItem struct {
	Qty     string   `json:"qty"`
	Notes   string   `json:"notes"`
	Product *Product `json:"product"`
}

type GenericError struct {
	Code    string  `json:"code"`
	Message *string `json:"message,omitempty"`
}

func (GenericError) IsCartListResponse() {}

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
