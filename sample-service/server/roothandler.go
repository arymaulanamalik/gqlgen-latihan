package server

import (
	"github.com/arymaulanamalik/gqlgen-latihan/sample-service/domains"
)

type RootHandler struct {
	CartHandler domains.Handler
}
