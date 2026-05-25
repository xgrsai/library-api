package router

import (
    "net/http"
    "library-api/handlers"
)

func Setup() {
    http.HandleFunc("/", handlers.PrintMsg)
    http.HandleFunc("/books", handlers.GetBooks)
}