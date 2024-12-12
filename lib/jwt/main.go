package main

import (
	"blog-example/lib/jwt/jose2go"
	"fmt"
)

func main() {
	joseToken := jose2go.NewJoseToken("test", 18)
	token, _ := joseToken.GetToken()
	fmt.Println(token)
}
