package db

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB

func InitializeDBConnection() {
    db, err := sqlx.Open("mysql", "DenysShpak:ПАРОЛЬ@tcp(127.0.0.1:3306)/SyncMe?parseTime=true")
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
