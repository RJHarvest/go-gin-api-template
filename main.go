package main

import (
  "fmt"
  "log"
  "github.com/gin-gonic/gin"
  "github.com/RJHarvest/gin-api-template/services/postgresql"
  "github.com/RJHarvest/gin-api-template/routers/v1"
  "github.com/RJHarvest/gin-api-template/routers"
  "github.com/RJHarvest/gin-api-template/middlewares"
)

func main() {
  postgresql.InitDB()
  router := gin.Default()
  router.Use(middlewares.CustomCors)
  api := router.Group("/api")
  {
    routers.InitRoutes(api)

    // v1 api routes
    apiV1 := api.Group("v1")
    v1.InitRoutes(apiV1)
  }

  err := router.Run(":5000")
  if err != nil {
    panic(fmt.Sprintf("ERROR: Failed to start server %v", err))
  }
  log.Printf("Server in running on port: %d", 5000)
}
