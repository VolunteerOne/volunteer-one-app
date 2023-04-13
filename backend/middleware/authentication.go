package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func BasicAuth(c *gin.Context) {
    // Get the token off the header
	accessToken, ok := c.Request.Header["Token"]

	if !ok {
		fmt.Println(`No "token" field in header`)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/validate it
	token, err := jwt.Parse(accessToken[0], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Println("Error: Something went wrong when parsing the token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

        // Check if refresh
        if claims["type"] == "refresh" {
            log.Println("Cannot use refresh token for normal authentication")
            c.JSON(http.StatusUnauthorized, gin.H{
                "message":       "Cannot use refresh token for normal authentication",
                "success":       false,
            })
            c.AbortWithStatus(http.StatusUnauthorized)
        }

		// Check if expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			log.Println("Access token is expired")
            c.JSON(http.StatusUnauthorized, gin.H{
                "message":       "Access token is expired",
                "success":       false,
            })
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		log.Println("good token")

		// // Find the user with token "user"
		// var user models.User
		// initializers.DB.First(&user, claims["sub"])
		//
		// if user.ID == 0 {
		// 	fmt.Println("Error in user.id")
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// }
		//
		// // Attach to request
		// c.Set("user", user)

		// Continue
		c.Next()
	} else {
		log.Println("Invalid token")
        c.JSON(http.StatusUnauthorized, gin.H{
            "message":       "Invalid token",
            "success":       false,
        })
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
