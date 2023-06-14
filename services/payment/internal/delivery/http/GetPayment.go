package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/services/payment/pb"
	"github.com/gin-gonic/gin"
)

func (hdl paymentHandler) GetPayment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.MustGet("userId").(string)
		res, err := hdl.client.GetPayment(ctx, &pb.GetPaymentRequest{
			UserId: userId,
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
		ctx.JSON(200, res)
	}
}
