package main

import (
	"github.com/gladiusio/gladius-application-server/pkg/controller"
	. "github.com/gladiusio/gladius-controld/pkg/config"
	"github.com/gladiusio/gladius-controld/pkg/routing"
	"github.com/gladiusio/gladius-utils/init/manager"
	"github.com/jinzhu/gorm"
	"log"
)

var Database *gorm.DB

func main() {
	// Define some variables
	name, displayName, description := Config()

	var err error
	Database, err = controller.Initialize(nil)
	if err != nil {
		log.Fatal("Could not open database")
	}

	// Run the function "run" in newtworkd as a service
	manager.RunService(name, displayName, description, initialize)
}

func initialize() {
	// Run the function "run" in newtworkd as a service
	router := ApplicationServerRouter(Database)
	port := "3333"
	routing.Start(router, &port)
}
