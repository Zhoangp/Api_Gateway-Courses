package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/services/payment/pb"
	"github.com/gin-gonic/gin"
)

func (hdl paymentHandler) CheckOut() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.MustGet("userId").(string)
		var cartId struct {
			CartId string `json:"cartId"`
		}
		if err := ctx.ShouldBind(&cartId); err != nil {
			panic(err)
		}
		fmt.Println(cartId)
		res, err := hdl.client.CheckOutWithPaypal(ctx, &pb.CheckoutRequest{
			UserId: userId,
			CartId: cartId.CartId,
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
