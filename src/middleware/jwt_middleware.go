package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var defaultKey = []byte("default-secret")

func getKey() []byte {
	if k := os.Getenv("JWT_SECRET"); k != "" {
		return []byte(k)
	}
	return defaultKey
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authH := c.GetHeader("Authorization")
		if !strings.HasPrefix(authH, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid auth header"})
			return
		}
		tokenStr := strings.TrimPrefix(authH, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return getKey(), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}
		uidF, ok := claims["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id"})
			return
		}
		c.Set("user_id", uint(uidF))
		c.Next()
	}
}
