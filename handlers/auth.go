package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

var allowedGroups = []string{"/goappreader", "/goappwriter"}

func KeycloakAuthMiddleware() gin.HandlerFunc {
	provider, err := oidc.NewProvider(context.Background(), "http://keycloak:8080/realms/myrealm")
	if err != nil {
		panic(err)
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: "go-app"})

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// Check if token has "Bearer " prefix
		token = strings.TrimPrefix(token, "Bearer ")
		token = strings.Trim(token, "\"")
		idToken, err := verifier.Verify(context.Background(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid token:%v", err.Error())})
			return
		}
		var claims struct {
			Groups []string `json:"groups"`
		}
		if err := idToken.Claims(&claims); err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("Failed to parse token claims:%v", err.Error())})
			return
		}
		fmt.Println(claims)
		// Check if user has at least one required group
		authorized := false
		for _, userGroup := range claims.Groups {
			for _, allowedGroup := range allowedGroups {
				if userGroup == allowedGroup {
					authorized = true
					break
				}
			}
			if authorized {
				break
			}
		}
		if !authorized {
			c.AbortWithStatusJSON(http.StatusForbidden,
				gin.H{"error": "Access denied",
					"required_groups": allowedGroups,
					"user_groups":     claims.Groups})
			return
		}
		c.Set("userGroups", claims.Groups)
		c.Next()
	}
}
