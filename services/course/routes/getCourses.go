package routes

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/pkg/common"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
	"github.com/gin-gonic/gin"
)

type HandleCourse struct {
	cf     *config.Config
	client pb.CourseServiceClient
}

func NewCourseHandler(cf *config.Config, client pb.CourseServiceClient) *HandleCourse {
	return &HandleCourse{cf, client}
}
func (hdl HandleCourse) GetCoursesWithPagination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var a common.Paging
		if err := ctx.ShouldBind(&a); err != nil {
			panic(common.ErrBadRequest(err))
		}
		a.FullFill()
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
