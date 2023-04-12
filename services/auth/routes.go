package auth

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/auth/routes"
	"github.com/Zhoangp/Api_Gateway-Courses/services/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/login", svc.Login)
	routes.POST("/register", svc.Register)
	routes.GET("/token", svc.NewToken)
	routes.POST("/account", svc.GetTokenVerifyAccount)
	mdware := middleware.NewMiddlewareManager(c)

	routes.Use(mdware.RequireVerifyToken())
	routes.PATCH("/account/:token", svc.VerifyAccount)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)

}
func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
func (svc *ServiceClient) NewToken(ctx *gin.Context) {
	routes.NewToken(ctx, svc.Client)
}
func (sv *ServiceClient) VerifyAccount(ctx *gin.Context) {
	routes.VerifyAccount(ctx, sv.Client)
}
func (svc *ServiceClient) GetTokenVerifyAccount(ctx *gin.Context) {
	routes.GetTokenVerifyAccount(ctx,svc.Client)
}
