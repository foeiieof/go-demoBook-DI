package main

import (
	"demo-1/config"
	"demo-1/routes"

	"github.com/gin-gonic/gin"
)

func main() {
  config.Initail()
  r := gin.Default()
  routes.SetupRoutes(r)
  r.Run(":8080")
}
