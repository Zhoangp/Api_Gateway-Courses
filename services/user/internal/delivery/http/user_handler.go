package http

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/pb"
)

type UserHandler struct {
	cf     *config.Config
	client pb.UserServiceClient
}

func NewUserHandler(cf *config.Config, client pb.UserServiceClient) *UserHandler {
	return &UserHandler{cf, client}
}
