package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    log.Printf("Got a request")
    // vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "This is v2\n")
}


func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    // r.HandleFunc("/products", ProductsHandler)
    // r.HandleFunc("/articles", ArticlesHandler)
    http.Handle("/", r)
    srv := &http.Server{
        Handler:      r,
        Addr:         ":8080",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
    log.Print("Starting the server")
    log.Fatal(srv.ListenAndServe())
}