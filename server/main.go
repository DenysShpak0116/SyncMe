package main

import (
    "fmt"
    "log"
    "net/http"
    "server/controllers"
    "server/db"
)

func main() {
    db.InitializeDBConnection()
    http.HandleFunc("/", handler)
    http.HandleFunc("/CreateComment", controllers.CreateCommentHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World!")
}
