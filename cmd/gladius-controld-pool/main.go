package main

import (
	"github.com/gladiusio/gladius-controld/pkg/routing"
	"github.com/gladiusio/gladius-utils/init/manager"
	. "github.com/gladiusio/gladius-controld/pkg/config"
)

func main() {
	// Define some variables
	name, displayName, description := Config()

	// Run the function "run" in newtworkd as a service
	manager.RunService(name, displayName, description, initialize)
}

func initialize() {
	// Run the function "run" in newtworkd as a service
	router := PoolManagerRouter()
	routing.Start(router)
}