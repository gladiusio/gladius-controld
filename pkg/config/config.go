package config

import (
	"github.com/gladiusio/gladius-controld/pkg/routing"
	"github.com/gladiusio/gladius-utils/config"
	"github.com/gorilla/mux"
	"log"
)

func Config() (Name, DisplayName, Description string) {
	// Setup config handling
	config.SetupConfig("gladius-controld", config.ControlDaemonDefaults())
	return "GladiusControlDaemon", "Gladius Control Daemon", "Gladius Control Daemon"
}

func NodeRouter() (*mux.Router) {
	// Run the function "run" in newtworkd as a service
	router, err := routing.InitializeRouter()
	if err != nil {
		println("Failed to initialized router")
	}

	router, err = routing.AppendP2PEndPoints(router)
	if err != nil {
		log.Fatalln("Failed to append P2P Endpoints")
	}

	router, err = routing.AppendWalletManagementEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Account Management Endpoints")
	}

	router, err = routing.AppendAccountManagementEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Account Management Endpoints")
	}

	router, err = routing.AppendStatusEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Status Endpoints")
	}

	router, err = routing.AppendNodeManagerEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Node Manager Endpoints")
	}

	router, err = routing.AppendMarketEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Market Endpoints")
	}

	return router
}

func PoolManagerRouter() (*mux.Router) {
	// Run the function "run" in newtworkd as a service
	router, err := routing.InitializeRouter()
	if err != nil {
		println("Failed to initialized router")
	}

	router, err = routing.AppendP2PEndPoints(router)
	if err != nil {
		log.Fatalln("Failed to append P2P Endpoints")
	}

	router, err = routing.AppendWalletManagementEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Account Management Endpoints")
	}

	router, err = routing.AppendAccountManagementEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Account Management Endpoints")
	}

	router, err = routing.AppendStatusEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Status Endpoints")
	}

	router, err = routing.AppendMarketEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Market Endpoints")
	}

	router, err = routing.AppendPoolManagerEndpoints(router)
	if err != nil {
		println("Failed to append pool manager endpoints")
	}

	return router
}

func ApplicationServerRouter() (*mux.Router) {
	// Run the function "run" in newtworkd as a service
	router, err := routing.InitializeRouter()
	if err != nil {
		println("Failed to initialized router")
	}

	router, err = routing.AppendAccountManagementEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Account Management Endpoints")
	}

	router, err = routing.AppendStatusEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Status Endpoints")
	}

	router, err = routing.AppendServerEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Server Endpoints")
	}

	router, err = routing.AppendApplicationEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Application Endpoints")
	}

	return router
}