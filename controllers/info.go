package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 首页页面
// curl -i -X GET http://0.0.0.0:8080/v1/info
func GetInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
