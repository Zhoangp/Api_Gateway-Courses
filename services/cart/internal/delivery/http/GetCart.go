package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/services/cart/pb"
	"github.com/gin-gonic/gin"
)

func (hdl cartHandler) GetCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.MustGet("userId").(string)
		res, err := hdl.client.GetCart(ctx, &pb.GetCartRequest{
			Id: userId,
		})
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if res.Error != nil {
			fmt.Println(res.Error)
			panic(res.Error)
		}
		ctx.JSON(200, res)
	}
}
