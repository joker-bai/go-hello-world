package v1

import (
	"github.com/gin-gonic/gin"
	"go-hello-world/app/http/controllers"
	"go-hello-world/pkg/response"
	"net/http"
)

type RootController struct {
	controllers.BaseController
}

func (r *RootController) Root(c *gin.Context) {
	response.WriteResponse(c, http.StatusOK, nil, gin.H{
		"result": "主页面",
	})
}
