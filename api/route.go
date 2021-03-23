package api

import (
	"github.com/fs714/go-amis-admin/api/v1"
	"github.com/fs714/go-amis-admin/utils/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.Conf.DefaultConf.RunMode)
	gin.DisableConsoleColor()
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/api/v1/health"},
	}))
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	apiv1 := r.Group("/api/v1")

	{
		apiv1.GET("/health", v1.Health)
	}

	return r
}
