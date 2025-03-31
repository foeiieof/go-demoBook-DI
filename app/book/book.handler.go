package app

import (
	m "demo-1/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
  service BookService
}

func NewBookHandler(service BookService) *BookHandler {
  return &BookHandler{service}
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
  books, err := h.service.GetAllBooks()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message":""})
    return
  }
  c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookById(c *gin.Context) {
  id,_ := strconv.Atoi(c.Param("id"))
  book, err := h.service.GetBookById(uint(id))
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"message":"Book not found"})
    return
  }
  c.JSON(http.StatusOK,book)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
  var book m.Book
  if err := c.ShouldBindJSON(&book); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid data"})
    return
  }
  createdBook, err := h.service.CreateBook(book)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not create book"})
    return
  }

  c.JSON(http.StatusCreated, createdBook)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
  id,_ := strconv.Atoi(c.Param("id"))
  var book m.Book
  if err := c.ShouldBindJSON(&book); err != nil {
    c.JSON(http.StatusBadGateway, gin.H{"message":"Invalid data"})
    return
    // what about this return ?? if we could not put on it ???
  }
  book.ID = uint(id)
  updatedBook, err := h.service.UpdateBook(book)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not updat book"})
    return
  }
  c.JSON(http.StatusOK,updatedBook)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
  id,_ := strconv.Atoi(c.Param("id"))
  if err := h.service.DeleteBook(uint(id)); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not delete book"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message":"Book deleted successfully"})
}

