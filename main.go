package main

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/pkg/utils"
	"github.com/Zhoangp/Api_Gateway-Courses/services/auth"
	"github.com/Zhoangp/Api_Gateway-Courses/services/cart"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course"
	file_service "github.com/Zhoangp/Api_Gateway-Courses/services/file-service"
	"github.com/Zhoangp/Api_Gateway-Courses/services/middleware"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	//Get environment variable
	env := os.Getenv("ENV")
	//Load config
	fileName := "config/config-local.yml"
	if env == "app" {
		fileName = "config/config-app.yml"
	}
	cf, err := config.LoadConfig(fileName)
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	// Run migrations when deploy the application
	if env == "app" {
		utils.RunDBMigration(cf)
	}
	//create middleware manager instance
	mdware := middleware.NewMiddlewareManager(cf)

	//config cors middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"authorization", "content-type"}
	config.AllowFiles = true

	//new instance of the Gin web framework
	r := gin.Default()

	//apply middleware for all requests
	r.Use(cors.New(config), mdware.Recover())

	//define routers for each service
	auth.RegisterAuthRoutes(r, cf, mdware)
	user.RegisterUserRoutes(r, cf, mdware)
	file_service.RegisterFileRoute(r, cf, mdware)
	course.RegisterCourseService(r, cf, mdware)
	cart.NewCartRoutes(r, cf, mdware)

	r.Run(cf.Services.Port)
}
