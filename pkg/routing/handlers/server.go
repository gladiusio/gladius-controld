package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/gladiusio/gladius-application-server/pkg/controller"
)

// Retrieve Pool Information
func PublicPoolInformationHandler(database *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		poolInformation, err := controller.PoolInformation(database)
		if err != nil {
			ErrorHandler(w, r, "Could retrieve Public Information", err, http.StatusBadRequest)
			return
		}

		ResponseHandler(w, r, "null", true, nil, poolInformation, nil)
	}
}

type PoolContainsWallet struct {
	ContainsWallet bool
}

func PoolContainsNode(database *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		walletAddress := vars["walletAddress"]
		containsWallet, err := controller.NodeInPool(database, walletAddress)
		if err != nil {
			ErrorHandler(w, r, "Could not query server", err, http.StatusInternalServerError)
			return
		}

		ResponseHandler(w, r, "null", true, nil, PoolContainsWallet{ContainsWallet: containsWallet}, nil)
	}
}

func PoolNodes(database *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		nodes, err := controller.NodesAccepted(database)
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve nodes", err, http.StatusInternalServerError)
			return
		}

		var nodeAddresses []string

		for _, node := range nodes {
			nodeAddresses = append(nodeAddresses, node.Wallet)
		}

		ResponseHandler(w, r, "null", true, nil, nodeAddresses, nil)

	}
}
