package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
	"github.com/gin-gonic/gin"
)

func (hdl *CourseHandler) DeleteCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userId := c.MustGet("userId").(string)

		req := &pb.DeleteCourseRequest{
			CourseId: id,
			UserId:   userId,
		}

		res, err := hdl.client.DeleteCourse(c, req)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		if res.Error != nil {
			panic(res.Error)
		}

		c.JSON(200, gin.H{"code": 200, "message": "delete sucessfully"})
	}
}
