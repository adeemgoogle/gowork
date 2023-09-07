package config

// Config - конфигурация
type Config struct {
	ServerPort  string `env:"SERVER_PORT" envDefault:":8080"`
	ReleaseMode bool   `env:"RELEASE_MODE" envDefault:"false"`

	DBName string `env:"DATABASE_NAME" envDefault:"gowork"`
	DSName string `env:"DS_NAME" envDefault:"postgres"`
	DBURL  string `env:"DB_URL" envDefault:"postgres://postgres:postgres@localhost:5432/gowork?sslmode=disable"`

	//DBURL  string `env:"DB_URL" envDefault:"postgres://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable" envExpand:"true"`

	WeatherBaseURL  string `env:"WEATHER_BASE_URL" envDefault:"https://pro.openweathermap.org/data/2.5"`
	WeatherImageURL string `env:"WEATHER_IMAGE_URL" envDefault:"https://openweathermap.org"`
	WeatherAppId    string `env:"WEATHER_APP_ID" envDefault:"51e51b22fb137270e2e89bd2bc7c4acc"`
}
