package router

import (
	"zehan/gin/app/middleware"
	"zehan/gin/utils"

	"github.com/gin-gonic/gin"
)

func (r *Router) Router_Permission(k *utils.Kit) {

	resource := k.App.Group("pm")

	resource.Use(middleware.ApiCheck("permission.csv"))
	{
		resource.GET("test", func(c *gin.Context) {
			return
		})
	}
}
