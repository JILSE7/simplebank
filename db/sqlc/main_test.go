package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDvriver = "postgres"
	dbSource  = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDvriver, dbSource)

	if err != nil {
		log.Fatal("cannot conncet to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
