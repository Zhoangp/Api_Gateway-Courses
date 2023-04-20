package routes

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/pkg/common"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/pb"
	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	common.SQLModel
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

func Update(ctx *gin.Context, client pb.UserServiceClient) {
	var user updateRequest
	email := ctx.MustGet("emailUser").(string)
	if err := ctx.ShouldBind(&user); err != nil {
		panic(err)
	}

	res, err := client.UpdateUser(ctx, &pb.UpdateUserRequest{
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
