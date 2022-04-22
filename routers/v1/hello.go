package v1

import (
  "github.com/gin-gonic/gin"
  "github.com/RJHarvest/gin-api-template/controllers/v1/hello"
)

func SetHelloRoutes(r *gin.RouterGroup) {
  r.GET("/hello", func(c *gin.Context) {
    message := hello.GetHello()
    c.JSON(200, gin.H{
      "message": message,
    })
  })
}
