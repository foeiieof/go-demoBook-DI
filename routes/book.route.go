package routes

import (
	book "demo-1/app/book"

  "demo-1/config"
	"github.com/gin-gonic/gin"
)

 
func BookRoutes(r *gin.RouterGroup){
    repository := book.NewBookRepository(config.DB)
    service := book.NewBookService(repository)
    bookHandlers := book.NewBookHandler(service)

    r.GET("/", bookHandlers.GetAllBooks)
    r.GET("/:id", bookHandlers.GetBookById)
    r.POST("/", bookHandlers.CreateBook)
    r.PUT("/:id",bookHandlers.UpdateBook)
    r.DELETE("/:id",bookHandlers.DeleteBook)
}

