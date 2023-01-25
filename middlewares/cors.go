package middlewares

import (
  "time"
  "github.com/gin-contrib/cors"
  "github.com/RJHarvest/gin-api-template/config"
)

var DefaultCors = cors.Default()
var CustomCors = cors.New(cors.Config{
  AllowOrigins:     config.GetConfig().CorsAllowOrigins,
  AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
  AllowHeaders:     []string{"Origin"},
  ExposeHeaders:    []string{"Content-Length"},
  AllowCredentials: true,
  AllowOriginFunc: func(origin string) bool {
    return origin == "https://github.com"
  },
  MaxAge: 12 * time.Hour,
})
