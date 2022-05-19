package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-hello-world/app/http/controllers/v1"
)

func SetupRouter(router *gin.Engine) {
	ruc := new(v1.RootController)
	router.GET("/", ruc.Root)

	huc := new(v1.HealthController)
	router.GET("/health", huc.HealthCheck)
}
