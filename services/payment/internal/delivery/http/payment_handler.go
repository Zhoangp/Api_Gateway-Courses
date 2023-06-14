package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/payment/pb"
	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	cf     *config.Config
	client pb.PaymentServiceClient
}

func NewPaymentHandler(cf *config.Config, client pb.PaymentServiceClient) *paymentHandler {
	return &paymentHandler{cf, client}
}
func (hdl *paymentHandler) ConnectToPaypalAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.MustGet("userId").(string)
		var t pb.Token
		if err := ctx.ShouldBind(&t); err != nil {
			fmt.Println(err)
			panic(err)
		}
		res, err := hdl.client.ConnectPaypalAccount(ctx, &pb.ConnectPaypalRequest{
			IdentifyToken: &pb.Token{
				Token: t.Token,
			},
			UserId: userId,
		})
		if err != nil {
			fmt.Println(err)

			panic(err)
		}
		ctx.JSON(200, res)

	}
}
