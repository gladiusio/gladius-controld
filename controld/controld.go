package controld

import (
	"github.com/gladiusio/gladius-application-server/pkg/controller"
	"github.com/gladiusio/gladius-controld/pkg/config"
	"github.com/gladiusio/gladius-controld/pkg/routing"
	gladiusConfig "github.com/gladiusio/gladius-utils/config"
	"github.com/gladiusio/gladius-utils/init/manager"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var Database *gorm.DB

func initializeConfiguration() (config.Configuration, error) {
	gladiusConfig.SetupConfig("gladius-controld", map[string]string{})
	return config.DefaultConfiguration()
}

func initializeDatabase(databaseConfig config.DatabaseConfig) {
	db, err := gorm.Open(databaseConfig.Type, databaseConfig.GormConnectionString())
	if err != nil {
		log.Fatal("Could not open database")
	}

	Database, err = controller.Initialize(db)
	if err != nil {
		log.Fatal("Could not migrate database")
	}
}

func initializeService(router *mux.Router, configuration config.ConfigurationOptions) {
	// Router Setup
	cRouter := routing.ControlRouter{
		Router: router,
		Port:   configuration.Port,
		Debug:  configuration.Debug,
	}

	manager.RunService(configuration.Name, configuration.DisplayName, configuration.Description, cRouter.Start)
}

func InitializeNodeManager() {
	// Grab default configuration
	configuration, err := initializeConfiguration()
	if err != nil {
		log.Print(err)
	}

	nodeConfig := configuration.NodeManager.Config
	initializeService(config.NodeRouter(), nodeConfig)
}

func InitializePoolManager() {
	// Grab default configuration
	configuration, err := initializeConfiguration()
	if err != nil {
		log.Print(err)
	}

	poolConfig := configuration.PoolManager.Config

	initializeDatabase(configuration.ApplicationServer.Database)
	initializeService(config.PoolManagerRouter(Database), poolConfig)
}

func InitializeApplicationServer() {
	// Grab default configuration
	configuration, err := initializeConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	asConfig := configuration.ApplicationServer.Config

	initializeDatabase(configuration.ApplicationServer.Database)
	initializeService(config.ApplicationServerRouter(Database), asConfig)
}
