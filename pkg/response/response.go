package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

func WriteResponse(c *gin.Context, code int, error interface{}, data interface{}) {
	if error != nil {
		c.JSON(code, ErrorResponse{
			Code:    code,
			Message: error.(string),
			Details: data,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
