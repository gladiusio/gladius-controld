package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/routing"
	"github.com/gladiusio/gladius-utils/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type ConfigurationOptions struct {
	Name        string
	DisplayName string
	Description string
	Debug       bool
	Port        string
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
	Provider           string
	MarketAddress      string
	PoolManagerAddress string
}

type Configuration struct {
	Version    string
	Build      int
	Blockchain BlockchainConfig
	Directory  struct {
		Base   string
		Wallet string
	}
	NodeManager struct {
		Config ConfigurationOptions
	}
	PoolManager struct {
		Database DatabaseConfig
		Config   ConfigurationOptions
	}
	ApplicationServer struct {
		Database DatabaseConfig
		Config   ConfigurationOptions
	}
}

func (configuration Configuration) defaults() Configuration {
	baseDir, err := config.GetGladiusBase()
	if err != nil {
		baseDir = ""
	}

	return Configuration{
		Version: "0.5.0",
		Build:   20180821,
		Blockchain: BlockchainConfig{
			Provider:           "https://ropsten.infura.io/tjqLYxxGIUp0NylVCiWw",
			MarketAddress:      "0xc4dfb5c9e861eeae844795cfb8d30b77b78bbc38",
			PoolManagerAddress: "0x1f136d7b6308870ed334378f381c9f56d04c3aba",
		},
		Directory: struct {
			Base   string
			Wallet string
		}{
			Base:   baseDir,
			Wallet: filepath.Join(baseDir, "wallet"),
		},
		NodeManager: struct {
			Config ConfigurationOptions
		}{
			Config: struct {
				Name        string
				DisplayName string
				Description string
				Debug       bool
				Port        string
			}{
				Name:        "GladiusNodeControlDaemon",
				DisplayName: "Gladius Node Manager",
				Description: "Gladius Control Daemon",
				Debug:       false,
				Port:        "3001",
			},
		},
	}
}

func DefaultConfiguration() (Configuration, error) {
	var configuration Configuration

	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("\n\nUnable to find gladius-controld.toml in project root, or default directories below.\n\nError: \n%v", err)
		log.Printf("\n\nUsing Default Node Manager Configuration")

		configuration = configuration.defaults()

		viper.SetDefault("blockchain", configuration.Blockchain)
		viper.SetDefault("nodeManager", configuration.NodeManager)

		jsonBytes, err := json.Marshal(configuration)
		if err != nil {
			log.Fatalf("Unable to marshal configuration struct to load defaults, %v", err)
		}

		viper.SetConfigType("json")
		viper.ReadConfig(bytes.NewBuffer(jsonBytes))
	} else {
		err = viper.Unmarshal(&configuration)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
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
