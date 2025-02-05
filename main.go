package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Render! 🚀")
}

func main() {
    port := os.Getenv("PORT") // Render provides a dynamic port

    if port == "" {
        port = "8080" // Default fallback
    }

    http.HandleFunc("/", handler)
    fmt.Println("Server running on port " + port)
    http.ListenAndServe(":"+port, nil)
}

