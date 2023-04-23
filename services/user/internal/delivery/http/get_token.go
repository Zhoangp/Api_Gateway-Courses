package http

import (
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/internal/model"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/pb"
	"github.com/gin-gonic/gin"
)

func (hdl UserHandler) GetTokenResetPass() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data model.GetTokenRequest
		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}
		res, err := hdl.client.GetTokenResetPass(ctx, &pb.GetTokenResetPassRequest{
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
}
