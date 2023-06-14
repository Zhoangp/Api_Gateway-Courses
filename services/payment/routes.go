package payment

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/middleware"
	"github.com/Zhoangp/Api_Gateway-Courses/services/payment/internal/delivery/http"
	"github.com/gin-gonic/gin"
)

func NewPaymentRoutes(r *gin.Engine, cf *config.Config, mdw *middleware.MiddleareManager) {
	router := r.Group("/payment")
	client := NewPaymentService(cf)
	paymentHdl := http.NewPaymentHandler(cf, client)
	router.Use(mdw.RequiredAuth())
	router.POST("", paymentHdl.CheckOut())
	router.POST("/capture", paymentHdl.Capture())
	router.POST("/account", paymentHdl.ConnectToPaypalAccount())
	router.GET("/", paymentHdl.GetPayment())
	router.GET("/paypal", paymentHdl.GetPaypal())

}
