package main

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/auth-service"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/file-service"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/middleware"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/user-service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cf, err := config.LoadConfig("config/config-local.yml")
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	mdware := middleware.NewMiddlewareManager(cf)
	r := gin.Default()
	r.Use(mdware.Recover())

	_ = *auth_service.RegisterAuthRoutes(r, cf)
	user_service.RegisterUserRoutes(r, cf)
	file_service.RegisterFileRoute(r, cf)
	r.Run(cf.Services.Port)
}
