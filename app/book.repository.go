package app

import (
	m "demo-1/models"

	"gorm.io/gorm"
)

type BookRepository interface {
  FindAll() ([]m.Book,error)
  FindById(id uint) (m.Book,error)
  Create(book m.Book) (m.Book,error)
  Update(book m.Book) (m.Book,error)
  Delete(id uint) error
}

type bookRepository struct { db *gorm.DB }

func NewBookRepository(db *gorm.DB) BookRepository {
  return &bookRepository{db}
}

func (r *bookRepository) FindAll() ([]m.Book, error) {
  var books []m.Book
  if err := r.db.Find(&books).Error; err!= nil {
    return nil, err
  }
  return books, nil
}

func (r *bookRepository) FindById(id uint) (m.Book, error) {
  var book m.Book
  if err := r.db.First(&book,id).Error; err != nil {
    return book, err
  }
  return book, nil
}

func (r *bookRepository) Create(book m.Book) (m.Book, error) {
  if err := r.db.Create(&book).Error; err != nil {
    return book, err
  }
  return book, nil
}

func (r *bookRepository) Update(book m.Book) (m.Book, error) {
  if err:= r.db.Save(&book).Error; err != nil {
    return book, err
  }
  return book, nil
}

func (r *bookRepository) Delete(id uint) error{
  if err:= r.db.Delete(&m.Book{}, id).Error; err != nil {
    return err
  }
  return nil
}

