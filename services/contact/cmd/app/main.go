package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"webServiceGolang/pkg/store/postgres"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer func(Pool *pgx.Conn, ctx context.Context) {
		err := Pool.Close(ctx)
		if err != nil {

		}
	}(conn.Pool, context.Background())
}
