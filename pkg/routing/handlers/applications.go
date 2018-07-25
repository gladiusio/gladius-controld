package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/gladiusio/gladius-application-server/pkg/db/models"
	"github.com/gladiusio/gladius-application-server/pkg/controller"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gorilla/mux"
)

func PoolNewApplicationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestPayload models.NodeRequestPayload
	err := decoder.Decode(&requestPayload)

	requestPayload.IPAddress = r.RemoteAddr

	if err != nil {
		ErrorHandler(w, r, "Could not decode request payload", err, http.StatusBadRequest)
		return
	}

	db, err := controller.Initialize(nil)
	if err != nil {
		ErrorHandler(w, r, "Could not apply to pool", err, http.StatusBadRequest)
		return
	}

	controller.NodeApplyToPool(db, requestPayload)
	viewApplication(w, r, requestPayload.Wallet)
}

func PoolEditApplicationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestPayload models.NodeRequestPayload
	err := decoder.Decode(&requestPayload)
	if err != nil {
		ErrorHandler(w, r, "Could not decode request payload", err, http.StatusBadRequest)
		return
	}

	db, err := controller.Initialize(nil)
	if err != nil {
		ErrorHandler(w, r, "Could not apply to pool", err, http.StatusBadRequest)
		return
	}

	controller.NodeUpdateProfile(db, requestPayload)
	viewApplication(w, r, requestPayload.Wallet)
}

func PoolViewApplicationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wallet := vars["wallet"]
	viewApplication(w, r, wallet)
}

func getProfile(wallet string) (models.NodeProfile, error) {
	db, err := controller.Initialize(nil)
	if err != nil {
		return models.NodeProfile{}, err
	}

	profile, err := controller.NodeProfile(db, wallet)
	if err != nil {
		return models.NodeProfile{}, err
	}

	return profile, err
}

func viewApplication(w http.ResponseWriter, r *http.Request, wallet string) {
	profile, err := getProfile(wallet)
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve profile for wallet: " + wallet, err, http.StatusBadRequest)
		return
	}

	ResponseHandler(w, r, "null", true, nil, profile, nil)
}

type statusResponse struct {
	Accepted bool `json:"accepted"`
	NodeAccepted bool `json:"nodeAcceptance"`
	PoolAccepted bool `json:"poolAcceptance"`
}

func PoolStatusViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wallet := vars["wallet"]

	profile, err := getProfile(wallet)
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve profile for wallet: " + wallet, err, http.StatusBadRequest)
		return
	}

	response := statusResponse{
		Accepted: profile.Accepted.Valid && profile.Accepted.Bool,
		NodeAccepted: profile.NodeAccepted.Valid && profile.NodeAccepted.Bool,
		PoolAccepted: profile.PoolAccepted.Valid && profile.PoolAccepted.Bool,
	}

	ResponseHandler(w, r, "null", true, nil, response, nil)
}