package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strings"

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
	PoolUrl            string
	PoolManagerAddress string
}

type P2PConfig struct {
	BindPort       int
	AdvertisePort  int
	VerifyOverride bool
}

type Configuration struct {
	Version    string
	Build      int
	Blockchain BlockchainConfig
	P2P        P2PConfig
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

// Returns the singleton viper config as a parsed struct
func ViperConfiguration() Configuration {
	var viperConfiguration Configuration
	viper.Unmarshal(&viperConfiguration)

	return viperConfiguration
}

func (configuration Configuration) defaults() Configuration {
	baseDir, err := config.GetGladiusBase()
	if err != nil {
		baseDir = ""
	}

	return Configuration{
		Version: "0.6.0",
		Build:   20180829,
		Blockchain: BlockchainConfig{
			Provider:           "https://mainnet.infura.io/tjqLYxxGIUp0NylVCiWw",
			MarketAddress:      "0x27a9390283236f836a0b3c8dfdbed2ed854322fc",
			PoolUrl:            "http://174.138.111.1/api/",
			PoolManagerAddress: "0x9717EaDbfE344457135a4f1fA8AE3B11B4CAB0b7",
		},
		P2P: P2PConfig{
			AdvertisePort:  7946,
			BindPort:       7946,
			VerifyOverride: false,
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

func DefaultConfiguration() {
	var configuration Configuration

	// Path of used config value
	configFile := viper.ConfigFileUsed()

	if configFile == "" {
		log.Printf("\n\nUnable to find gladius-controld.toml in project root, or default directories below.\n")
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
	}

	// Setup environment vars, they look like CONTROLD_OBJECT_KEY
	viper.SetEnvPrefix("CONTROLD")
	r := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(r)
	viper.AutomaticEnv()
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
