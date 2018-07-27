package crypto

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"os"
		"github.com/spf13/viper"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	)

const (
	md5       = 1
	sha1      = 2
	ripemd160 = 3
	sha256    = 8
	sha384    = 9
	sha512    = 10
	sha224    = 11
)

type Key struct {
	openpgp.Entity
}

func DecryptData(message string) (string, error) {
	var pathTemp = viper.GetString("DirKeys")
	keyringFileBuffer, err := os.Open(pathTemp + "/private.asc")
	if err != nil {
		return "", err
	}

	defer keyringFileBuffer.Close()

	entityList, err := openpgp.ReadArmoredKeyRing(keyringFileBuffer)

	// Decode the base64 string
	dec, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		// TODO this resolves the prior keybase message submission issue, but should return an err instead of nil
		return "", nil
	}

	// Decrypt it with the contents of the private key
	md, err := openpgp.ReadMessage(bytes.NewBuffer(dec), entityList, nil, nil)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return "", err
	}
	decStr := string(bytes)

	return decStr, nil
}

func (key *Key) armor() (string, error) {
	buf := new(bytes.Buffer)
	armor, err := armor.Encode(buf, openpgp.PublicKeyType, nil)
	if err != nil {
		return "", err
	}
	key.Serialize(armor)
	armor.Close()

	return buf.String(), nil
}

func (key *Key) armorPrivate() (string, error) {
	buf := new(bytes.Buffer)
	armor, err := armor.Encode(buf, openpgp.PrivateKeyType, nil)
	if err != nil {
		return "", err
	}

	key.SerializePrivate(armor, nil)
	armor.Close()

	return buf.String(), nil
}
