package handlers

import (
    "net/http"
    "encoding/json"
    "library-api/storage"
    "library-api/models"
    "fmt"
)

func PrintMsg(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Book Library API is running")
}

func HandleEndpoint(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // Logic to retrieve data
        w.Header().Set("Content-Type", "application/json") // set json response

        data, err := json.Marshal(storage.Books)
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        w.Write([]byte(data))

    case http.MethodPost:
        // Logic to create or update data
        w.Header().Set("Content-Type", "application/json") // set json response

        var books []models.Book // empty variable - pass slice cause there will  be invalid jssonn if we will create multiple books
        err := json.NewDecoder(r.Body).Decode(&books) // take json?
		if err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		// set id
		for i := range books {
			books[i].ID = len(storage.Books) + i + 1
		}

		storage.Books = append(storage.Books, books...) // ... syntax for all slice elements take and append them all

        data, err := json.Marshal(books)
        if err != nil {
            http.Error(w, "error", 500)
            return
        }

		w.WriteHeader(http.StatusCreated)
        w.Write([]byte(data))

    default:
        // Handle unsupported methods
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}