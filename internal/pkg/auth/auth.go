package auth

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/theruziev2004/bookstore/internal/pkg"

	"github.com/golang-jwt/jwt/v4"
)

const userClaimGet = "USER_CLAIM"

var jwtKey = ""

func SetJWTKey(secret string) {
	if secret == "" {
		log.Fatal("jwt key is empty")
	}
	jwtKey = secret
}

func GetJWTSecret() []byte {
	if jwtKey == "" {
		log.Fatal("jwt key is empty")
	}
	return []byte(jwtKey)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		tokenString = strings.TrimSpace(tokenString)
		// Parse the token
		var claim jwt.RegisteredClaims
		_, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return GetJWTSecret(), nil
		})

		if err != nil {
			pkg.SendString(w, http.StatusBadRequest, "auth error")
			return
		}

		r = r.WithContext(WriteTokenToContext(r.Context(), claim))
		next.ServeHTTP(w, r)
	})
}

func WriteTokenToContext(ctx context.Context, claim jwt.RegisteredClaims) context.Context {
	return context.WithValue(ctx, userClaimGet, claim)
}

func ReadTokenFromContext(ctx context.Context) jwt.RegisteredClaims {
	claim := ctx.Value(userClaimGet).(jwt.RegisteredClaims)
	return claim
}
