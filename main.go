package main

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/auth"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/middleware"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/user"
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
	_ = *auth.RegisterAuthRoutes(r, cf)
	user.RegisterUserRoutes(r, cf)
	r.Run(cf.Services.Port)
}
