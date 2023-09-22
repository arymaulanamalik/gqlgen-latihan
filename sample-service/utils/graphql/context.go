package graphql

import (
	"context"
	"net/http"
)

type ctxKey string

const (
	requestHeaderCtx = ctxKey("reqHeader")
)

func WithReqHeader(ctx context.Context, header http.Header) context.Context {
	return context.WithValue(ctx, requestHeaderCtx, header)
}

func GetReqHeader(ctx context.Context) http.Header {
	val, ok := ctx.Value(requestHeaderCtx).(http.Header)
	if !ok {
		val = http.Header{}
	}

	return val
}
