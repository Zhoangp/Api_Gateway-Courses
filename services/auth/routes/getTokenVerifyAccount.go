package routes

import (
	"github.com/Zhoangp/Api_Gateway-Courses/services/auth/pb"
	"github.com/gin-gonic/gin"
)
type request struct {
	Email string `json:"email"`
}
func GetTokenVerifyAccount(ctx *gin.Context, client pb.AuthServiceClient) {
	var rq request
	if err := ctx.ShouldBind(&rq); err != nil {
		panic(err)
	}
	res, err := client.GetTokenVeriryAccount(ctx, &pb.VerifyAccountRequest{
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
