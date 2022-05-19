package v1

import (
	"github.com/gin-gonic/gin"
	"go-hello-world/app/http/controllers"
	"go-hello-world/pkg/response"
	"net/http"
)

type HealthController struct {
	controllers.BaseController
}

func (h *HealthController) HealthCheck(c *gin.Context) {
	response.WriteResponse(c, http.StatusOK, nil, gin.H{
		"result": "健康检测页面",
		"status": "OK",
	})
}
