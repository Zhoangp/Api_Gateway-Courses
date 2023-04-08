package routes

import (
	"github.com/Zhoangp/Api_Gateway-Courses/internal/user/pb"
	"github.com/gin-gonic/gin"
)

type getOtpData struct {
	Email string `json:"email"`
}

func GetTokenResetPass(ctx *gin.Context, client pb.UserServiceClient) {
	var data getOtpData
	if err := ctx.ShouldBind(&data); err != nil {
		panic(err)
	}
	res, err := client.GetTokenResetPass(ctx, &pb.GetTokenResetPassRequest{
		Email: data.Email,
	})
	if err != nil {
		panic(err)
	}
	if res.Error != nil {
		panic(res.Error)
	}
	ctx.JSON(200, gin.H{"Message": "Success"})
}
