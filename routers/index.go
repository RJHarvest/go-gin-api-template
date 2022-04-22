package routers

import (
  "github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.RouterGroup) {
  SetPingRoutes(g)
}
