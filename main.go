package db

import (
    "database/sql"
    "log"
    "time"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    dsn := "host=db port=5432 user=postgres password=postgres dbname=projectdb sslmode=disable"
    var err error

    for i := 0; i < 10; i++ {
        DB, err = sql.Open("postgres", dsn)
        if err == nil {
            err = DB.Ping()
        }
        if err == nil {
            break
        }
        log.Println("Waiting for Postgres to be ready...")
        time.Sleep(2 * time.Second)
    }

    if err != nil {
        log.Fatalf("Cannot connect to Postgres: %v", err)
    }

    log.Println("Connected to Postgres")
}
