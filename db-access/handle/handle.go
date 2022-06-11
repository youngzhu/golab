package handle

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func OpenWithString() (*sql.DB, error) {
	conn := "root:root@tcp(localhost:3306)/golab"

	return sql.Open("mysql", conn)
}

func OpenWithProperties() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "golab",
	}

	return sql.Open("mysql", cfg.FormatDSN())
}

func OpenWithConnector() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "golab",
	}

	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return sql.OpenDB(connector), nil
}
