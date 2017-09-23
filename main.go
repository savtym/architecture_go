package main

import (
    "net/http"
    "./models/morze"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./views/src")));
    http.HandleFunc("/api/morze", morze.Handler);

    http.ListenAndServe(":8080", nil);
}