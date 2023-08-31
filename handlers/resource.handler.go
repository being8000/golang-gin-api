package handlers

import (
	"zehan/gin/utils"

	"github.com/gin-gonic/gin"
)

type ResouceHandler struct {
	Kit *utils.Kit
}

func (handler *ResouceHandler) ReadResource(c *gin.Context) {
	// some stuff
	// blahblah...

	c.JSON(200, utils.RestResponse{Code: 1, Message: "read resource successfully", Data: "resource"})
}

func (handler *ResouceHandler) WriteResource(c *gin.Context) {
	// some stuff
	// blahblah...

	c.JSON(200, utils.RestResponse{Code: 1, Message: "write resource successfully", Data: "resource"})
}
