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

var apiRouter *mux.Router

func Start(router *mux.Router, port *string)  {
	if port != nil {
		fmt.Println("Starting API at http://localhost:" + *port)
		log.Fatal(http.ListenAndServe(":" + *port, ghandlers.CORS()(router)))
	} else {
		fmt.Println("Starting API at http://localhost:" + PORT)
		log.Fatal(http.ListenAndServe(":"+PORT, ghandlers.CORS()(router)))
	}
}

func InitializeRouter() (*mux.Router, error) {
	// Main Router
	router := mux.NewRouter()
	if DEBUG {
		router.Use(loggingMiddleware)
	}

	return router, nil
}

func InitializeAPISubRoutes(router *mux.Router) {
	// Base API Sub-Routes
	if apiRouter == nil {
		apiRouter = router.PathPrefix("/api").Subrouter()
		apiRouter.Use(responseMiddleware)
		apiRouter.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	}
}

func AppendNodeEndpoints(router *mux.Router) (*mux.Router, error) {
	// Initialize Base API sub-route
	InitializeAPISubRoutes(router)
	// P2P setup
	peerNetwork := peer.New()
	p2pRouter := apiRouter.PathPrefix("/p2p").Subrouter()
	// P2P Message Routes
	p2pRouter.HandleFunc("/message/sign", handlers.CreateSignedMessageHandler).
		Methods(http.MethodPost)
	p2pRouter.HandleFunc("/message/verify", handlers.VerifySignedMessageHandler).
		Methods(http.MethodPost)
	// P2P State Routes
	p2pRouter.HandleFunc("/state/push_message", handlers.PushStateMessageHandler(peerNetwork)).
		Methods(http.MethodPost)
	p2pRouter.HandleFunc("/state/", handlers.GetFullStateHandler(peerNetwork)).
		Methods(http.MethodGet)
	p2pRouter.HandleFunc("/state/", handlers.PushStateMessageHandler(peerNetwork)).
		Methods(http.MethodPost)

	// Key Management
	walletRouter := apiRouter.PathPrefix("/keystore").Subrouter()
	walletRouter.HandleFunc("/account/create", handlers.KeystoreAccountCreationHandler).
		Methods(http.MethodPost)
	walletRouter.HandleFunc("/account", handlers.KeystoreAccountRetrievalHandler)
	walletRouter.HandleFunc("/account/open", handlers.KeystoreAccountUnlockHandler).
		Methods(http.MethodPost)
	walletRouter.HandleFunc("/pgp/view/public", handlers.KeystorePGPPublicKeyRetrievalHandler)
	walletRouter.HandleFunc("/pgp/create", handlers.KeystorePGPCreationHandler).
		Methods(http.MethodPost)

	// Account Management
	accountRouter := apiRouter.PathPrefix("/account/{address:0[xX][0-9a-fA-F]{40}}").Subrouter()
	accountRouter.HandleFunc("/balance/{symbol:[a-z]{3}}", handlers.AccountBalanceHandler)
	accountRouter.HandleFunc("/transactions", handlers.AccountTransactionsHandler).
		Methods(http.MethodPost)

	// Status Sub-Routes
	statusRouter := apiRouter.PathPrefix("/status").Subrouter()
	statusRouter.HandleFunc("/", handlers.StatusHandler).
		Methods(http.MethodGet, http.MethodPut).
		Name("status")
	statusRouter.HandleFunc("/tx/{tx:0[xX][0-9a-fA-F]{64}}", handlers.StatusTxHandler).
		Methods(http.MethodGet).
		Name("status-tx")

	// Node Sub-Routes
	nodeRouter := apiRouter.PathPrefix("/node").Subrouter()
	// Node pool applications
	nodeRouter.HandleFunc("/applications", handlers.NodeViewAllApplicationsHandler).
		Methods(http.MethodGet)
	// Node application to Pool
	nodeRouter.HandleFunc("/applications/{poolAddress:0[xX][0-9a-fA-F]{40}}/new", handlers.NodeNewApplicationHandler).
		Methods(http.MethodPost)
	nodeRouter.HandleFunc("/applications/{poolAddress:0[xX][0-9a-fA-F]{40}}/view", handlers.NodeViewApplicationHandler).
		Methods(http.MethodGet)

	// Pool Sub-Routes
	poolRouter := apiRouter.PathPrefix("/pool").Subrouter()
	// Retrieve owned Pool if available
	poolRouter.HandleFunc("/", nil)
	// Pool object, data, public key, etc
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}", handlers.PoolRetrievePublicKeyHandler) // TODO temp to display public key
	// Pool Retrieve Data
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}/data", handlers.PoolPublicDataHandler).
		Methods(http.MethodGet)

	// Market Sub-Routes
	marketRouter := apiRouter.PathPrefix("/market").Subrouter()
	marketRouter.HandleFunc("/pools", handlers.MarketPoolsHandler)

	return router, nil
}

func AppendPoolManagerEndpoints(router *mux.Router) (*mux.Router, error) {
	// Initialize Base API sub-route
	InitializeAPISubRoutes(router)

	// Pool
	poolRouter := apiRouter.PathPrefix("/pool").Subrouter()
	// Pool data, both public and private data can be set here
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}/data", handlers.PoolPublicDataHandler).
		Methods(http.MethodPost)
	// Retrieve nodes with query parameters for inc data, approved, pending, rejected
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}/nodes/{status:.*}", handlers.PoolRetrieveNodesHandler)
	// Retrieve node application
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}/node/{nodeAddress:0[xX][0-9a-fA-F]{40}}/application", handlers.PoolRetrieveNodeApplicationHandler)
	// Retrieve or update the status of a node's application
	poolRouter.HandleFunc("/{poolAddress:0[xX][0-9a-fA-F]{40}}/node/{nodeAddress:0[xX][0-9a-fA-F]{40}}/{status}", handlers.PoolUpdateNodeStatusHandler).
		Methods(http.MethodPut)

	// Market
	marketRouter := apiRouter.PathPrefix("/market").Subrouter()
	marketRouter.HandleFunc("/pools/owned", handlers.MarketPoolsOwnedHandler)
	marketRouter.HandleFunc("/pools/create", handlers.MarketPoolsCreateHandler).
		Methods(http.MethodPost)

	// Applications
	applicationRouter := apiRouter.PathPrefix("/application").Subrouter()
	applicationRouter.HandleFunc("/new", handlers.PoolNewApplicationHandler).
		Methods(http.MethodPost)
	applicationRouter.HandleFunc("/edit", handlers.PoolEditApplicationHandler).
		Methods(http.MethodPost)
	applicationRouter.HandleFunc("/view/{wallet:0[xX][0-9a-fA-F]{40}}", handlers.PoolViewApplicationHandler).
		Methods(http.MethodGet)
	applicationRouter.HandleFunc("/status/{wallet:0[xX][0-9a-fA-F]{40}}", handlers.PoolStatusViewHandler)

	return router, nil
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
	if r.Method == http.MethodPost {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}

	// Return the request as a string
	return strings.Join(request, "\n")
}
