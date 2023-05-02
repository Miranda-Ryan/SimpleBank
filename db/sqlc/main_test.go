package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	driverName = "postgres"
	dbSource   = "postgres://root:secret@127.0.0.1:5432/simplebank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(driverName, dbSource)
	if err != nil {
		log.Fatal("Could not connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
