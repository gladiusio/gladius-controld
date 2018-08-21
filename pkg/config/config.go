package config

import (
	"fmt"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/routing"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

type ConfigurationOptions struct {
	Name        string
	DisplayName string
	Description string
	Debug       bool
	Port        string
}

type ApplicationServerConfig struct {
	Database DatabaseConfig
	Config   ConfigurationOptions
}

func (databaseConfig *DatabaseConfig) GormConnectionString() string {
	sslMode := "disable"

	if databaseConfig.SSL {
		sslMode = "require"
	}

	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.User,
		databaseConfig.Name,
		databaseConfig.Password,
		sslMode,
	)

	return connection
}

type DatabaseConfig struct {
	Type     string
	Host     string
	Port     string
	User     string
	Name     string
	Password string
	SSL      bool
}

type BlockchainConfig struct {
	Provider      string
	MarketAddress string
}

type NodeManagerConfig struct {
	Config ConfigurationOptions
}

type PoolManagerConfig struct {
	Database DatabaseConfig
	Config   ConfigurationOptions
}

type Configuration struct {
	Version           string
	Blockchain        BlockchainConfig
	NodeManager       NodeManagerConfig
	PoolManager       PoolManagerConfig
	ApplicationServer ApplicationServerConfig
}

func DefaultConfiguration() (Configuration, error) {
	var configuration Configuration

	viper.SetConfigName("gladius-controld-defaults")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("\n\nUnable to find gladius-controld.toml in project root, or default directories below.\n\nError: \n%v", err)
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return configuration, nil
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

func PoolManagerRouter(db *gorm.DB) *mux.Router {
	router, ga := setupRouter()

	err := routing.AppendWalletManagementEndpoints(router, ga)
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

	err = routing.AppendPoolManagerEndpoints(router, ga, db)
	if err != nil {
		println("Failed to append pool manager endpoints")
	}

	return router
}

func ApplicationServerRouter(db *gorm.DB) *mux.Router {
	router, _ := setupRouter()

	err := routing.AppendAccountManagementEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Account Management Endpoints")
	}

	err = routing.AppendStatusEndpoints(router)
	if err != nil {
		log.Fatalln("Failed to append Status Endpoints")
	}

	err = routing.AppendServerEndpoints(router, db)
	if err != nil {
		log.Fatalln("Failed to append Server Endpoints")
	}

	err = routing.AppendApplicationEndpoints(router, db)
	if err != nil {
		log.Fatalln("Failed to append Application Endpoints")
	}

	return router
}
