package routes

import (
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/pb"
	"github.com/gin-gonic/gin"
)

type requestInstrcutor struct {
	Website  string `json:"website"`
	LinkedIn string `json:"linkedin"`
	Youtube  string `json:"youtube"`
	Bio      string `json:"bio"`
}

func NewInstructor(c *gin.Context, client pb.UserServiceClient) {
	email := c.MustGet("emailUser").(string)
	var data requestInstrcutor
	if err := c.ShouldBind(&data); err != nil {
		panic(err)
	}
	res, err := client.NewInstructor(c, &pb.NewInstructorRequest{
		Email:    email,
		Website:  data.Website,
		Linkedin: data.LinkedIn,
		Youtube:  data.Youtube,
		Bio:      data.Bio,
	})
	if err != nil {
		panic(err)
	}
	if res.Error != nil {
		panic(res.Error)
	}
	c.JSON(200, gin.H{"message": "Create successfully!"})
}
