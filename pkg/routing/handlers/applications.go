package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gladiusio/gladius-controld/pkg/p2p/message"
	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
	"github.com/pkg/errors"
	"net"
	"net/http"
	"strings"

	"github.com/gladiusio/gladius-common/pkg/db/models"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type statusResponse struct {
	Accepted     bool `json:"accepted"`
	NodeAccepted bool `json:"nodeAcceptance"`
	PoolAccepted bool `json:"poolAcceptance"`
}

func getSignedMessage(r *http.Request) (signature.SignedMessage, error) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var signedMessage signature.SignedMessage
	err := decoder.Decode(&signedMessage)
	if err != nil {
		return signature.SignedMessage{}, err
	}

	return signedMessage, nil
}

func getMessage(signedMessage signature.SignedMessage) (message.Message, error) {
	if signedMessage.Message == nil || !signedMessage.IsVerified() {
		return message.Message{}, errors.New("Message failed verification")
	}

	var unsignedMessage message.Message
	rawMessage, err := signedMessage.Message.MarshalJSON()
	if err != nil {
		return message.Message{}, err
	}
	json.Unmarshal(rawMessage, &unsignedMessage)

	return unsignedMessage, nil
}

func getRequestPayload(r *http.Request) (models.NodeRequestPayload, error) {
	signedMessage, err := getSignedMessage(r)
	if err != nil {
		return models.NodeRequestPayload{}, err
	}

	unsignedMessage, err := getMessage(signedMessage)
	if err != nil {
		return models.NodeRequestPayload{}, err
	}

	var requestPayload models.NodeRequestPayload
	messageContent, err := unsignedMessage.Content.MarshalJSON()
	if err != nil {
		return models.NodeRequestPayload{}, err
	}

	json.Unmarshal(messageContent, &requestPayload)

	requestPayload.IPAddress = getIPAddress(r)

	return requestPayload, nil
}

type ipRange struct {
	start net.IP
	end   net.IP
}

var privateRanges = []ipRange{
	{
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	{
		start: net.ParseIP("100.64.0.0"),
		end:   net.ParseIP("100.127.255.255"),
	},
	{
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	{
		start: net.ParseIP("192.0.0.0"),
		end:   net.ParseIP("192.0.0.255"),
	},
	{
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	{
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

func getIPAddress(r *http.Request) string {
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
