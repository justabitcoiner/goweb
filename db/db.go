package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func Connect() {
	var err error
	conn, err = pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/goweb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}

func Disconnect() {
	conn.Close(context.Background())
}
