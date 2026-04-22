package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Content string
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Baza bilan bog'lanishda xato!")
	}
	db.AutoMigrate(&Message{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"Message": "Zar, men seni juda yaxshi ko'raman!"})
	})

	r.POST("/reply", func(c *gin.Context) {
		reply := c.PostForm("reply_text")
		db.Create(&Message{Content: reply})
		c.HTML(http.StatusOK, "index.html", gin.H{"Message": "Xabaringiz yuborildi! ❤️"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
