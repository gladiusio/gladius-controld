package main

import (
	. "github.com/gladiusio/gladius-controld/pkg/config"
	"github.com/gladiusio/gladius-controld/pkg/routing"
	"github.com/gladiusio/gladius-utils/init/manager"
)

func main() {
	// Define some variables
	name, displayName, description := Config()

	// Run the function "run" in newtworkd as a service
	manager.RunService(name, displayName, description, initialize)
}

func initialize() {
	router := NodeRouter()
	routing.Start(router, nil)
}
