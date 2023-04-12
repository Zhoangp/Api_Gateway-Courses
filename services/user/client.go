package user

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.UserServiceClient
}
func NewServiceClient(cf *config.Config) pb.UserServiceClient {
	maxSize := 1024 * 1024 * 10
	fmt.Println(maxSize)
	client, err := grpc.Dial(cf.Services.UserUrl,grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxSize), grpc.MaxCallSendMsgSize(maxSize)))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}



	return pb.NewUserServiceClient(client)
}
