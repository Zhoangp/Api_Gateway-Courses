package routes

import (
	"github.com/Zhoangp/Api_Gateway-Courses/internal/auth-service/pb"
	"github.com/gin-gonic/gin"
)
func VerifyAccount(ctx *gin.Context, client pb.AuthServiceClient) {
	email := ctx.MustGet("emailUser").(string)

	res, err := client.VerifyAccount(ctx, &pb.VerifyAccountRequest{
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
