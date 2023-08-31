package router

import (
	"zehan/gin/app/database"
	"zehan/gin/app/pkg"
	"zehan/gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

type Router struct {
}

func (router *Router) mount(k *utils.Kit) {
	router.Router_Redis(k)
	router.Router_User(k)
	router.Router_Example(k)
	router.Router_BasicAuth(k)
	router.Router_Resource(k)
	router.Router_Permission(k)
}

func New(app *gin.Engine) {
	// connect database
	mysql := &database.Mysql{}
	mysql.Connect()
	// connect Redis
	redis := &pkg.Redis{}
	router := &Router{}
	router.mount(&utils.Kit{
		App: app,
		// Valid: validator.New(),
		DB:     mysql.DB,
		PG:     paginate.New(),
		Redis:  redis.NewRedisClient(),
		Casbin: pkg.NewCasbin(mysql.DB),
	})
}
