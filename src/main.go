package main

import (
	"context"
	"github.com/adeemgoogle/gowork/src/service/location"
	"github.com/adeemgoogle/gowork/src/service/user"
	"io/ioutil"
	"log"
	"os/signal"
	"syscall"

	httpClient "github.com/adeemgoogle/gowork/src/common/http"
	"github.com/adeemgoogle/gowork/src/config"
	"github.com/adeemgoogle/gowork/src/database"
	"github.com/adeemgoogle/gowork/src/database/drivers"
	"github.com/adeemgoogle/gowork/src/handler"
	"github.com/adeemgoogle/gowork/src/service/weather"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
)

// main - основная функция
func main() {
	var conf config.Config
	if err := env.Parse(&conf); err != nil {
		log.Fatalf("[ERROR] while parsing env variables to config struct: %v", err)
	}

	h := InitHandler(conf)
	r := InitRouter(h, conf.ReleaseMode)
	r.Run(conf.ServerPort)
}

// InitHandler - инициализация обработчика
func InitHandler(conf config.Config) *handler.Handler {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ds, err := database.New(drivers.Config{
		DBName: conf.DBName,
		DSName: conf.DSName,
		URL:    conf.DBURL,
	})

	if err = ds.Connect(ctx); err != nil {
		log.Fatalf("[ERROR] while connecting to a database: %s, %v", conf.DSName, err)
	}

	err = ds.RunMigrations()
	if err != nil {
		log.Fatalf("[ERROR] migaration: %v", err)
	}

	weatherCli := httpClient.NewClient(conf.WeatherBaseURL)
	weatherSrv := weather.NewService(conf, ds.WeatherRepository(), weatherCli)
	userSrv := user.NewService(conf, ds.UserRepository())
	locationSrv := location.NewService(conf, ds.LocationRepository())

	h := handler.NewHandler(weatherSrv, userSrv, locationSrv, conf)
	return h
}

// InitRouter - инициализация ендпойнтов
func InitRouter(h *handler.Handler, production bool) *gin.Engine {
	if production {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	// создание экземпляра приложения с предварительно настроенными параметрами
	// этот экземпляр приложения содержит роутер, мидлвары и другие компоненты, необходимые для обработки HTTP-запросов
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// ендпойнт для проверки работоспособности приложения
	r.GET("/healthz")

	// группа /api/v1 содержит ендпойнты первой версии
	apiv1 := r.Group("/api/v1")
	apiv1.GET("/weather", h.GetAllWeatherData)
	apiv1.POST("/user", h.CreateUser)
	apiv1.GET("/user", h.GetUser)
	apiv1.GET("/locations", h.GetLocations)
	return r
}
