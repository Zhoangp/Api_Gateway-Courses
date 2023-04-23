package course

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/routes"
	"github.com/Zhoangp/Api_Gateway-Courses/services/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCourseService(r *gin.Engine, cf *config.Config) {
	mdware := middleware.NewMiddlewareManager(cf)

	client := NewCourseServiceClient(cf)
	hdlCourse := routes.NewCourseHandler(cf, client)
	router := r.Group("/courses")
	router.Use(mdware.RequiredAuth())
	router.POST("", hdlCourse.GetCoursesWithPagination())
}
