package go_jose

import (
	"blog-example/lib/jwt/inf"
	"encoding/json"
	"fmt"
	"github.com/go-jose/go-jose/v4"
	"log"
	"time"
)

var jsonClaims []byte

func init() {
	claims := map[string]interface{}{
		"user_name": "huizhou92",
		"age":       30,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	jsonClaims, _ = json.Marshal(claims)
}

type GoJoseToken struct {
	User inf.UserInfo
}

func NewGoJoseToken(name string, age int) *GoJoseToken {
	return &GoJoseToken{
		User: inf.UserInfo{
			UserName: name,
			Age:      age,
		},
	}
}

func (gjt *GoJoseToken) GetJwtToken() (string, error) {

	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{
		Algorithm: cert.Alg,
		Key:       inf.SecretKey,
	}, nil)
	if err != nil {
		log.Fatalf("创建加密器失败: %v", err)
	}
	object, err := encrypter.Encrypt(jsonClaims)
	if err != nil {
		log.Fatalf("Failed to encrypt data: %v", err)
	}
	serialized, err := object.CompactSerialize()
	if err != nil {
		log.Fatalf("序列化加密对象失败: %v", err)
	}
	fmt.Printf("加密后: %s\n", serialized)
	return serialized, nil
}
