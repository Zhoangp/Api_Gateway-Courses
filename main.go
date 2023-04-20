package main

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/pkg/utils"
	"github.com/Zhoangp/Api_Gateway-Courses/services/auth"
	file_service "github.com/Zhoangp/Api_Gateway-Courses/services/file-service"
	"github.com/Zhoangp/Api_Gateway-Courses/services/middleware"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	fileName := "config/config-local.yml"
	if env == "app" {
		fileName = "config/config-app.yml"
	}
	cf, err := config.LoadConfig(fileName)
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	if env == "app" {
		utils.RunDBMigration(cf)
	}
	mdware := middleware.NewMiddlewareManager(cf)
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"authorization", "content-type"}
	config.AllowFiles = true
	r.Use(cors.New(config), mdware.Recover())

	_ = *auth.RegisterAuthRoutes(r, cf)
	user.RegisterUserRoutes(r, cf)
	file_service.RegisterFileRoute(r, cf)
	r.Run(cf.Services.Port)
}
