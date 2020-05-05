package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// main is the entry point of any go application.
func main() {
	// Create a new gorilla router.
	r :=mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to to the go-docker API.")
	})


	// TODO: implement TLS 
	// At this stage we only accept http requests.
	fmt.Println("Server listening on port 80")
	http.ListenAndServe(":80", r)
}