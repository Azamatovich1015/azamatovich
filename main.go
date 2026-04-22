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
	// Render-dagi DATABASE_URL orqali ulanamiz
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Ma'lumotlar bazasiga ulanib bo'lmadi!")
	}

	// Xabarlar jadvalini yaratish
	db.AutoMigrate(&Message{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Message": "Zar, men seni juda yaxshi ko'raman!",
		})
	})

	// Javob yozilganda ishlaydigan qism