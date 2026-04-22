package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "Ishlayapti",
			"message": "25-aprel kuni ishga muvaffaqiyatli borishingizni tilayman!",
		})
	})

	// Global portni olish (Render uchun shart)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
