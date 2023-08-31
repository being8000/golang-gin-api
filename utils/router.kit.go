package utils

import (
	"zehan/gin/app/pkg"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// type Validator struct {
// 	Validate *validator.Validate
// }

type Kit struct {
	App    *gin.Engine          // Fiber App Object
	DB     *gorm.DB             // `Mysql`
	PG     *paginate.Pagination // `Pagination Tool`
	Redis  *redis.Client        // `redis`
	Casbin *pkg.Casbin
}
