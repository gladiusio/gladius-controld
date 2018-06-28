package signature

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"regexp"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/p2p/message"
	"github.com/tdewolff/minify"
	mjson "github.com/tdewolff/minify/json"
)

// SignedMessage is a type representing a signed message
type SignedMessage struct {
	Message   *json.RawMessage `json:"message"`
	Hash      []byte           `json:"hash"`
	Signature []byte           `json:"signature"`
	Address   string           `json:"address"`
}

// ParseSignedMessage returns a signed message to be passed into the VerifySignedMessage method
func ParseSignedMessage(message, hash, signature, address string) (*SignedMessage, error) {
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)
	messageMin, err := m.String("text/json", message)
	if err != nil {
		panic(err)
	}

	h := json.RawMessage(messageMin)
	dHash, err := b64.StdEncoding.DecodeString(hash)
	if err != nil {
		return nil, errors.New("error decoding hash")
	}
	dSignature, err := b64.StdEncoding.DecodeString(signature)
	if err != nil {
		return nil, errors.New("error decoding signature")
	}
	return &SignedMessage{Message: &h, Hash: dHash, Signature: dSignature, Address: address}, nil
}

// CreateSignedMessage creates a signed state from the message where
func CreateSignedMessage(message *message.Message, passphrase string) (string, error) {
	ga := blockchain.NewGladiusAccountManager()
	success, err := ga.UnlockAccount(passphrase)
	if success == false || err != nil {
		return "", errors.New("Error unlocking wallet")
	}

	// Create a serailized JSON string
	messageBytes := message.Serialize()

	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)
	messageBytes, err = m.Bytes("text/json", messageBytes)
	if err != nil {
		panic(err)
	}

	hash := crypto.Keccak256(messageBytes)
	account, err := ga.GetAccount()

	if err != nil {
		return "", err
	}

	signature, err := ga.Keystore().SignHash(*account, hash)
	if err != nil {
		return "", errors.New("Error signing message")
	}

	address, err := ga.GetAccountAddress()
	if err != nil {
		return "", err
	}

	addressString := address.String()

	h := json.RawMessage(messageBytes)

	// Create the signed message
	signed := &SignedMessage{
		Message:   &h,
		Hash:      hash,
		Signature: signature,
		Address:   addressString,
	}

	// Encode the struct as a json
	bytes, err := json.Marshal(signed)
	if err != nil {
		panic(err)
	}

	return string(bytes), err
}

// VerifySignedMessage Verifies that a signed message is valid
func VerifySignedMessage(sm *SignedMessage) (bool, error) {
	// Check if address is part of pool
	// TODO: Check real address against pool
	addressInPool := true
	// Check if hash matches the message
	m, _ := sm.Message.MarshalJSON()
	hashMatches := bytes.Equal(sm.Hash, crypto.Keccak256(m))

	pub, err := crypto.SigToPub(sm.Hash, sm.Signature)
	if err != nil {
		return false, errors.New("Error signing message")
	}

	// Check if the signature is valid
	signatureValid := crypto.VerifySignature(crypto.CompressPubkey(pub), sm.Hash, sm.Signature[:64])

	// Check if the address matches
	addressMatches := crypto.PubkeyToAddress(*pub).String() == sm.Address

	if addressInPool && addressMatches && hashMatches && signatureValid {
		return true, nil
	}

	return false, nil
}
