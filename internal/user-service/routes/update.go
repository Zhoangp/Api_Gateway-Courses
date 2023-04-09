package routes

import (
	"github.com/Zhoangp/Api_Gateway-Courses/internal/user-service/pb"
	"github.com/Zhoangp/Api_Gateway-Courses/pkg/common"
	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	common.SQLModel
	Email     string `json:"email" gorm:"column:email"`
	Password  string `json:"password" gorm:"column:password"`
	FirstName string `json:"firstName" gorm:"column:firstName"`
	LastName  string `json:"lastName" gorm:"column:lastName"`
	Phone     string `json:"phone" gorm:"column:phoneNumber"`
	Address   string `json:"address" gorm:"column:address"`
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
		panic(err)
	}
	if res.Error != nil {
		panic(res.Error)
	}
	ctx.JSON(200, gin.H{"Message": "Update user-service successfully"})
}
