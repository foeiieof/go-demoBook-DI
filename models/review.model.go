package models

import "time"

type Review struct {
  ID        uint    `gorm:"primaryKey" json:"id"`
  BookID    uint    `json:"book_id"`
  Reviewer  *string `json:"reviewer"`
  Rating    *int    `json:"rating"`
  Comment   string  `json:"comment"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`

  Book Book `gorm:"foreignKey:BookID" json:"book"`
}
