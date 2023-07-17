package main

import (

	"gorm.io/gorm"
	_ "net/http"
	"github.com/adeemgoogle/gowork/mypackage"
	
	
	// 	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func main() {
	var cityName string = "Almaty"
	 mypackage.CurrentData(cityName)

}
