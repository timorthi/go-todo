package main

import (
    "io"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        io.WriteString(w, "Hello, world!\n")
    })
    http.ListenAndServe(":7000", nil)
}
