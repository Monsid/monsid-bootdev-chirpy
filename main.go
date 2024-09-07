package main

import (
    "log"
    "net/http"
)

func main() {
    // Create a new ServeMux
    mux := http.NewServeMux()

    // Create a readiness handler
    readinessHandler := func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    }

    // Register the readiness handler for the /healthz path
    mux.HandleFunc("/healthz", readinessHandler)

    // Create a file server handler
    fileServer := http.FileServer(http.Dir("."))

    // Use the Handle method to register the file server for the /app/ path
    mux.Handle("/app/", http.StripPrefix("/app", fileServer))

    // Create a new Server struct
    server := &http.Server{
        Addr:    ":8080", // Set the address to listen on
        Handler: mux,    // Use the ServeMux as the handler
    }

    // Start the server
    log.Println("Server is starting on port 8080...")
    if err := server.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}