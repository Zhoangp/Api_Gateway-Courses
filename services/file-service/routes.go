package file_service

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/file-service/internal/http"
	"github.com/Zhoangp/Api_Gateway-Courses/services/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterFileRoute(r *gin.Engine, cf *config.Config, middleware *middleware.MiddleareManager) {

	hdl := http.NewFileHandler(cf, InitServiceClient(cf))
	route := r.Group("/file")
	route.Use(middleware.RequiredAuth())
	route.POST("", hdl.UploadAvatar())
	route.POST("/asset", hdl.UploadAsset())

}
