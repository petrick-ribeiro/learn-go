package main

import (
	"fmt"
	"login-auth/config"
	"login-auth/database"
	"login-auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
  err := config.Load()
  if err != nil {
    panic(err)
  }

  database.Connect()

  app := fiber.New()
  app.Use(cors.New(cors.Config{
    AllowCredentials: true,
  }))
  routes.Setup(app)

  app.Listen(fmt.Sprintf(":%s", config.GetServerPort()))
}
