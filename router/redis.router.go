package router

import (
	"encoding/json"
	"net/http"
	"zehan/gin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RedisSet struct {
	Key   string      `json:"key" binding:"required"`
	Value interface{} `json:"value" binding:"required"`
}

func (router *Router) Router_Redis(k *utils.Kit) {
	self := k
	app := self.App.Group("redis")
	{
		app.POST("/set", func(c *gin.Context) {
			var set RedisSet
			if err := c.ShouldBindJSON(&set); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if _, err := self.Redis.Set(c, set.Key, set.Value, 0).Result(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}
			if _, err := self.Redis.HMSet(c, "list", "key1", "value1", "key2", "value2").Result(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		})
		app.GET("/get/:type/:key", func(ctx *gin.Context) {
			var red struct {
				Type string `uri:"type"`
				Key  string `uri:"key"`
			}

			if err := ctx.ShouldBindUri(&red); err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{"error": err})
				return
			}
			bt, _ := json.Marshal(red)
			zap.L().Info("Redis:", zap.Binary("b", bt))
			switch red.Type {
			case "list":
				res, err := self.Redis.HMGet(ctx, red.Key).Result()
				if err != nil {
					ctx.JSON(http.StatusBadRequest, err)
				}
				ctx.JSON(http.StatusOK, res)
			case "string":
				res, err := self.Redis.Get(ctx, red.Key).Result()
				if err != nil {
					ctx.JSON(http.StatusBadRequest, err)
				}
				ctx.JSON(http.StatusOK, res)
			}
			return
		})
	}

}
