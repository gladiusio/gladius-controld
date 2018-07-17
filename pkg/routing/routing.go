package routing

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gladiusio/gladius-controld/pkg/p2p/peer"
	"github.com/gladiusio/gladius-controld/pkg/routing/handlers"
	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	PORT  = "3001"
	DEBUG = false
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

	// P2P setup
	peer := peer.New()
	p2pRouter := apiRouter.PathPrefix("/p2p").Subrouter()

	// P2P Message Routes
	p2pRouter.HandleFunc("/message/sign", handlers.CreateSignedMessageHandler).
		Methods("POST")
	p2pRouter.HandleFunc("/message/verify", handlers.VerifySignedMessageHandler).
		Methods("POST")

	p2pRouter.HandleFunc("/discovery/introduce", handlers.IntroductionHandler(peer)).
		Methods("POST")

	// P2P State Routes
	p2pRouter.HandleFunc("/state/push_message", handlers.PushStateMessageHandler(peer)).
		Methods("POST")
	p2pRouter.HandleFunc("/state/", handlers.GetFullStateHandler(peer)).
		Methods("GET")
	p2pRouter.HandleFunc("/state/signatures", handlers.GetSignatureListHandler(peer)).
		Methods("GET")
	p2pRouter.HandleFunc("/state/content_diff", handlers.GetContentHandler(peer)).
		Methods("POST")

	// Key Management
	walletRouter := apiRouter.PathPrefix("/keystore").Subrouter()
	walletRouter.HandleFunc("/account/create", handlers.KeystoreAccountCreationHandler).
		Methods("POST")
	walletRouter.HandleFunc("/account", handlers.KeystoreAccountRetrievalHandler)
	walletRouter.HandleFunc("/account/open", handlers.KeystoreAccountUnlockHandler).
		Methods("POST")
	walletRouter.HandleFunc("/pgp/view/public", handlers.KeystorePGPPublicKeyRetrievalHandler)
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
	nodeRouter.HandleFunc("/", handlers.NodeFactoryNodeAddressHandler)
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
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}/data", handlers.PoolPublicDataHandler).
		Methods("GET", "POST")
	// Retrieve nodes with query parameters for inc data, approved, pending, rejected
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}/nodes/{status:.*}", handlers.PoolRetrieveNodesHandler)
	// Retrieve node application
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}/node/{nodeAddress:0[xX][0-9a-fA-F]{40}}/application", handlers.PoolRetrieveNodeApplicationHandler)
	// Retrieve or update the status of a node's application
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}/node/{nodeAddress:0[xX][0-9a-fA-F]{40}}/{status}", handlers.PoolUpdateNodeStatusHandler).
		Methods("PUT")

	// Market Sub-Routes
	marketRouter := apiRouter.PathPrefix("/market").Subrouter()
	marketRouter.HandleFunc("/pools", handlers.MarketPoolsHandler)
	marketRouter.HandleFunc("/pools/owned", handlers.MarketPoolsOwnedHandler)
	marketRouter.HandleFunc("/pools/create", handlers.MarketPoolsCreateHandler).
		Methods("POST")

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
