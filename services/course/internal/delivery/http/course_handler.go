package http

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
)

type CourseHandler struct {
	cf     *config.Config
	client pb.CourseServiceClient
}

func NewCourseHandler(cf *config.Config, client pb.CourseServiceClient) *CourseHandler {
	return &CourseHandler{cf, client}
}
