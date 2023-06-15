package main

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/pkg/utils"
	"github.com/Zhoangp/Api_Gateway-Courses/services/auth"
	"github.com/Zhoangp/Api_Gateway-Courses/services/cart"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course"
	file_service "github.com/Zhoangp/Api_Gateway-Courses/services/file-service"
	"github.com/Zhoangp/Api_Gateway-Courses/services/middleware"
	"github.com/Zhoangp/Api_Gateway-Courses/services/payment"
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
	if env == "railway" {
		fileName = "config/config-railway.yml"
	} else if env == "app" {
		fileName = "config/config-app.yml"
	}

	cf, err := config.LoadConfig(fileName)
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	//Run migrations when deploy the application
	if env == "app" || env == "railway" {
		host := os.Getenv("HOST")
		port := os.Getenv("PORT")
		user := os.Getenv("USER")
		password := os.Getenv("PASS")
		dbName := os.Getenv("DBNAME")
		cf.Mysql.Host = host
		cf.Mysql.Port = port
		cf.Mysql.User = user
		cf.Mysql.Password = password
		cf.Mysql.DBName = dbName
		utils.RunDBMigration(cf)
	}
	mdware := middleware.NewMiddlewareManager(cf)

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
	payment.NewPaymentRoutes(r, cf, mdware)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if env == "railway" {
		r.Run("0.0.0.0:" + port)
	} else {
		r.Run(cf.Services.Port)
	}
}
