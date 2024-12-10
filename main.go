package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/irwanhub2016/go-crud/config"
    "github.com/irwanhub2016/go-crud/routes"
)

func main() {
    config.Connect()
    router := routes.SetupRoutes()

    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
