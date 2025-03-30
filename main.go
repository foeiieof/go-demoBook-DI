package main

import (
	"demo-1/app"
	"demo-1/config"
	"demo-1/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  config.Initail()

  repository := app.NewBookRepository(config.DB)
  service := app.NewBookService(repository)
  handler := app.NewBookHandler(service)

  r := gin.Default()

  routes.SetupRoutes(r, handler)

  r.GET("/", func(c *gin.Context){ c.JSON(http.StatusOK, gin.H{"message": "OK"}) })

  r.Run(":8080")
}

