package server

import (
	"net/http"

	gqlUtils "github.com/arymaulanamalik/gqlgen-latihan/sample-service/utils/graphql"
)

func ContextWithReqHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := gqlUtils.WithReqHeader(r.Context(), r.Header)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}
