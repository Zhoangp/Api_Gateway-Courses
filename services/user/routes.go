package user

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/middleware"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/internal/delivery/http"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, c *config.Config, middleware *middleware.MiddleareManager) {
	client := NewUserServiceClient(c)
	userHandler := http.NewUserHandler(c, client)
	routes := r.Group("/user")
	routes.POST("/otp", userHandler.GetTokenResetPass())
	routes.PUT("/password/:token", middleware.RequireVerifyToken(), userHandler.ResetPassword())

	routes.Use(middleware.RequiredAuth())
	routes.PUT("/", userHandler.Update())
	routes.PATCH("/avatar", userHandler.UpdateAvatar())

	routes2 := r.Group("/instructor")
	routes2.Use(middleware.RequiredAuth())
	routes2.POST("/", userHandler.NewInstructor())

}
