package main

import (
	controllers "github.com/adeemgoogle/gowork/src/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	controllers.PostCurrentData(r)
	controllers.PutCurrentData(r)
	controllers.PostHourlyData(r)
	controllers.PutHourlyData(r)
	controllers.PostDailyData(r)
	controllers.PutDailyData(r)
	controllers.RunAPI(r)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080

}
