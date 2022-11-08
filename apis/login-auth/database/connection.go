package database

import (
	"fmt"
	"login-auth/config"
  "login-auth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect()  {
  cfg := config.GetDB()

  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    cfg.Host, cfg.User, cfg.Pass, cfg.DB, cfg.Port,
  )

  conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  conn.AutoMigrate(&models.User{})

}
