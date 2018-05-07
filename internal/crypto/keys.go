package crypto

import (
	"bytes"
	"encoding/base64"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
	"io/ioutil"
	"os"
	"strings"
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

func CreateKeyPair(name, comment, email string) (string, error) {
	var pathTemp string = os.Getenv("HOME") + "/.config/gladius/keys/"

	entity, _ := openpgp.NewEntity(name, comment, email, nil)
	key := Key{*entity}

	for _, id := range key.Identities {
		id.SelfSignature.PreferredSymmetric = []uint8{
			uint8(packet.CipherAES256),
			uint8(packet.CipherAES192),
			uint8(packet.CipherAES128),
			uint8(packet.CipherCAST5),
			uint8(packet.Cipher3DES),
		}

		id.SelfSignature.PreferredHash = []uint8{
			sha256,
			sha1,
			sha384,
			sha512,
			sha224,
		}

		id.SelfSignature.PreferredCompression = []uint8{
			uint8(packet.CompressionZLIB),
			uint8(packet.CompressionZIP),
		}

		id.SelfSignature.SignUserId(id.UserId.Id, key.PrimaryKey, key.PrivateKey, nil)
	}

	// Self-sign the Subkeys
	for _, subkey := range key.Subkeys {
		subkey.Sig.SignKey(subkey.PublicKey, key.PrivateKey, nil)
	}

	publicKey, err := key.armor()

	if err != nil {
		return "", err
	}

	privateKey, err := key.armorPrivate()

	if err != nil {
		return "", err
	}

	os.MkdirAll(pathTemp, os.ModePerm)

	ioutil.WriteFile(pathTemp+"private.asc", []byte(privateKey), 0600)
	ioutil.WriteFile(pathTemp+"public.asc", []byte(publicKey), 0644)

	return pathTemp + "private.asc", nil
}

func EncryptData(data string) (string, error) {
	var pathTemp string = os.Getenv("HOME") + "/.config/gladius/keys/"
	keyringFileBuffer, _ := ioutil.ReadFile(pathTemp + "public.asc")

	publicKey := string(keyringFileBuffer)

	return EncryptMessage(data, publicKey)
}

func EncryptMessage(message, publicKey string) (string, error) {
	publicKeyReader := strings.NewReader(publicKey)

	entityList, err := openpgp.ReadArmoredKeyRing(publicKeyReader)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	w, err := openpgp.Encrypt(buf, entityList, nil, nil, nil)
	if err != nil {
		return "", err
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		return "", err
	}
	err = w.Close()
	if err != nil {
		return "", err
	}

	// Encode to base64
	bytes, err := ioutil.ReadAll(buf)
	if err != nil {
		return "", err
	}
	encStr := base64.StdEncoding.EncodeToString(bytes)

	return encStr, nil
}

func DecryptData(message string) (string, error) {
	var pathTemp string = os.Getenv("HOME") + "/.config/gladius/keys/"
	keyringFileBuffer, _ := os.Open(pathTemp + "private.asc")

	defer keyringFileBuffer.Close()

	entityList, err := openpgp.ReadArmoredKeyRing(keyringFileBuffer)

	// Decode the base64 string
	dec, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", err
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
