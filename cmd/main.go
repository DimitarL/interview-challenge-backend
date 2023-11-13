package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/DimitarL/rental/internal/api"
	"github.com/DimitarL/rental/internal/service"
	"github.com/DimitarL/rental/internal/store"
	"github.com/caarlos0/env/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	DBConnURL string `env:"DB_CONN_URL" envDefault:"postgres://root:root@localhost:5434/testingwithrentals"`
}

func main() {
	appCtx := context.Background()

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		slog.With(slog.String("error", err.Error())).Error("failed parsing config from env")
		os.Exit(1)
	}

	pool, err := pgxpool.New(appCtx, cfg.DBConnURL)
	if err != nil {
		slog.With(slog.String("error", err.Error())).Error("failed connecting to DB")
		os.Exit(1)
	}

	st := store.NewStore(pool)
	svc := service.NewService(st)

	getRentalHandler := api.NewGetRentalHandler(svc)
	getRentalsHandler := api.NewGetRentalsHandler(svc)
	router := api.NewRouterBuilder(getRentalHandler, getRentalsHandler)

	err = http.ListenAndServe("127.0.0.1:8080", router.Build())
	if err != nil {
		slog.With(slog.String("error", err.Error())).Error("failed to listen and serve")
		os.Exit(1)
	}
}
