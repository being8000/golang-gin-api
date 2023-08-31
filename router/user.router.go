package router

import (
	"zehan/gin/handlers"
	"zehan/gin/utils"
)

func (r *Router) Router_User(k *utils.Kit) {
	handler := &handlers.UserHandler{
		Kit: k,
	}

	app := k.App.Group("user")
	app.POST("/login", handler.Login)
	app.POST("/addPolicy", handler.AddPolicy)
}
