package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	// cfg := mysql.Config{
	// 	User:                 "root",
	// 	Passwd:               "root",
	// 	Net:                  "tcp",
	// 	Addr:                 "localhost:5432",
	// 	DBName:               "mytestdb",
	// 	AllowNativePasswords: true,
	// }

	connStr := "user=root password=root dbname=mytestdb host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return db, nil

}
