package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
	"github.com/gin-gonic/gin"
)

func (hdl *CourseHandler) CreateCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var course pb.CreateCourseRequest
		if err := c.ShouldBind(&course); err != nil {
			panic(err)
		}

		userId := c.MustGet("userId").(string)
		course.InstructorId = userId
		res, err := hdl.client.CreateCourse(c, &course)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if res.Error != nil {
			fmt.Println(res.Error)
			panic(res.Error)
		}
		c.JSON(200, res)

	}
}
