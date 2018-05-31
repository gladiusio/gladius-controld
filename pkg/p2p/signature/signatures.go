package signature

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/p2p/message"
)

// SignedMessage is a type representing a signed message
type SignedMessage struct {
	Message   []byte `json:"message"`
	Hash      []byte `json:"hash"`
	Signature []byte `json:"signature"`
	Address   string `json:"address"`
}

// ParseSignedMessage returns a signed message to be passed into the VerifySignedMessgae method
func ParseSignedMessage(message, hash, signature, address string) (*SignedMessage, error) {
	dMessage, err := b64.StdEncoding.DecodeString(message)
	if err != nil {
		return nil, errors.New("error decoding message")
	}
	dHash, err := b64.StdEncoding.DecodeString(hash)
	if err != nil {
		return nil, errors.New("error decoding hash")
	}
	dSignature, err := b64.StdEncoding.DecodeString(signature)
	if err != nil {
		return nil, errors.New("error decoding signature")
	}
	fmt.Println("Decoded signature: ", dSignature)
	return &SignedMessage{Message: dMessage, Hash: dHash, Signature: dSignature, Address: address}, nil
}

// CreateSignedMessage creates a signed state from the message where
func CreateSignedMessage(message *message.Message, passphrase string) (string, error) {
	ga := blockchain.NewGladiusAccountManager()
	err := ga.UnlockAccount(passphrase)
	if err != nil {
		return "", errors.New("Error unlocking wallet")
	}

	// Create a serailized JSON string
	messageBytes := message.Serialize()

	hash := crypto.Keccak256(messageBytes)
	signature, err := ga.Keystore().SignHash(ga.GetAccount(), hash)
	fmt.Println("Signed signature: ", signature)

	if err != nil {
		return "", errors.New("Error signing message")
	}

	// Create the signed message
	signed := &SignedMessage{Message: messageBytes, Hash: hash, Signature: signature, Address: ga.GetAccountAddress().String()}

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
	hashMatches := bytes.Equal(sm.Hash, crypto.Keccak256(sm.Message))

	pub, err := crypto.SigToPub(sm.Hash, sm.Signature)
	if err != nil {
		return false, errors.New("Error signing message")
	}

	fmt.Println(len(sm.Signature))
	// Check if the signature is valid
	signatureValid := crypto.VerifySignature(crypto.CompressPubkey(pub), sm.Hash, sm.Signature[:64])

	// Check if the address matches
	addressMatches := crypto.PubkeyToAddress(*pub).String() == sm.Address
	fmt.Println(addressInPool, hashMatches, signatureValid, addressMatches)
	if addressInPool && addressMatches && hashMatches && signatureValid {
		return true, nil
	}

	return false, nil
}
