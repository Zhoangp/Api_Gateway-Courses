package http

import (
	"errors"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
	"github.com/gin-gonic/gin"
)

func (hdl *CourseHandler) GetCourseContent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.MustGet("userId").(string)
		courseId := ctx.Param("courseId")
		if courseId == "" {
			panic(errors.New("need course id"))
		}
		res, err := hdl.client.GetCourseContent(ctx, &pb.GetCourseContentRequest{
			UserId:   userId,
			CourseId: courseId,
		})
		if err != nil {
			panic(err)
		}
		if res.Error != nil {
			panic(res.Error)
		}
		ctx.JSON(200, res)
	}
}
