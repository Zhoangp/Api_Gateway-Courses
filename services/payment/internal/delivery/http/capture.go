package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/services/payment/pb"
	"github.com/gin-gonic/gin"
)

func (hdl paymentHandler) Capture() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//userId := ctx.MustGet("userId").(string)
		orderId := ctx.Query("orderId")
		token := ctx.Query("token")

		res, err := hdl.client.CaptureWithPaypal(ctx, &pb.CaptureRequest{
			Token:   token,
			OrderId: orderId,
		})
		if err != nil {
			fmt.Print(err)
			panic(err)
		}
		if res.Error != nil {
			fmt.Print(res.Error)
			panic(res.Error)
		}
		fmt.Println(res)
		ctx.JSON(200, gin.H{"message": "Capture successful", "code": 200})
	}
}
