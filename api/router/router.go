package router

import (
	"github.com/nanami-wq/get-service-start/api/controller"
	"github.com/nanami-wq/get-service-start/api/middleware"

	"github.com/gin-gonic/gin"
)

func GenerateRouter(r *gin.Engine) {
	r.Use(gin.Recovery(), middleware.CorsWare(), middleware.ResponseMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	auth := controller.NewAuthController()
	auth.RegisterPublic(r.Group("/api"))

	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	auth.RegisterProtected(protected)
}
