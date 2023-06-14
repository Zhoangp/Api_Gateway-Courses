package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
	"github.com/gin-gonic/gin"
)

func (hdl *CourseHandler) PublishCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userId := c.MustGet("userId").(string)

		publish := &pb.PublishCourseRequest{
			CourseId: id,
			UserId:   userId,
		}

		res, err := hdl.client.PublishCourse(c, publish)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		if res.Error != nil {
			panic(res.Error)
		}
		if res.ErrorsValidate != nil {
			c.AbortWithStatusJSON(400, gin.H{"message": "This course is not eligible for publication yet", "eligibles": res.ErrorsValidate})
		} else {
			c.JSON(200, gin.H{"code": 200, "message": "publish sucessfully"})
		}
	}
}
