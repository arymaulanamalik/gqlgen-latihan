package domains

import (
	"context"
	"errors"

	"github.com/arymaulanamalik/gqlgen-latihan/sample-service/config"
)

type Service interface {
	GetCartByID(ctx context.Context, id string) (*Cart, error)
}

type service struct {
	conf config.Config
}

func NewService(conf config.Config) Service {
	return &service{
		conf: conf,
	}
}

func (ths *service) GetCartByID(ctx context.Context, id string) (*Cart, error) {

	result := &Cart{
		// Items: []CartItem{
		// 	{
		// 		Qty:   "3",
		// 		Notes: "ok",
		// 		Product: Product{
		// 			ID:   "1",
		// 			Name: "Product-1",
		// 		},
		// 	},
		// 	{
		// 		Qty:   "2",
		// 		Notes: "ok",
		// 		Product: Product{
		// 			ID:   "2",
		// 			Name: "Product-2",
		// 		},
		// 	},
		// },
	}

	err := errors.New("FAILED SOMETHING")

	return result, err
}
