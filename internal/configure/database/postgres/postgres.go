package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgDB struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewPostgresDB(ctx context.Context, log *slog.Logger, dsn string) (*PgDB, error) {
	const op = "postgres.NewPostgresDB"

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	PgDB := &PgDB{
		db:  db,
		log: log,
	}

	if err = PgDB.pingContext(ctx); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return PgDB, nil
}

func (pg *PgDB) GetDB() *pgxpool.Pool {
	return pg.db
}

func (pg *PgDB) Close() {
	pg.db.Close()
}

func (pg *PgDB) pingContext(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	status := "up"
	if err := pg.db.Ping(ctx); err != nil {
		status = "down"
		pg.log.Error("database status", slog.String("status", status))
		return fmt.Errorf("failed to ping database: %w", err)
	}
	pg.log.Info("database status", slog.String("status", status))

	return nil
}
