package routes

import (
	"context"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/auth-service/pb"
	"github.com/gin-gonic/gin"
)

type registerRequestBody struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Role        string `json:"role"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := registerRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		panic(err)
	}
	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		PhoneNumber: body.PhoneNumber,
		Email:       body.Email,
		Password:    body.Password,
		Address:     body.Address,
		Role:        body.Role,
	})

	if err != nil {
		panic(err)
	}
	if res.Error != nil {
		panic(res.Error)
	}

	ctx.JSON(200, gin.H{"Message": "Create user-service successfully"})
}
