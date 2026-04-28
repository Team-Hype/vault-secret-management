package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Team-Hype/vault-secret-management/internal/config"
	"github.com/Team-Hype/vault-secret-management/internal/postgres"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	configPath := flag.String("config", "/vault/secrets/db.env", "path to config file")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	var cfg config.Config
	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
		slog.Error("read config error", slog.String("error", err.Error()))
		os.Exit(1)
	}

	pool, err := pgxpool.New(ctx, cfg.Postgres.ToDSN())
	if err != nil {
		slog.Error("create postgres pool error", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer pool.Close()

	dbInfo := postgres.NewDBInfo(pool)
	infos, err := dbInfo.Scan(ctx)
	if err != nil {
		slog.Error("db info error", slog.String("error", err.Error()))
		os.Exit(1)
	}

	for _, info := range infos {
		slog.Info("db info", slog.String("key", info.Key), slog.String("value", info.Value))
	}
}
