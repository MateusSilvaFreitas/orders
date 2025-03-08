package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase() {
	cfg := mysql.Config{
		User: "root",
		Passwd: "root",
		Net: "tcp",
		Addr: "localhost:3306",
		DBName: "orders",
	}

	var err error

	DB, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to database orders!")
}