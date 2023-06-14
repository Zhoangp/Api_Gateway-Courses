package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/pkg/common"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
	"github.com/gin-gonic/gin"
)

func (hdl CourseHandler) GetCourse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		fmt.Println(id)
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
func (hdl CourseHandler) GetCourses() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var a common.Paging
		if err := ctx.ShouldBind(&a); err != nil {
			panic(common.ErrBadRequest(err))
		}
		a.FullFill()
		res, err := hdl.client.GetCourses(ctx, &pb.GetCoursesRequest{
			Page:     int32(a.Page),
			PageSize: int32(a.Limit),
		})
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if res.Error != nil {
			panic(err)
		}
		ctx.JSON(200, res)
	}
}
