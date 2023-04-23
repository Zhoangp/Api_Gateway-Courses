package course

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/internal/delivery/http"
	"github.com/Zhoangp/Api_Gateway-Courses/services/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCourseService(r *gin.Engine, cf *config.Config, middleware *middleware.MiddleareManager) {

	client := NewCourseServiceClient(cf)
	hdlCourse := http.NewCourseHandler(cf, client)
	router := r.Group("/courses")
	router.Use(middleware.RequiredAuth())
	router.POST("", hdlCourse.GetCoursesWithPagination())
}
