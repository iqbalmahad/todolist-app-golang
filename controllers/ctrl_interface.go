package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Create(c *fiber.Ctx) error
	Read(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Index(c *fiber.Ctx) error
}

type TodoController interface {
	Create(c *fiber.Ctx) error
	Read(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Index(c *fiber.Ctx) error
}

type AuthController interface {
	Login(c *fiber.Ctx) error
	Accessible(c *fiber.Ctx) error
	Restricted(c *fiber.Ctx) error
}
