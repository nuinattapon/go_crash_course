package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbUser   = "dev"
	dbPass   = "Welcome1"
	dbName   = "nui"
	dbHost   = "tcp(192.168.1.6:3306)"
	// dbSource = "dev:Welcome1@tcp(192.168.1.6:3306)/nui?parseTime=true"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	// create database connection string
	dbSource := dbUser + ":" + dbPass + "@" + dbHost + "/" + dbName + "?parseTime=true"
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to the database")
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
