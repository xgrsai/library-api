package handlers

import (
	"net/http"
	"encoding/json"
	"library-api/storage"
	"fmt"
)

func PrintMsg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Book Library API is running")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(storage.Books)
}