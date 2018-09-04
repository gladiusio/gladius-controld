package controld

import (
	"log"

	"github.com/gladiusio/gladius-application-server/pkg/controller"
	"github.com/gladiusio/gladius-controld/pkg/config"
	"github.com/gladiusio/gladius-controld/pkg/routing"
	gladiusConfig "github.com/gladiusio/gladius-utils/config"
	"github.com/gladiusio/gladius-utils/init/manager"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var Database *gorm.DB

func initializeConfiguration() {
	// Read in gladius-utils directory locations
	gladiusConfig.SetupConfig("gladius-controld", map[string]string{})
	// Set Defaults if no configuration file is found
	config.DefaultConfiguration()
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

// TODO: Refactor so this doesn't have to be here. All calls should be to viper.Get() not the struct
func initializeNodeManagerService(router *mux.Router, configuration config.ConfigurationOptions) {
	// Router Setup
	cRouter := routing.ControlRouter{
		Router: router,
		Port:   viper.GetString("NodeManager.Config.Port"),
		Debug:  configuration.Debug,
	}

	manager.RunService(configuration.Name, configuration.DisplayName, configuration.Description, cRouter.Start)
}

func InitializeNodeManager() {
	// Grab default configuration
	initializeConfiguration()

	configuration := config.ViperConfiguration()

	nodeConfig := configuration.NodeManager.Config
	initializeNodeManagerService(config.NodeRouter(), nodeConfig)
}

func InitializePoolManager() {
	// Grab default configuration
	initializeConfiguration()

	configuration := config.ViperConfiguration()

	poolConfig := configuration.PoolManager.Config

	initializeDatabase(configuration.ApplicationServer.Database)
	initializeService(config.PoolManagerRouter(Database), poolConfig)
}

func InitializeApplicationServer() {
	// Grab default configuration
	initializeConfiguration()

	configuration := config.ViperConfiguration()

	asConfig := configuration.ApplicationServer.Config

	initializeDatabase(configuration.ApplicationServer.Database)
	initializeService(config.ApplicationServerRouter(Database), asConfig)
}
