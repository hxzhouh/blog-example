package inf

type UserInfo struct {
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	//jwt.RegisteredClaims
}

var SecretKey []byte

func init() {
	SecretKey = []byte("12345612345612341234561234561234")
}

type Token interface {
	GetToken() (string, error)
	ParseToken(token string) (*UserInfo, error)
}
