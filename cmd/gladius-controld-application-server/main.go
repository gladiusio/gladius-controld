package main

import (
	"github.com/gladiusio/gladius-application-server/pkg/controller"
	. "github.com/gladiusio/gladius-controld/pkg/config"
	"github.com/gladiusio/gladius-controld/pkg/routing"
	"github.com/gladiusio/gladius-utils/init/manager"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var Database *gorm.DB

func main() {
	// Define some variables
	name, displayName, description := Config()

	configuration, err := DefaultConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(configuration.ApplicationServer.Database.Type, configuration.ApplicationServer.Database.GormConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	Database, err = controller.Initialize(db)
	if err != nil {
		log.Fatal("Could not open database")
	}

	// Run the function "run" in newtworkd as a service
	router := ApplicationServerRouter(Database)

	cRouter := routing.ControlRouter{
		Router: router,
		Port:   configuration.ApplicationServer.Config.Port,
	}

	// Run the function "run" in newtworkd as a service
	manager.RunService(name, displayName, description, cRouter.Start)
}