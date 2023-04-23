package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/auth/internal/model"
	"github.com/Zhoangp/Api_Gateway-Courses/services/auth/pb"
	errpb "github.com/Zhoangp/Api_Gateway-Courses/services/error/pb"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	cf     *config.Config
	client pb.AuthServiceClient
}

func NewAuthHandler(cf *config.Config, client pb.AuthServiceClient) *AuthHandler {
	return &AuthHandler{cf, client}
}
func (hdl AuthHandler) GetTokenVerifyAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rq model.GetTokenVerifyAccountRequest
		if err := ctx.ShouldBind(&rq); err != nil {
			panic(err)
		}
		res, err := hdl.client.GetTokenVeriryAccount(ctx, &pb.VerifyAccountRequest{
			Email: rq.Email,
		})
		if err != nil {
			panic(err)
		}
		if res.Error != nil {
			panic(res.Error)
		}
		ctx.JSON(200, gin.H{"Message": "Token has been sent to your email!"})
	}

}
func (hdl AuthHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		b := model.LoginRequest{}

		if err := ctx.BindJSON(&b); err != nil {
			panic(err)
		}
		res, err := hdl.client.Login(ctx, &pb.LoginRequest{
			Email:    b.Email,
			Password: b.Password,
		})
		if err != nil {
			panic(err)
		}
		if res.Error != nil {
			panic(res.Error)
		}
		fmt.Println(res)
		ctx.SetCookie("refresh_token", res.RefreshToken, 3600*720, "/", "localhost", false, true)
		ctx.JSON(200, &res)

	}
}
func (hdl AuthHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := model.RegisterRequest{}
		if err := ctx.BindJSON(&body); err != nil {
			panic(err)
		}
		res, err := hdl.client.Register(ctx, &pb.RegisterRequest{
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

		ctx.JSON(200, gin.H{"Message": "Create User-Service successfully"})
	}

}
func (hdl AuthHandler) NewToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken, err := ctx.Cookie("refresh_token")
		if err != nil {
			panic(&errpb.ErrorResponse{
				Code:    401,
				Message: "Refresh token not found! Please login again",
			})
		}
		res, err := hdl.client.NewToken(ctx, &pb.NewTokenRequest{RefreshToken: refreshToken})
		if err != nil {

			panic(err)
		}
		if res.Error != nil {
			panic(res.Error)
		}
		ctx.JSON(200, &res)
	}

}
func (hdl AuthHandler) VerifyAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.MustGet("emailUser").(string)

		res, err := hdl.client.VerifyAccount(ctx, &pb.VerifyAccountRequest{
			Email: email,
		})

		if err != nil {
			panic(err)
		}

		if res.Error != nil {
			panic(res.Error)
		}
		ctx.JSON(200, gin.H{"Message": "Account has been verified"})
	}

}
