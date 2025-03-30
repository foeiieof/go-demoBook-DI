package config

import (
	"demo-1/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initail() {

  var err error

  envFile := os.Getenv("ENV_FILE")
  if envFile == ""{ envFile = "env/.local.env" }

  err = godotenv.Load(envFile)
  if err != nil { log.Panicf("Error loading .env: %s",err) }

  dataBase := os.Getenv("DATABASE_URL")

  DB, err = gorm.Open(mysql.Open(dataBase), &gorm.Config{})
  if err != nil {
    log.Fatal("Could not connet to database", err)
    return
  }

  modelList := []interface{} { &models.Book{}, }
  DB.AutoMigrate(modelList...)
}



