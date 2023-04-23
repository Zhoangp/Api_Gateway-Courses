package course

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
	"google.golang.org/grpc"
)

func NewCourseServiceClient(cf *config.Config) pb.CourseServiceClient {
	cc, err := grpc.Dial(cf.Services.CourseUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewCourseServiceClient(cc)
}
