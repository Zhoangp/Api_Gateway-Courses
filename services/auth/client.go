package auth

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/auth/pb"
	"google.golang.org/grpc"
)

// InitServiceClient func used to define auth service client
func NewAuthServiceClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Services.AuthUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
