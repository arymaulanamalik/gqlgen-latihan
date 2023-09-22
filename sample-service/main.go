package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arymaulanamalik/gqlgen-latihan/sample-service/config"
	"github.com/arymaulanamalik/gqlgen-latihan/sample-service/domains"
	"github.com/arymaulanamalik/gqlgen-latihan/sample-service/graph"
	"github.com/arymaulanamalik/gqlgen-latihan/sample-service/server"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewConfig()
	cartSvc := domains.NewService(*cfg)
	cartHandler := domains.NewHandler(cartSvc)

	rootHandler := server.RootHandler{CartHandler: cartHandler}
	gqlResolver := &graph.Resolver{CartService: cartSvc}

	r := server.SetupRouter(rootHandler, gqlResolver)
	serve(r, *cfg)
}

func serve(r *chi.Mux, cfg config.Config) {
	addr := cfg.Address + ":" + cfg.Port
	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		logrus.WithFields(logrus.Fields{
			"addr": cfg.Address,
			"port": cfg.Port,
		}).Info("Starting HTTP server")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server forced to shutdown: %v\n", err)
	}

	logrus.Info("Server Stopped")
}
