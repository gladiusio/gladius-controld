package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nfeld9807/rest-api/internal/handlers"
	"log"
	"net/http"
	"strings"
)

func main() {
	var DEBUG = true
	var port = ":3001"

	fmt.Println("Starting server at http://localhost" + port)

	// Main Router
	router := mux.NewRouter()
	if DEBUG {
		router.Use(loggingMiddleware)
	}

	// Base API Sub-Routes
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/", handlers.APIHandler)

	// Status Sub-Routes
	statusRouter := apiRouter.PathPrefix("/status").Subrouter()
	statusRouter.HandleFunc("/", handlers.StatusHandler).
		Methods("GET", "PUT").
		Name("status")
	statusRouter.HandleFunc("/tx/{tx:0[xX][0-9a-fA-F]{64}}", handlers.StatusHandler).
		Methods("GET").
		Name("status-tx")

	// Market Sub-Routes
	marketRouter := apiRouter.PathPrefix("/market").Subrouter()
	marketRouter.HandleFunc("/pools", handlers.MarketPoolsHandler).
		Methods("GET").
		Name("market-pools")
	marketRouter.HandleFunc("/pools/create", handlers.MarketPoolsCreateHandler).
		Methods("POST").
		Name("market-pools-create")

	log.Fatal(http.ListenAndServe(port, router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println()
		log.Println(formatRequest(r))
		log.Println()

		next.ServeHTTP(w, r)
	})
}

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string

	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)

	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))

	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}

	// Return the request as a string
	return strings.Join(request, "\n")
}
