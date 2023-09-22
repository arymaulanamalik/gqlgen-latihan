package graph

import "github.com/arymaulanamalik/gqlgen-latihan/sample-service/domains"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CartService domains.Service
}
