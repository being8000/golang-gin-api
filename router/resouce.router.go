package router

import (
	"zehan/gin/app/middleware"
	"zehan/gin/handlers"
	"zehan/gin/utils"
)

func (r *Router) Router_Resource(k *utils.Kit) {
	handler := &handlers.ResouceHandler{
		Kit: k,
	}

	resource := k.App.Group("api")
	resource.Use(middleware.Authenticate())
	{
		resource.GET("/resource", middleware.Authorize("resource", "read", k.Casbin), handler.ReadResource)
		resource.POST("/resource", middleware.Authorize("resource", "write", k.Casbin), handler.WriteResource)
	}
}
