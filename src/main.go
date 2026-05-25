package main

import (
    "fmt"
    "net/http"
    "library-api/router"
)

func main() {
    router.Setup()
    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}