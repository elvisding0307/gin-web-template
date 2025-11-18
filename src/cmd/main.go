package main

import (
	"gin-web-template/internal/config"
	"gin-web-template/internal/dao"
	"gin-web-template/internal/routers"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

// @title gin-web-template
// @version 1.0
// @description gin-web-template
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4387
// @BasePath /api/v1
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Init() {
	var (
		cfg *config.Config
		err error
	)
	if err = config.InitConfig(); err != nil {
		log.Fatalln("Unable to load the config.")
	}
	if err = dao.InitMysqlDB(); err != nil {
		log.Fatalln("Unable to initialize the mysql database.")
	}
	if err = dao.InitRedisDB(); err != nil {
		log.Fatalln("Unable to initialize the redis database.")
	}

	if cfg, err = config.ServerConfig(); err != nil {
		log.Println("无法获取服务器配置: ", err)
	}
	if strings.ToLower(cfg.ServerMode) == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	// Initialize the router
	router := routers.CreateRouter()

	// Start the server
	addr := cfg.GetBindAddr()
	log.Println("Server is running on ", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func main() {
	Init()
}
