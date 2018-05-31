package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gladiusio/gladius-controld/pkg/p2p/message"
	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
)

type stateBody struct {
	State      string `json:"state"`
	Passphrase string `json:"passphrase"`
}

type signatureBody struct {
	Message   string `json:"message"`
	Hash      string `json:"hash"`
	Signature string `json:"signature"`
	Address   string `json:"address"`
}

func stateBodyDecoder(w http.ResponseWriter, r *http.Request) (*stateBody, error) {
	decoder := json.NewDecoder(r.Body)
	var body stateBody
	err := decoder.Decode(&body)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return &body, nil
}

func signatureBodyDecoder(w http.ResponseWriter, r *http.Request) (*signatureBody, error) {
	decoder := json.NewDecoder(r.Body)
	var body signatureBody
	err := decoder.Decode(&body)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return &body, nil
}

func PeerToPeerStateUpdateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := stateBodyDecoder(w, r)
	if err != nil {
		ErrorHandler(w, r, "Could not find `passphrase` or could not find `state` in request", err, http.StatusInternalServerError)
		return
	}
	message := message.New(body.State)

	signed, err := signature.CreateSignedMessage(message, body.Passphrase)
	if err != nil {
		ErrorHandler(w, r, "Could not find sign message. Passphrase likely incorrect.", err, http.StatusInternalServerError)
		return
	}
	ResponseHandler(w, r, "null", signed)
}

func VerifySignedMessageHandler(w http.ResponseWriter, r *http.Request) {
	body, err := signatureBodyDecoder(w, r)
	if err != nil {
		ErrorHandler(w, r, "Missing one or more of: `message`, `hash`, `signature`, `address`", err, http.StatusInternalServerError)
		return
	}

	parsed, err := signature.ParseSignedMessage(body.Message, body.Hash, body.Signature, body.Address)
	if err != nil {
		ErrorHandler(w, r, "Couldn't parse body", err, http.StatusInternalServerError)
		return
	}
	verified, err := signature.VerifySignedMessage(parsed)
	if err != nil {
		ErrorHandler(w, r, "Error veryfing signature", err, http.StatusInternalServerError)
		return
	}

	ResponseHandler(w, r, "null", strconv.FormatBool(verified))
}
