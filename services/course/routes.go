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
	routerCategory := r.Group("/categories")
	routerCategory.GET("", hdlCourse.GetCategories())
	router := r.Group("/courses")
	router.GET("", hdlCourse.GetCourses())
	router.GET("/:id", hdlCourse.GetCourse())
	router.Use(middleware.RequiredAuth())
	router.PUT("/:id", hdlCourse.UpdateCourse())
	router.PATCH("/publish/:id", hdlCourse.PublishCourse())

	router.POST("", hdlCourse.CreateCourse())
	router.POST("/content/:courseId", hdlCourse.GetCourseContent())
	router.POST("/:courseId/enrollment", hdlCourse.Enroll())
	router.GET("/enrollment", hdlCourse.GetEnrollments())
	router.GET("/prices", hdlCourse.GetListPrice())
	router.GET("/owner", hdlCourse.GetCoursesWithInstructor())
	router.DELETE("/:id", hdlCourse.DeleteCourse())

}
