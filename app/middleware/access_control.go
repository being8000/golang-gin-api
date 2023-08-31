package middleware

import (
	"errors"
	"fmt"
	"zehan/gin/app/pkg"
	"zehan/gin/utils"

	"github.com/allegro/bigcache"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// Authenticate determines if current subject has logged in.
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get session id
		sessionId, _ := c.Cookie("current_subject")
		// Get current subject
		sub, err := utils.GlobalCache.Get(sessionId)
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			c.AbortWithStatusJSON(401, utils.RestResponse{Message: "user hasn't logged in yet"})
			return
		}
		c.Set("current_subject", string(sub))
		c.Next()
	}
}

// Authorize determines if current subject has been authorized to take an action on an object.
func Authorize(obj string, act string, casbin *pkg.Casbin) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, existed := c.Get("current_subject")
		if !existed {
			c.AbortWithStatusJSON(401, utils.RestResponse{Message: "user hasn't logged in yet"})
			return
		}
		// casbin enforces policy
		ok, err := enforce(val.(string), obj, act, casbin)
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

func enforce(sub string, obj string, act string, cs *pkg.Casbin) (bool, error) {
	enforcer, err := casbin.NewEnforcer("./app/app/rbac_model.conf", cs.Adapter)
	if err != nil {
		return false, fmt.Errorf("failed to create casbin enforcer: %w", err)
	}
	// Load policies from DB dynamically
	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok, err := enforcer.Enforce(sub, obj, act)
	return ok, err
}
