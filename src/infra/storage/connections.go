package storage

import (
	"context"
	"fmt"

	"github.com/castmetal/golang-api-boilerplate/src/config"
	"github.com/castmetal/golang-api-boilerplate/src/domains/common/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresConnection(ctx context.Context, config config.EnvStruct) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(ctx, config.DB.URL)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	logger.Info(ctx, "connected to db")

	return db, nil
}
