package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgresql://admin:admin@123@localhost:5432/go-bank-postgres?sslmode=disable"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
