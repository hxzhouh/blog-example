package jose2go

import (
	"blog-example/lib/jwt/inf"
	"encoding/json"
	jose "github.com/dvsekhvalnov/jose2go"
	"time"
)

type JoseToken struct {
	User inf.UserInfo
	Exp  int64
}

func NewJoseToken(name string, age int) *JoseToken {
	return &JoseToken{
		User: inf.UserInfo{
			UserName: name,
			Age:      age,
		},
	}
}

func (joseToken *JoseToken) GetToken() (string, error) {

	jose.Header("typ", "JWT")
	jose.Header("alg", "HS256")
	joseToken.Exp = time.Now().Add(time.Hour * 24).Unix()
	p, _ := json.Marshal(joseToken)

	token, err := jose.Sign(string(p), jose.HS256, inf.SecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (joseToken *JoseToken) ParseToken(token string) (*inf.UserInfo, error) {
	claims, _, err := jose.Decode(token, inf.SecretKey)
	if err != nil {
		return nil, err
	}
	var userInfo inf.UserInfo
	_ = json.Unmarshal([]byte(claims), &userInfo)
	return &userInfo, nil
}
