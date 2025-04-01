package routes

import (
	book "demo-1/app/book"

  "demo-1/config"
	"github.com/gin-gonic/gin"
)

 
func BookRoutes(r *gin.RouterGroup){

    repository := book.NewBookRepository(config.DB)
    service := book.NewBookService(repository)
    handlers := book.NewBookHandler(service)

    r.GET("/", handlers.GetAllBooks)
    r.GET("/:id", handlers.GetBookById)
    r.POST("/", handlers.CreateBook)
    r.PUT("/:id",handlers.UpdateBook)
    r.DELETE("/:id", handlers.DeleteBook)
}

