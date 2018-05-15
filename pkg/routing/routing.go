package routing

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gladiusio/gladius-controld/pkg/routing/handlers"
	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	STATIC_DIR = "/static/"
	PORT       = "3001"
	DEBUG      = true
)

func Start() {
	fmt.Println("Starting API at http://localhost:" + PORT)

	// Main Router
	router := mux.NewRouter()
	if DEBUG {
		router.Use(loggingMiddleware)
	}

	// Base API Sub-Routes
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(responseMiddleware)
	apiRouter.HandleFunc("/manager", handlers.APIHandler)
	apiRouter.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)

	// Key Management
	walletRouter := apiRouter.PathPrefix("/keystore").Subrouter()
	walletRouter.HandleFunc("/wallet/create", handlers.KeystoreCreationHandler).
		Methods("POST")
	walletRouter.HandleFunc("/wallets", handlers.KeystoreWalletsRetrievalHandler).
		Methods("GET")
	walletRouter.HandleFunc("/wallet/{index:[0-9]*}", handlers.KeystoreWalletRetrievalHandler).
		Methods("GET")
	walletRouter.HandleFunc("/wallet/{index:[0-9]*}/open", handlers.KeystoreWalletOpenHandler).
		Methods("POST")
	walletRouter.HandleFunc("/pgp/create", handlers.KeystorePGPCreationHandler).
		Methods("POST")

	// Status Sub-Routes
	statusRouter := apiRouter.PathPrefix("/status").Subrouter()
	statusRouter.HandleFunc("/", handlers.StatusHandler).
		Methods("GET", "PUT").
		Name("status")
	statusRouter.HandleFunc("/tx/{tx:0[xX][0-9a-fA-F]{64}}", handlers.StatusTxHandler).
		Methods("GET").
		Name("status-tx")

	// Node Sub-Routes
	nodeRouter := apiRouter.PathPrefix("/node").Subrouter()
	// Retrieve owned Node if available
	nodeRouter.HandleFunc("/", handlers.NodeFactoryNodeAddressHandler).
		Methods("GET")
	// Node for address
	nodeRouter.HandleFunc("/{nodeAddress:0[xX][0-9a-fA-F]{40}}", nil)
	// Node Creation
	nodeRouter.HandleFunc("/create", handlers.NodeFactoryCreateNodeHandler).
		Methods("POST")
	// Node Data
	nodeRouter.HandleFunc("/{nodeAddress:0[xX][0-9a-fA-F]{40}}/data", handlers.NodeRetrieveDataHandler).
		Methods("GET")
	nodeRouter.HandleFunc("/{nodeAddress:0[xX][0-9a-fA-F]{40}}/data", handlers.NodeSetDataHandler).
		Methods("POST")
	// Node application to Pool
	nodeRouter.HandleFunc("/{nodeAddress:0[xX][0-9a-fA-F]{40}}/apply/{poolAddress:0[xX][0-9a-fA-F]{40}}", handlers.NodeApplyToPoolHandler)
	// Node application status
	nodeRouter.HandleFunc("/{nodeAddress:0[xX][0-9a-fA-F]{40}}/application/{poolAddress:0[xX][0-9a-fA-F]{40}}", handlers.NodeApplicationStatusHandler)

	// Pool Sub-Routes
	poolRouter := apiRouter.PathPrefix("/pool").Subrouter()
	// Retrieve owned Pool if available
	poolRouter.HandleFunc("/", nil)
	// Pool object, data, public key, etc
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}", handlers.PoolRetrievePublicKeyHandler) // TODO temp to display public key
	// Pool data, both public and private data can be set here
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}/data", nil)
	// Retrieve nodes with query parameters for inc data, approved, pending, rejected
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}/nodes", nil)
	// Retrieve or update the status of a node's application
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}/node/{nodeAddress:0[xX][0-9a-fA-F]{40}}/status", nil)

	// Market Sub-Routes
	marketRouter := apiRouter.PathPrefix("/market").Subrouter()
	marketRouter.HandleFunc("/pools", handlers.MarketPoolsHandler).
		Methods("GET").
		Name("market-pools")
	marketRouter.HandleFunc("/pools/create", handlers.MarketPoolsCreateHandler).
		Methods("POST").
		Name("market-pools-create")

	log.Fatal(http.ListenAndServe(":"+PORT, ghandlers.CORS()(router)))
}

func responseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
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
