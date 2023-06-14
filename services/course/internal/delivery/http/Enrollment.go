package http

import (
	"errors"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
	"github.com/gin-gonic/gin"
)

func (hdl *CourseHandler) Enroll() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("userId").(string)
		courseId := c.Param("courseId")
		if courseId == "" {
			panic(errors.New("need course id"))
		}
		res, err := hdl.client.Enrollment(c, &pb.EnrollmentRequest{
			UserId:   userId,
			CourseId: courseId,
		})
		if err != nil {
			panic(err)
		}
		if res.Error != nil {
			panic(res.Error)
		}
		c.JSON(200, gin.H{"message": "Enrollment successful!"})
	}
}

func (hdl *CourseHandler) GetEnrollments() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("userId").(string)
		res, err := hdl.client.GetEnrollments(c, &pb.GetEnrollmentsRequest{
			UserId: userId,
		})
		if err != nil {
			panic(err)
		}
		if res.Error != nil {
			panic(res.Error)
		}
		c.JSON(200, res)
	}
}
