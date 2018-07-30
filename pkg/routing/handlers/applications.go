package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/gladiusio/gladius-application-server/pkg/db/models"
	"github.com/gladiusio/gladius-application-server/pkg/controller"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gorilla/mux"
	"net"
	"bytes"
	"strings"
)

type ipRange struct {
	start net.IP
	end   net.IP
}

var privateRanges = []ipRange{
	ipRange{
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	ipRange{
		start: net.ParseIP("100.64.0.0"),
		end:   net.ParseIP("100.127.255.255"),
	},
	ipRange{
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	ipRange{
		start: net.ParseIP("192.0.0.0"),
		end:   net.ParseIP("192.0.0.255"),
	},
	ipRange{
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	ipRange{
		start: net.ParseIP("198.18.0.0"),
		end:   net.ParseIP("198.19.255.255"),
	},
}

// isPrivateSubnet - check to see if this ip is in a private subnet
func isPrivateSubnet(ipAddress net.IP) bool {
	// my use case is only concerned with ipv4 atm
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		// iterate over all our ranges
		for _, r := range privateRanges {
			// check if this ip is in a private range
			if inRange(r, ipAddress) {
				return true
			}
		}
	}
	return false
}

// inRange - check to see if a given ip address is within a range given
func inRange(r ipRange, ipAddress net.IP) bool {
	// strcmp type byte comparison
	if bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0 {
		return true
	}
	return false
}

func getIPAdress(r *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		addresses := strings.Split(r.Header.Get(h), ",")
		// march from right to left until we get a public address
		// that will be the address right before our proxy.
		for i := len(addresses) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addresses[i])
			// header can contain spaces too, strip those out.
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() || isPrivateSubnet(realIP) {
				// bad address, go to next
				continue
			}
			return ip
		}
	}
	return ""
}

func PoolNewApplicationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestPayload models.NodeRequestPayload
	err := decoder.Decode(&requestPayload)

	requestPayload.IPAddress = getIPAdress(r)

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

func getProfile(wallet string) (controller.FullProfile, error) {
	db, err := controller.Initialize(nil)
	if err != nil {
		return controller.FullProfile{}, err
	}

	profile, err := controller.NodePoolApplication(db, wallet)
	if err != nil {
		return controller.FullProfile{}, err
	}

	return profile, err
}

func viewApplication(w http.ResponseWriter, r *http.Request, wallet string) {
	profile, err := getProfile(wallet)
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve profile for wallet: "+wallet, err, http.StatusBadRequest)
		return
	}

	ResponseHandler(w, r, "null", true, nil, profile, nil)
}

type statusResponse struct {
	Accepted     bool `json:"accepted"`
	NodeAccepted bool `json:"nodeAcceptance"`
	PoolAccepted bool `json:"poolAcceptance"`
}

func PoolStatusViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wallet := vars["wallet"]

	profile, err := getProfile(wallet)
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve profile for wallet: "+wallet, err, http.StatusBadRequest)
		return
	}

	response := statusResponse{
		Accepted:     profile.NodeProfile.Accepted.Valid && profile.NodeProfile.Accepted.Bool,
		NodeAccepted: profile.NodeProfile.NodeAccepted.Valid && profile.NodeProfile.NodeAccepted.Bool,
		PoolAccepted: profile.NodeProfile.PoolAccepted.Valid && profile.NodeProfile.PoolAccepted.Bool,
	}

	ResponseHandler(w, r, "null", true, nil, response, nil)
}
