package routes

import (
	"github.com/Zhoangp/Api_Gateway-Courses/internal/user-service/pb"
	"github.com/gin-gonic/gin"
)

type changePassRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func ResetPassword(ctx *gin.Context, client pb.UserServiceClient) {
	var data changePassRequest
	if err := ctx.ShouldBind(&data); err != nil {
		panic(err)
	}
	email := ctx.MustGet("emailUser").(string)
	pass := ctx.MustGet("password").(string)
	res, err := client.ChangePassword(ctx, &pb.ChangePasswordRequest{
		Email:       email,
		Password:    pass,
		NewPassword: data.NewPassword,
	})

	if err != nil {
		panic(err)
	}
	if res.Error != nil {
		panic(res.Error)
	}
	ctx.JSON(200, gin.H{"Message": "Change Password Success!"})
}
