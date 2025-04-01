package models

import "time"

type Book struct {
  ID            uint    `json:"id"`
  Title         string  `json:"title"`
  Description   string  `json:"description"`
  AuthorID      uint    `json:"author_id"`
  PublishedAt   time.Time   `json:"published_at"`
  CreatedAt     time.Time   `json:"created_at"`
  UpdatedAt     time.Time   `json:"updated_at"`

  Author  Author `gorm:"foreignKey:AuthorID" json:"author"`
}

