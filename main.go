package main

import (
	"demo-1/app/book"
	"demo-1/config"
	"demo-1/routes"

	"github.com/gin-gonic/gin"
)

func main() {

  config.Initail()

  repository := app.NewBookRepository(config.DB)
  service := app.NewBookService(repository)
  handler := app.NewBookHandler(service)

  r := gin.Default()

  routes.SetupRoutes(r, handler)

  r.Run(":8080")
}
