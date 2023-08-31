package middleware

import (
	"log"
	"zehan/gin/app/pkg"
	"zehan/gin/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// Authorize determines if current subject has been authorized to take an action on an object.
func ApiCheck(policy string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// val, existed := c.Get("current_subject")
		// zap.S().Info(val)
		// if !existed {
		// 	c.AbortWithStatusJSON(401, utils.RestResponse{Message: "user hasn't logged in yet"})
		// 	return
		// }
		enforcer := pkg.GetEnforcer(policy)
		// casbin enforces policy
		ok, str, err := enforcer.EnforceEx("test", c.Request.URL, c.Request.Method)
		log.Println(str)
		if err != nil {
			zap.L().Error("Error: ", zap.Error(err))
			c.AbortWithStatusJSON(500, utils.RestResponse{Message: "error occurred when authorizing user"})
			return
		}
		if !ok {
			c.AbortWithStatusJSON(403, utils.RestResponse{Message: "forbidden"})
			return
		}
		c.Next()
	}
}
