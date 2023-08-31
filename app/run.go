package app

import (
	"zehan/gin/app/middleware"
	"zehan/gin/app/pkg"
	"zehan/gin/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunApplicationContext() {
	pkg.InitConfig()
	pkg.InitZapLogger()
	r := gin.New()
	r.SetTrustedProxies(nil)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig)) // CORS configuraion
	r.Use(middleware.RequestID())
	// logger := zap.L()
	// r.Use(gin.LoggerWithWriter(writer))
	r.Use(middleware.Logger())
	r.Use(gin.RecoveryWithWriter(pkg.GetLogWriter()))
	// r.Use(gin.Logger())
	// r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	// r.Use(ginzap.RecoveryWithZap(logger, true))
	// initialize routers
	router.New(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
