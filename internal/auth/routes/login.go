package routes

import (
	"context"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/auth/pb"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		panic(err)
	}
	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})
	if err != nil {
		panic(err)
	}
	if res.Error != nil {
		panic(res.Error)
	}
	ctx.SetCookie("refresh_token", res.RefreshToken, 3600*720, "/", "localhost", false, true)
	ctx.JSON(200, &res)

}
