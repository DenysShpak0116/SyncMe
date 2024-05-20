package db

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB

func InitializeDBConnection() {
    db, err := sqlx.Open("mysql", "SyncMeAdmin:Smad_mysql123@tcp(syncme.mysql.database.azure.com:3306)/syncme?tls=skip-verify")

    if err != nil {
        panic(err.Error())
    }

    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Success")
    DBClient = db
}
