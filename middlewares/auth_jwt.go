package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/patipan-patisampita/gin-framework9/configs"
	"github.com/patipan-patisampita/gin-framework9/models"
)

func AuthJWT() gin.HandlerFunc {

	return gin.HandlerFunc(func(c *gin.Context) {

		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			defer c.AbortWithStatus(http.StatusUnauthorized)
		}

		tokenHeader := c.GetHeader("Authorization")
		accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
		// fmt.Println(accessToken)
		jwtSecretKey := os.Getenv("JWT_SECRET")

		token, _ := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			defer c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			// global value result
			claims := token.Claims.(jwt.MapClaims)
			var user models.User
			configs.DB.First(&user, claims["user_id"])
			c.Set("user", user)
			// return to next method if token is exist
			c.Next()
		}

	})
}
