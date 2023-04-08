package user

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/middleware"
	"github.com/Zhoangp/Api_Gateway-Courses/internal/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, c *config.Config) {
	svc := ServiceClient{
		Client: NewServiceClient(c),
	}
	mdware := middleware.NewMiddlewareManager(c)
	routes := r.Group("/user")

	routes.POST("/otp", svc.GetTokenResetPass)
	routes.PUT("/password/:token", mdware.RequireVerifyToken(), svc.ResetPassword)

	routes.Use(mdware.RequiredAuth())
	routes.PUT("/", svc.UpdateUser)

}
func (svc *ServiceClient) UpdateUser(ctx *gin.Context) {
	routes.Update(ctx, svc.Client)
}
func (svc *ServiceClient) GetTokenResetPass(ctx *gin.Context) {
	routes.GetTokenResetPass(ctx, svc.Client)
}
func (svc *ServiceClient) ResetPassword(ctx *gin.Context) {
	routes.ResetPassword(ctx, svc.Client)
}
