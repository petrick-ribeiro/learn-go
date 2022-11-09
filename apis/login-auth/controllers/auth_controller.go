package controllers

import (
	"login-auth/database"
	"login-auth/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
  var data map[string]string

  err := c.BodyParser(&data)
  if err != nil {
    return err
  }

  password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
  user := models.User{
    Name:     data["name"],
    Email:    data["email"],
    Password: password,
  }

  database.DB.Create(&user)

  return c.JSON(user)
} 

func Login(c *fiber.Ctx) error {
  var data map[string]string

  err := c.BodyParser(&data)
  if err != nil {
    return err
  }

  var user models.User

  database.DB.Where("email = ?", data["email"]).First(&user)

  if user.ID == 0 {
    c.Status(fiber.StatusNotFound)
    return c.JSON(fiber.Map{
      "error": true,
      "message": "user not found.",
    })
  }

  if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
    c.Status(fiber.StatusBadRequest)
    return c.JSON(fiber.Map{
      "error":   true,
      "message": "wrong password",
    })
  }

  claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
    Issuer: strconv.Itoa(int(user.ID)),
    ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
  })

  token, err := claims.SignedString([]byte(SecretKey))
  if err != nil {
    c.Status(fiber.StatusInternalServerError)
    return c.JSON(fiber.Map{
      "error":   true,
      "message": "login error",
    })
  }

  cookie := fiber.Cookie{
    Name:    "jwt",
    Value:   token,
    Expires: time.Now().Add(time.Hour * 24),
  }

  c.Cookie(&cookie)

  return c.JSON(fiber.Map{
    "error":   false,
    "message": "success",
  })
}

func User(c *fiber.Ctx) error {
  cookie := c.Cookies("jwt")

  token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
    return []byte(SecretKey), nil
  })
  if err != nil {
    c.Status(fiber.StatusUnauthorized)
    return c.JSON(fiber.Map{
      "error":   true,
      "message": "unauthenticated",
    })
  }
  
  claims := token.Claims.(*jwt.StandardClaims)

  var user models.User

  database.DB.Where("id = ?", claims.Issuer).First(&user)

  return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
  cookie := fiber.Cookie{
    Name: "jwt",
    Value: "-",
    Expires: time.Now().Add(-time.Hour),
    HTTPOnly: true,
  }

  c.Cookie(&cookie)

  return c.JSON(fiber.Map{
    "error":   false,
    "message": "success",
  })
}
