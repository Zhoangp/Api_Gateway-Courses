package middleware

import "github.com/Zhoangp/Api_Gateway-Courses/config"

type middleareManager struct {
	cfg *config.Config
}

func NewMiddlewareManager(cfg *config.Config) *middleareManager {
	return &middleareManager{cfg}
}
