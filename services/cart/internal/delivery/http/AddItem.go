package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/services/cart/internal/model"
	"github.com/Zhoangp/Api_Gateway-Courses/services/cart/pb"
	"github.com/gin-gonic/gin"
	"reflect"
)

func (hdl cartHandler) AddItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data model.AddItem
		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}
		res, err := hdl.client.AddItem(ctx, &pb.AddItemRequest{
			CartId:   data.CartId,
			CourseId: data.CourseId,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(reflect.TypeOf(res.Error))

		if res.Error != nil {
			panic(res.Error)
		}
		ctx.JSON(200, gin.H{"status": 200, "message": "Add course to cart successfully!"})
	}
}
