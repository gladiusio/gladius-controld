package config

import (
	"log"

	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/routing"
	"github.com/gladiusio/gladius-utils/config"
	"github.com/gorilla/mux"
)

func Config() (Name, DisplayName, Description string) {
	// Setup config handling
	config.SetupConfig("gladius-controld", config.ControlDaemonDefaults())
	return "GladiusControlDaemon", "Gladius Control Daemon", "Gladius Control Daemon"
}

func setupRouter() (*mux.Router, *blockchain.GladiusAccountManager) {
	router, err := routing.InitializeRouter()
	if err != nil {
		println("Failed to initialized router")
	}

	// Create a new GladiusAccountManager for the routes, this is so we can have
	// a shared account system between all endpoints
	ga := blockchain.NewGladiusAccountManager()

	return router, ga
}

func NodeRouter() *mux.Router {
	router, ga := setupRouter()

	err := routing.AppendP2PEndPoints(router, ga)
	if err != nil {
		log.Fatalln("Failed to append P2P Endpoints")
	}

	err = routing.AppendWalletManagementEndpoints(router, ga)
	if err != nil {
		log.Fatalln("Failed to append Account Management Endpoints")
	}

	err = routing.AppendAccountManagementEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Account Management Endpoints")
	}

	err = routing.AppendStatusEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Status Endpoints")
	}

	err = routing.AppendNodeManagerEndpoints(router, ga)
	if err != nil {
		log.Fatalln("Failed to append Node Manager Endpoints")
	}

	err = routing.AppendMarketEndpoints(router, ga)
	if err != nil {
		log.Fatalln("Failed to append Market Endpoints")
	}

	return router
}

func PoolManagerRouter() *mux.Router {
	router, ga := setupRouter()

	err := routing.AppendP2PEndPoints(router, ga)
	if err != nil {
		log.Fatalln("Failed to append P2P Endpoints")
	}

	err = routing.AppendWalletManagementEndpoints(router, ga)
	if err != nil {
		log.Fatalln("Failed to append Account Management Endpoints")
	}

	err = routing.AppendAccountManagementEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Account Management Endpoints")
	}

	err = routing.AppendStatusEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Status Endpoints")
	}

	err = routing.AppendMarketEndpoints(router, ga)
	if err != nil {
		log.Fatalln("Failed to append Market Endpoints")
	}

	err = routing.AppendPoolManagerEndpoints(router, ga)
	if err != nil {
		println("Failed to append pool manager endpoints")
	}

	return router
}

func ApplicationServerRouter() *mux.Router {
	router, _ := setupRouter()

	err := routing.AppendAccountManagementEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Account Management Endpoints")
	}

	err = routing.AppendStatusEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Status Endpoints")
	}

	err = routing.AppendServerEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Server Endpoints")
	}

	err = routing.AppendApplicationEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Application Endpoints")
	}

	return router
}
