package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type example struct {
	message    string
	otherField bool
}

func ExampleHandler(c *gin.Context) {

	respObject := &example{
		message:    "Hello",
		otherField: true,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": respObject.message,
	})
}
