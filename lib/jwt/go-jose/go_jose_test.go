package go_jose

import (
	"fmt"
	"github.com/go-jose/go-jose/v4"
	"log"
	"os"
	"testing"
)

func TestGoJoseToken_GetJwtToken(t *testing.T) {
	joseToken := NewGoJoseToken("test", 18)
	token, err := joseToken.GetJwtToken()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(token)
	}
}

func TestGoJoseToken_GetJwtRSAToken(t *testing.T) {
	joseToken := NewGoJoseToken("test", 18)
	publicKey, err := LoadPublicKey(cert.PublicKey)
	token, err := joseToken.GetRSAToken(publicKey)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(token)
	}
	privateKey, err := LoadPrivateKey(cert.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
	claims, err := joseToken.DecodeRSAToken(token, privateKey)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(claims)
	}
	fmt.Println(claims)
}

func TestCheckTokenByJwe(t *testing.T) {

	privateKey, err := LoadPrivateKey(cert.PrivateKey)

	privateJSONWEBKey := &jose.JSONWebKey{
		Key:       privateKey,
		KeyID:     "my-ec-key",
		Algorithm: string(cert.Alg),
		Use:       "huizhou92",
	}
	publicKey, err := LoadPublicKey(cert.PublicKey)
	publicJSONWEBKey := &jose.JSONWebKey{
		Key:       publicKey,
		KeyID:     "my-ec-key",
		Algorithm: string(cert.Alg),
		Use:       "huizhou92",
	}

	fmt.Println(privateJSONWEBKey)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{
		Algorithm: cert.Alg,         // 使用 ECDH 密钥交换算法
		Key:       publicJSONWEBKey, // 使用公钥进行加密
	}, nil)
	if err != nil {
		log.Fatalf("failed to create encrypter: %v", err)
	}

	jwe, err := encrypter.Encrypt(jsonClaims)
	if err != nil {
		log.Fatalf("Failed to sign data: %v", err)
	}
	jweToken, err := jwe.CompactSerialize()

	if err != nil {
		log.Fatalf("failed to serialize JWE: %v", err)
	}

	fmt.Printf("Generated JWE Token: %s\n", jweToken)
	//JSONWebEncryption, err := jose.ParseEncryptedJSON(jweToken, []jose.KeyAlgorithm{cert.Alg}, []jose.ContentEncryption{cert.Enc})

	object, err := jose.ParseEncrypted(jweToken, []jose.KeyAlgorithm{cert.Alg}, []jose.ContentEncryption{cert.Enc})
	if err != nil {
		log.Fatalf("failed to parse JWE token: %v", err)
	}

	// 使用从 JWKs 中获取的公钥验证 JWT

	fmt.Printf("Decrypted JWE Payload: %s\n", object.GetAuthData())
}

func TestLoadPrivateKey(t *testing.T) {

	keyBytes, err := os.ReadFile("ec_private.pem")
	if err != nil {
		t.Error(err)
	}
	key, err := LoadPrivateKey(keyBytes)
	if err != nil {
		t.Error(err)
	}
	t.Log(key)
}
