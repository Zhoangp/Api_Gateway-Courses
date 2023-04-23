package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/internal/model"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/pb"
	"github.com/gin-gonic/gin"
)

func (hdl UserHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.UpdateUserRequest
		email := ctx.MustGet("emailUser").(string)
		if err := ctx.ShouldBind(&user); err != nil {
			panic(err)
		}

		res, err := hdl.client.UpdateUser(ctx, &pb.UpdateUserRequest{
			Id:          uint32(user.Id),
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       email,
			Password:    user.Password,
			PhoneNumber: user.Phone,
			Address:     user.Address,
		})
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if res.Error != nil {
			panic(res.Error)
		}
		ctx.JSON(200, gin.H{"Message": "Update User-Service successfully"})
	}
}
