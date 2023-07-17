package Serv

import (

	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"github.com/adeemgoogle/gowork/mypackage"
)

type City struct {
	gorm.Model
	ID      int     `json:"id"`
	Name    string  `json:"Name"`
	Country string  `json:"Country"`
	Lon     float64 `json:"Lon"`
	Lat     float64 `json:"Lat"`
}

var DB *gorm.DB

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	var err error
	dsn := "host=localhost user=postgres password=123 dbname=go port=5432 sslmode=disable "

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")

	DB.AutoMigrate(&mypackage.Current{})
	// DB.AutoMigrate(&Air_quality{})

}

func Server() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/createCity", func(c *gin.Context) {

		var city City
		if err := c.ShouldBindJSON(&city); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.Create(&city)
		c.JSON(200, gin.H{
			"city": "created",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080

}
