package jwt_v5

import (
	"blog-example/lib/jwt/inf"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type V5Token struct {
	User inf.UserInfo
	jwt.RegisteredClaims
}

func NewV5Token(name string, age int) *V5Token {
	return &V5Token{
		User: inf.UserInfo{
			UserName: name,
			Age:      age,
		},
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24))},
	}
}
func (v5 *V5Token) GetToken() (string, error) {
	v5.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, v5)
	return token.SignedString(inf.SecretKey)
}

func (v5 *V5Token) ParseToken(token string) (*inf.UserInfo, error) {
	t, err := jwt.ParseWithClaims(token, &V5Token{}, func(token *jwt.Token) (interface{}, error) {
		return inf.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	v5, ok := t.Claims.(*V5Token)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return &v5.User, nil
}

func main() {
	v5 := NewV5Token("test", 18)
	token, _ := v5.GetToken()
	fmt.Println(token)
	user, _ := v5.ParseToken(token)
	fmt.Println(user)
}
