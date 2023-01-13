package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context) {
	c.String(http.StatusBadRequest, "Bad Request")
}
