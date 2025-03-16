package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandleError(ginC *gin.Context, statusCode int, message string, err error) {
	if err != nil {
		message = fmt.Sprintf("%s: %v", message, err)
	}
	ginC.IndentedJSON(statusCode, gin.H{"message": message})
}