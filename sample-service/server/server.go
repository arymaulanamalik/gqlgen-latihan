package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/arymaulanamalik/gqlgen-latihan/sample-service/graph"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(rootHandler RootHandler, gqlResolver *graph.Resolver) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/carts", rootHandler.CartHandler.CartList)

	gqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: gqlResolver}))

	r.Get("/playground", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", ContextWithReqHeader(gqlServer))

	return r
}
