package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
	"github.com/gin-gonic/gin"
)

func (hdl *CourseHandler) UpdateCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userId := c.MustGet("userId").(string)

		var course pb.UpdateCourseRequest
		if err := c.ShouldBind(&course); err != nil {
			panic(err)
		}
		course.CourseId = id
		course.InstructorId = userId
		fmt.Println(course.InstructorId)
		res, err := hdl.client.UpdateCourse(c, &course)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if res.Error != nil {
			fmt.Println(res.Error)
			panic(res.Error)
		}
		c.JSON(200, gin.H{"code": 200, "message": "update sucessfully"})
	}
}
