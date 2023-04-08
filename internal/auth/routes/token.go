package routes

import (
	pb "github.com/Zhoangp/Api_Gateway-Courses/internal/auth/pb"
	errpb "github.com/Zhoangp/Api_Gateway-Courses/internal/error/pb"

	"github.com/gin-gonic/gin"
)

func NewToken(ctx *gin.Context, client pb.AuthServiceClient) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		panic(&errpb.ErrorResponse{
			Code:    401,
			Message: "Refresh token not found! Please login again",
		})
	}
	res, err := client.NewToken(ctx, &pb.NewTokenRequest{RefreshToken: refreshToken})
	if err != nil {

		panic(err)
	}
	if res.Error != nil {
		panic(res.Error)
	}
	ctx.JSON(200, &res)
}
