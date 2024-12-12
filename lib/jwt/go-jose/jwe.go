package go_jose

import (
	"encoding/json"
	"fmt"
	"github.com/go-jose/go-jose/v4"
	"log"
)

func (gjt *GoJoseToken) GetRSAToken(key interface{}) (string, error) {

	crypter, err := jose.NewEncrypter(cert.Enc,
		jose.Recipient{Algorithm: cert.Alg, Key: key}, nil)
	if err != nil {
		log.Fatalf("Failed to create signer: %v", err)
	}
	// Sign the claims
	object, err := crypter.Encrypt(jsonClaims)
	if err != nil {
		log.Fatalf("Failed to sign data: %v", err)
	}

	serialized, err := object.CompactSerialize()
	if err != nil {
		log.Fatalf("Failed to serialize signed object: %v", err)
	}
	fmt.Printf("Signed: %s\n", serialized)
	return serialized, nil
}

func (gjt *GoJoseToken) DecodeRSAToken(serialized string, key interface{}) (map[string]interface{}, error) {
	// Read the private key
	// Parse the encrypted object
	parsedObject, err := jose.ParseEncrypted(serialized, []jose.KeyAlgorithm{cert.Alg},
		[]jose.ContentEncryption{cert.Enc})
	if err != nil {
		log.Fatalf("Failed to parse encrypted object: %v", err)
	}

	// Decrypt the data
	decrypted, err := parsedObject.Decrypt(key)
	if err != nil {
		log.Fatalf("Failed to decrypt data: %v", err)
	}

	// Unmarshal the decrypted data to extract the claims
	var claims map[string]interface{}
	if err = json.Unmarshal(decrypted, &claims); err != nil {
		log.Fatalf("Failed to unmarshal claims: %v", err)
	}

	fmt.Printf("Decrypted claims: %v\n", claims)
	return claims, nil
}
