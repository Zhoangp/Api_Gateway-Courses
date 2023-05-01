package cart

import (
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/cart/internal/delivery/http"
	"github.com/Zhoangp/Api_Gateway-Courses/services/middleware"
	"github.com/gin-gonic/gin"
)

func NewCartRoutes(r *gin.Engine, cf *config.Config, mdw *middleware.MiddleareManager) {
	router := r.Group("/cart")
	client := NewCartService(cf)
	cartHdl := http.NewCartHandler(cf, client)
	router.Use(mdw.RequiredAuth())
	router.GET("", cartHdl.GetCart())
	router.POST("", cartHdl.AddItem())
}
