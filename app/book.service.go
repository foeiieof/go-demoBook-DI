package app

import (
  m "demo-1/models"  
)

type BookService interface {
  GetAllBooks() ([]m.Book,error)
  GetBookById(id uint) (m.Book, error)
  CreateBook(book m.Book) (m.Book,error)
  UpdateBook(book m.Book) (m.Book,error)
  DeleteBook(id uint) error
}

type bookService struct { 
  repo BookRepository
}

func NewBookService(repo BookRepository) BookService { 
  return &bookService{repo}
}

func (s *bookService) GetAllBooks() ([]m.Book,error){
  return s.repo.FindAll()
}

func (s *bookService) GetBookById(id uint) (m.Book, error) {
  return s.repo.FindById(id)
} 

func (s *bookService) CreateBook(book m.Book) (m.Book, error) {
 return s.repo.Create(book)
}

func (s *bookService) UpdateBook(book m.Book) (m.Book, error) {
  return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id uint) error {
  return s.repo.Delete(id)
}
