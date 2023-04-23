package middleware

import "github.com/Zhoangp/Api_Gateway-Courses/config"

type MiddleareManager struct {
	cfg *config.Config
}

func NewMiddlewareManager(cfg *config.Config) *MiddleareManager {
	return &MiddleareManager{cfg}
}
