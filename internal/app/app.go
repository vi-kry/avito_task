package app

import (
	"context"
	"fmt"
	"log/slog"

	"avito_task/internal/config"
	"avito_task/internal/configure/database/postgres"
	"avito_task/pkg/logger"
)

func Run() {
	cfg := config.InitConfig()

	log := logger.SetupLogger(cfg.Env)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := setupDatabase(ctx, log, &cfg)


	// ------- repo
	repoTender :=
}

func setupDatabase(ctx context.Context, log *slog.Logger, cfg *config.Config) *postgres.PgDB {
	postgresDB, err := postgres.NewPostgresDB(ctx, log, postgresDSN(&cfg.Postgres))
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	return postgresDB
}

func postgresDSN(psqlCfg *config.PostgresConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		psqlCfg.Host, psqlCfg.Port, psqlCfg.Username, psqlCfg.Password, psqlCfg.DbName, psqlCfg.SslMode)
}
