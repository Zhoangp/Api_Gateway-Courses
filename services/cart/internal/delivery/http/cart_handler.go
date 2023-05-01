package http

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/cart/pb"
)

type cartHandler struct {
	cf     *config.Config
	client pb.CartServiceClient
}

func NewCartHandler(cf *config.Config, client pb.CartServiceClient) *cartHandler {
	return &cartHandler{cf, client}
}
