package app

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnection() (*sql.DB) {
	LoadDotEnv("db.env")

	Driver := "mysql"
	User := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connectionUri := User + ":" + Password + "@tcp(127.0.0.1:3306)/" + dbName

	connection, err := sql.Open(Driver, connectionUri)

	if err != nil {
		fmt.Println("Something went wrong making the connection")
		panic(err.Error)
	}

	fmt.Println("DB connection stablished")
	return connection
}