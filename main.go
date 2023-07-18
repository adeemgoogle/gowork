package main

import (
	"github.com/adeemgoogle/gowork/Serv"
	// "github.com/adeemgoogle/gowork/mypackage"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

func main() {
	// var cityName string = "Almaty"
	// mypackage.CurrentData(cityName)

	DB := Serv.Init()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/createCity", func(c *gin.Context) {
		// Serv.Close(DB)


		var city int
		// if err := c.ShouldBindJSON(&city); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		DB.Create(&city)
		c.JSON(200, gin.H{
			"city" : "created",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080

	// Serv.Close(DB)
}



