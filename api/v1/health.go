package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": CodeSuccess,
		"msg":    "",
		"data": map[string]string{
			"status": "ok",
		},
	})
}
