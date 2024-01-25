package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler functions for routes
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact Page")
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Services Page")
}

// Router Set up
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/about", aboutHandler)
	r.HandleFunc("/contact", contactHandler)
	r.HandleFunc("/services", servicesHandler)

	// Server Start and Setup
	http.Handle("/", r)

	port := "8080"
	addr := fmt.Sprintf(":%s", port)

	// Server print info
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
