package postgres

import (
	"context"
	"net/http"

	"relat/internal/platform/config"
	"relat/internal/platform/httperr"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewPostgres() (*Postgres, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, config.EnvDatabaseUrl)
	if err != nil {
		return nil, httperr.NewHttpError("error initializing postgres pool", http.StatusInternalServerError)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, httperr.NewHttpError("error establishing postgres connection", http.StatusInternalServerError)
	}

	return &Postgres{ctx, pool}, nil
}

func (p *Postgres) Pool() *pgxpool.Pool {
	return p.pool
}
