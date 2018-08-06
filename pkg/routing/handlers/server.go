package handlers

import (
	"net/http"

	"github.com/gladiusio/gladius-application-server/pkg/controller"
)

// Retrieve Pool Information
func PublicPoolInformationHandler(w http.ResponseWriter, r *http.Request) {
	db, err := controller.Initialize(nil)
	defer db.Close()

	if err != nil {
		ErrorHandler(w, r, "Could retrieve Public Information", err, http.StatusBadRequest)
		return
	}

	poolInformation, err := controller.PoolInformation(db)
	if err != nil {
		ErrorHandler(w, r, "Could retrieve Public Information", err, http.StatusBadRequest)
		return
	}

	ResponseHandler(w, r, "null", true, nil, poolInformation, nil)
}
