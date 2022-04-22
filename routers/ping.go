package routers

import (
  "github.com/gin-gonic/gin"
)

func SetPingRoutes(r *gin.RouterGroup) {
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "status": "up",
    })
  })
}
