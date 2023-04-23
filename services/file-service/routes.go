package file_service

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/file-service/routes"
	"github.com/Zhoangp/Api_Gateway-Courses/services/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterFileRoute(r *gin.Engine, ctx *config.Config, middleware *middleware.MiddleareManager) {
	svc := ServiceClient{
		Client: InitServiceClient(ctx),
	}

	route := r.Group("/file")
	route.Use(middleware.RequiredAuth())
	route.POST("", svc.Upload)
}
func (svc *ServiceClient) Upload(ctx *gin.Context) {
	routes.Upload(ctx, svc.Client)
}
