package payment

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/payment/pb"
	"google.golang.org/grpc"
)

func NewPaymentService(cf *config.Config) pb.PaymentServiceClient {
	cc, err := grpc.Dial(cf.Services.PaymentUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewPaymentServiceClient(cc)
}
