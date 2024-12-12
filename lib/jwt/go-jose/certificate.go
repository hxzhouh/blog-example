package go_jose

import (
	"github.com/go-jose/go-jose/v4"
	"os"
)

type Cert struct {
	PrivateKey []byte
	PublicKey  []byte
	Alg        jose.KeyAlgorithm
	Enc        jose.ContentEncryption
}

var cert *Cert

func init() {
	cert = &Cert{}
	privateKey, err := os.ReadFile("ec_private.pem")
	if err != nil {
		panic(err)
	}
	cert.PrivateKey = privateKey
	publicKey, err := os.ReadFile("ec_public.pem")
	if err != nil {
		panic(err)
	}
	cert.PublicKey = publicKey
	cert.Alg = jose.ECDH_ES_A128KW
	cert.Enc = jose.A256GCM
}
