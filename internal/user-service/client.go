package user_service

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/user-service/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.UserServiceClient
}

func NewServiceClient(cf *config.Config) pb.UserServiceClient {
	client, err := grpc.Dial(cf.Services.UserUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)

	}
	return pb.NewUserServiceClient(client)
}
