package hello

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", hello)
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello, world</h1>")
}