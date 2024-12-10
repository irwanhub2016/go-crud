package config

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    var err error
    connStr := "user=irwan password=Lenteng@123 dbname=golang_crud sslmode=disable"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }

    if err = DB.Ping(); err != nil {
        panic(fmt.Sprintf("Failed to ping database: %v", err))
    }

    fmt.Println("Connected to the database!")
}
