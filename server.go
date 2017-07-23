package main

import (
    "encoding/json"
    "net/http"
)

type Book struct {
    Id int `json:"id"`
    Title string `json:"title"`
    Author string `json:"author"`
}

type Books []Book

func main() {
    http.HandleFunc("/", ShowBooks)
    http.ListenAndServe(":8080", nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
    books := Books{
        Book{1, "The Science", "Bill"},
        Book{2, "The Logic of Universe", "Stephen Halks"},
        Book{3, "Start Small, Stay Small", "Any Author"},
    }

    js, err := json.Marshal(books)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
	return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
