package http

import (
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/internal/model"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/pb"
	"github.com/gin-gonic/gin"
)

func (hdl UserHandler) NewInstructor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.MustGet("userId").(string)
		var data model.NewInstrcutorRequest
		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}
		res, err := hdl.client.NewInstructor(ctx, &pb.NewInstructorRequest{
			UserId:   userId,
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
		ctx.JSON(200, gin.H{"message": "Create successfully!"})
	}
}
