package routes

import (
	"demo-1/app/book"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, bookHandlers *app.BookHandler){

  bookRoutes := r.Group("/books")
  {
    bookRoutes.GET("/", bookHandlers.GetAllBooks)
    bookRoutes.GET("/:id", bookHandlers.GetBookById)
    bookRoutes.POST("/", bookHandlers.CreateBook)
    bookRoutes.PUT("/:id",bookHandlers.UpdateBook)
    bookRoutes.DELETE("/:id",bookHandlers.DeleteBook)
  }


}

