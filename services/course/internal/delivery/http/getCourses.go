package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/pkg/common"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
	"github.com/gin-gonic/gin"
)

func (hdl CourseHandler) GetCourses() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var a common.Paging
		a.Limit = 0
		a.Page = 0
		res, err := hdl.client.GetCoursesWithPagination(ctx, &pb.GetCoursesPaginationRequest{
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
