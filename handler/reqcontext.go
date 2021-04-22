package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInt64IdFromReqContext(c *gin.Context) int64 {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	return id
}
