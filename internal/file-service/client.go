package file_service

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/file-service/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct  {
	Client pb.FileServiceClient
}
func InitServiceClient(c *config.Config) pb.FileServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Services.FileUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewFileServiceClient(cc)
}