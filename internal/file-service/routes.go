package file_service

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/file-service/routes"
	"github.com/gin-gonic/gin"
)

func RegisterFileRoute(r *gin.Engine, ctx *config.Config) {
	svc := ServiceClient{
		Client: InitServiceClient(ctx),
	}
	route := r.Group("/file")
	route.POST("/", svc.Upload)
}
func (svc *ServiceClient) Upload(ctx *gin.Context) {
	routes.Upload(ctx, svc.Client)
}