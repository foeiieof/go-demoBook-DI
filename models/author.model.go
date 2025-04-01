package models

import "time"

type Author struct {
  ID        uint   `gorm:"primaryKey" json:"id"`
  Name      string `json:"name"`
  Bio       string `json:"bio"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`

  Books []Book `gorm:"foreignKey:AuthorID" json:"books"`
}
