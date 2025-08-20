package internal

import (
	"encoding/base64"
	"errors"
	"log"

	"github.com/Techeer-Hogwarts/search/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims represents the expected claims in the token
type JWTClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

var JWT_TOKEN string

func init() {
	JWT_TOKEN = config.GetEnvVarAsString("JWT_SECRET", "some_secret_key")
}

// ValidateJWT middleware checks the JWT token from cookies
func ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("access_token")
		if err != nil {
			log.Printf("JWT cookie: %s", cookie)
			c.Set("valid_jwt", false)
			c.Next()
			return
		}

		// Validate JWT
		claims, err := validateToken(cookie)
		log.Printf("JWT claims: %v", claims)
		if err != nil {
			log.Printf("Invalid JWT: %v", err)
			c.Set("valid_jwt", false)
			c.Next()
			return
		}
		c.Set("user_id", claims.ExpiresAt)
		c.Set("valid_jwt", true)
		c.Next()
	}
}

// validateToken verifies the JWT and extracts claims
func validateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		log.Printf("JWT header: %+v", token.Header)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}
		decoded, err := base64.StdEncoding.DecodeString(JWT_TOKEN)
		if err != nil {
			log.Printf("Failed to base64 decode JWT secret: %v", err)
			return nil, errors.New("invalid secret")
		}

		return decoded, nil
	})

	if err != nil || !token.Valid {
		log.Printf("JWT parse token: %v", token)
		log.Printf("JWT parse error: %v", err)
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
