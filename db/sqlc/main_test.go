package db

import (
	"context"
	"goBank/util"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	conn, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot log config", err)
	}
	testDB, err = pgxpool.New(context.Background(), conn.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
