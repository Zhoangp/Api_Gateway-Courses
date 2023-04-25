package http

import (
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
	"github.com/gin-gonic/gin"
)

func (hdl CourseHandler) GetCourse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		res, err := hdl.client.GetCourse(ctx, &pb.GetCourseRequest{Id: id})
		if err != nil {
			panic(err)
		}
		if res.Error != nil {
			panic(res.Error)
		}
		ctx.JSON(200, res)
	}
}
