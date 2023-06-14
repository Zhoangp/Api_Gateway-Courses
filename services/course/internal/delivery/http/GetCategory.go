package http

import (
	"github.com/Zhoangp/Api_Gateway-Courses/services/course/pb"
	"github.com/gin-gonic/gin"
)

func (hdl CourseHandler) GetCategories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := hdl.client.GetAllCategories(ctx, &pb.GetAllCategoriesRequest{})
		if err != nil {
			panic(err)
		}
		if res.Error != nil {
			panic(res.Error)
		}
		ctx.JSON(200, res)
	}
}
