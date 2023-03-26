package main

import (
	"GOHR/server/db"
	"GOHR/server/handler"
	"GOHR/server/main_service"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	_ "modernc.org/sqlite"
	"net/http"
	"time"
)

func main() {

	parseConfig() //viper

	dbInterface, dbClose := db.New(viper.GetString("db.sqlight_path"), viper.GetString("db.sqlight_driver"))
	defer dbClose()

	mainServiceInterface := main_service.New(dbInterface)
	handlerInterface := handler.New(mainServiceInterface)

	api := prepareAPIServer(viper.GetString("server.gin_mode"))

	setAPIRouts(api, handlerInterface)

	fmt.Println("http://127.0.0.1:9090")
	err := api.Run(viper.GetString("server.port"))
	if err != nil {
		fmt.Println("ERROR: RUN GIN SERVER %S", err.Error())
	}
}

func setAPIRouts(api *gin.Engine, handlerInterface handler.HandlerInterface) {
	api.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "Oops, NOT FOUND 404")
	})

	user := api.Group("user")
	user.GET("get_all", handlerInterface.GetAllUsers)
	user.POST("add_new", handlerInterface.AddNewUser)

	hr := api.Group("hr")
	hr.POST("add_new", handlerInterface.AddNewHR)

	web := api.Group("web", handlerInterface.AddNewHR)
	web.GET("/", handlerInterface.Index)
	web.GET("/profile", handlerInterface.Profile)
	web.GET("/logout", handlerInterface.Logout)
}

func parseConfig() {
	viper.SetConfigName("configs")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("ERROR READ CONFIG")
	}
}

func prepareAPIServer(serverMode string) *gin.Engine {
	gin.SetMode(serverMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output:    nil,
		SkipPaths: []string{"/favicon.ico"},
		Formatter: func(param gin.LogFormatterParams) string {
			if param.Latency > time.Minute {
				param.Latency = param.Latency.Truncate(time.Second)
			}
			return fmt.Sprintf("%v |%s %3d %s| %13v | %15s |%s %-7s %s %#v %s \n",

				param.TimeStamp.Format("2006/01/02 - 15:04:05"),
				"\u001B[97;106m\u001B[90;30m", param.StatusCode, "\033[0m",
				param.Latency,
				param.ClientIP,
				"\u001B[97;106m\u001B[97;30m", param.Method, "\033[0m",
				param.Path,
				param.ErrorMessage,
			)
		},
	}))
	r.Static("/assets/", "./server/web/assets/")
	r.LoadHTMLGlob("./server/web/templates/*")
	return r
}
