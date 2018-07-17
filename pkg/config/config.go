package config

import (
	"github.com/gladiusio/gladius-controld/pkg/routing"
	"github.com/gladiusio/gladius-utils/config"
	"github.com/gorilla/mux"
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

	router, err = routing.AppendNodeEndpoints(router)
	if err != nil {
		println("Failed to append node endpoints")
	}

	return router
}

func PoolManagerRouter() (*mux.Router) {
	// Run the function "run" in newtworkd as a service
	router := NodeRouter()
	router, err := routing.AppendPoolManagerEndpoints(router)
	if err != nil {
		println("Failed to append pool manager endpoints")
	}

	println("Initializing Pool Manager")

	return router
}