package middleware

import (
	"github.com/Zhoangp/Api_Gateway-Courses/internal/error/pb"
	"github.com/Zhoangp/Api_Gateway-Courses/pkg/common"
	"github.com/gin-gonic/gin"
)

func (m *middleareManager) Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if appErr, ok := err.(*pb.ErrorResponse); ok {
					ctx.AbortWithStatusJSON(int(appErr.Code), appErr)
					panic(err)
				}
				if appErr, ok := err.(*common.AppError); ok {
					ctx.AbortWithStatusJSON(int(appErr.StatusCode), appErr)
					panic(err)
				}
				ctx.AbortWithStatusJSON(500, &pb.ErrorResponse{
					Code:    500,
					Message: "Internal Server error",
				})
			}
		}()
		ctx.Next()
	}
}
