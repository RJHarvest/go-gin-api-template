package v1

import (
  "github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.RouterGroup) {
  SetHelloRoutes(g)
}
