package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"goBank/api"
	db "goBank/db/sqlc"
	"log"
)

const (
	dbSource      = "postgresql://admin:admin@123@localhost:5432/go-bank-postgres?sslmode=disable"
	serverAddress = "localhost:8082"
)

func main() {
	var err error
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
