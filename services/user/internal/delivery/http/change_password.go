package http

import (
	"errors"
	"github.com/Zhoangp/Api_Gateway-Courses/pkg/common"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/internal/model"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/pb"
	"github.com/gin-gonic/gin"
)

func (hdl UserHandler) ResetPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data model.ChangePassRequest
		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}
		email := ctx.MustGet("emailUser").(string)
		pass := ctx.MustGet("password").(string)
		key := ctx.MustGet("key").(string)
		if key != "forget" {
			panic(common.NewCustomError(errors.New("invalid token"), "invalid token"))
		}

		res, err := hdl.client.ChangePassword(ctx, &pb.ChangePasswordRequest{
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

}
