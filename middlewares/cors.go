package middlewares

import (
  "time"
  "github.com/gin-contrib/cors"
)

var DefaultCors = cors.Default()
var CustomCors = cors.New(cors.Config{
  AllowOrigins:     []string{"https://foo.com"},
  AllowMethods:     []string{"PUT", "PATCH"},
  AllowHeaders:     []string{"Origin"},
  ExposeHeaders:    []string{"Content-Length"},
  AllowCredentials: true,
  AllowOriginFunc: func(origin string) bool {
    return origin == "https://github.com"
  },
  MaxAge: 12 * time.Hour,
})
