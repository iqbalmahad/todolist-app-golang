package controllers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/iqbalmahad/todolist-app-golang.git/repositories"
	"gorm.io/gorm"
)

type authController struct {
	authRepo repositories.AuthRepo
}

func NewAutController(authRepo repositories.AuthRepo) AuthController {
	return &authController{
		authRepo: authRepo,
	}
}

func (ar *authController) Login(c *fiber.Ctx) error {
	usernameOrEmail := c.FormValue("username_or_email")
	password := c.FormValue("password")

	// Pengecekan username atau email dan password
	// Dapatkan pengguna berdasarkan username atau email
	user, err := ar.authRepo.GetByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	validCredentials := (usernameOrEmail == user.Username || usernameOrEmail == user.Email) && password == user.Password

	// Throws Unauthorized error jika kredensial tidak valid
	if !validCredentials {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"ID":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"name":     user.Name,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func (ar *authController) Accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func (ar *authController) Restricted(c *fiber.Ctx) error {
	username := c.Locals("user").(*jwt.Token)
	claims := username.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
