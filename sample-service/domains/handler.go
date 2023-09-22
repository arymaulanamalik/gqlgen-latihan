package domains

import (
	"net/http"

	"github.com/arymaulanamalik/gqlgen-latihan/sample-service/utils/response"
)

type Handler interface {
	CartList(w http.ResponseWriter, req *http.Request)
}

type handler struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return &handler{
		svc: svc,
	}
}

func (ths *handler) CartList(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	userID := req.Header.Get("x-user-id")
	if len(userID) == 0 {
		response.Error(w,
			response.WithMessage(response.CodeUnauthorized, "user id not found"),
			http.StatusUnauthorized)
		return
	}

	cart, err := ths.svc.GetCartByID(ctx, userID)
	if err != nil {
		response.Error(w,
			response.WithMessage(response.CodeInternal, "failed to get data"),
			http.StatusInternalServerError)
		return
	}

	response.Success(w, cart, http.StatusOK)
}
