package book

import (
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All books",
	})
}

func GetBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All books",
	})
}

func NewBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All books",
	})
}

func DeleteBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All books",
	})
}
