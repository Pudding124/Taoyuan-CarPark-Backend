package service

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"net/http"
	"taoyuan_carpark/cronjob"
	"taoyuan_carpark/logging"
	"taoyuan_carpark/park"
)

const routePrefix = "/car_park"

func Init(c *gin.Engine) {

	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	logging.Print(zerolog.InfoLevel, "123", "123", "123")

	v1 := c.Group(routePrefix)
	{
		v1.GET("/healthy", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
		v1.GET("/find", func(c *gin.Context) {
			result := park.Find(c)
			c.JSON(http.StatusOK, result)
		})
	}

	// setting config
	settingConfig()

	// enable update carpark
	cronjob.EnableSchedule()
}

func settingConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic("viper read config error：" + err.Error())
	}
}

func NewGin() *gin.Engine {
	app := gin.New()
	app.Use(cors.Default())

	return app
}
