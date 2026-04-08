package gin

import (
	"github.com/nanami-wq/get-service-start/api/middleware"
	"github.com/nanami-wq/get-service-start/api/router"
	"github.com/nanami-wq/get-service-start/config"

	"github.com/gin-gonic/gin"
)

func GinInit() *gin.Engine {
	r := gin.Default()
	config.MustLoad()
	router.GenerateRouter(r)
	middleware.InitSecret(config.GetConfig().JWT.Secret)
	return r
}
