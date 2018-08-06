package handlers

import (
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
