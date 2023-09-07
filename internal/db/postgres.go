package db

import (
	"fmt"
	"github.com/jackc/pgx"
	"time"
)

type DB struct {
	db *pgx.Conn
}

func NewClient(ctx context.Context, username, password, host, port, database string) {
	posgresURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)

	maxAttempts := 5

	for maxAttempts > 0 {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		connect, err := pgx.Connect()
	}
}
