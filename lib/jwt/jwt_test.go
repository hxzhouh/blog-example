package main

import (
	"blog-example/lib/jwt/jose2go"
	"blog-example/lib/jwt/jwt_v5"
	"testing"
)

// bench v5

func BenchmarkCreateToken_Jwt_v5(b *testing.B) {
	v5Token := jwt_v5.NewV5Token("test", 18)

	for i := 0; i < b.N; i++ {
		v5Token.GetToken()
	}
}

func BenchmarkParseToken_Jwt_v5(b *testing.B) {
	v5Token := jwt_v5.NewV5Token("test", 18)
	token, _ := v5Token.GetToken()

	for i := 0; i < b.N; i++ {
		v5Token.ParseToken(token)
	}
}

// bench jose2go

func BenchmarkCreateToken_Jose2go(b *testing.B) {
	joseToken := jose2go.NewJoseToken("test", 18)
	for i := 0; i < b.N; i++ {
		joseToken.GetToken()
	}
}

func BenchmarkParseToken_Jose2go(b *testing.B) {
	joseToken := jose2go.NewJoseToken("test", 18)
	token, _ := joseToken.GetToken()

	for i := 0; i < b.N; i++ {
		joseToken.ParseToken(token)
	}
}
