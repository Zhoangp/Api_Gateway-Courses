package cart

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/cart/pb"
	"google.golang.org/grpc"
)

func NewCartService(cf *config.Config) pb.CartServiceClient {
	cc, err := grpc.Dial(cf.Services.CartUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewCartServiceClient(cc)
}
