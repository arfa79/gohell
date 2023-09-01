package main

import (
    "fmt"
    "net/http"
    "gohell/handlers" // Import the handlers package
)

func main() {
    http.HandleFunc("/", hello.HelloHandler) // Use the new handler

    fmt.Println("Server is listening on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
