package middleware

import (
	"context"
	"myapp/tool"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var CtxKey = &contextKey{"key"}

type contextKey struct {
	name string
}

type Employee struct {
	ID int `json:"id"`
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")
		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		authTokens := strings.Split(authToken, " ")
		if authTokens[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		jwtToken, err := tool.TokenValidate(authTokens[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, ok := jwtToken.Claims.(*tool.MyClaim)
		if !ok || !jwtToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		ctx := context.WithValue(c.Request.Context(), CtxKey, &Employee{
			ID: claims.EmployeeID,
		})

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func AuthContext(c *gin.Context) *Employee {
	raw, _ := c.Request.Context().Value(CtxKey).(*Employee)
	return raw
}
