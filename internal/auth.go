package internal

import (
	"errors"

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
		// Get Cookie value
		cookie, err := c.Cookie("access_token")
		if err != nil {
			// Invalid JWT, allow the request to continue
			c.Set("valid_jwt", false) // Flag indicating invalid JWT
			c.Next()
			return
		}

		// Validate JWT
		claims, err := validateToken(cookie)
		if err != nil {
			// Invalid JWT, allow the request to continue
			c.Set("valid_jwt", false) // Flag indicating invalid JWT
			c.Next()
			return
		}
		// Attach user ID to context
		c.Set("user_id", claims.ExpiresAt)
		c.Set("valid_jwt", true)
		c.Next()
	}
}

// validateToken verifies the JWT and extracts claims
func validateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(JWT_TOKEN), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
