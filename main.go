package main

import (
	"github.com/adeemgoogle/gowork/Serv"
	"github.com/adeemgoogle/gowork/mypackage"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	var cityName string = "Almaty"
	mypackage.CurrentData(cityName)

	DB := Serv.Init()

	Serv.Close(DB)
}



