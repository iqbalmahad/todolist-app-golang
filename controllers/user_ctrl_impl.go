package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbalmahad/todolist-app-golang.git/models"
	"github.com/iqbalmahad/todolist-app-golang.git/repositories"
)

type userController struct {
	userRepo repositories.UserRepo
}

func NewUserController(userRepo repositories.UserRepo) UserController {
	return &userController{
		userRepo: userRepo,
	}
}

func (uc *userController) Create(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	if err := uc.userRepo.Create(user); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (uc *userController) Read(c *fiber.Ctx) error {
	userID := c.Params("id")
	user, err := uc.userRepo.Read(userID)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (uc *userController) Update(c *fiber.Ctx) error {
	userID := c.Params("id")
	var updateUser models.User
	if err := c.BodyParser(&updateUser); err != nil {
		return err
	}
	updateUser.ID = userID
	if err := uc.userRepo.Update(updateUser); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(updateUser)
}

func (uc *userController) Delete(c *fiber.Ctx) error {
	userID := c.Params("id")
	if err := uc.userRepo.Delete(userID); err != nil {
		return err
	}
	return c.Status(fiber.StatusNoContent).SendString("")
}

func (uc *userController) Index(c *fiber.Ctx) error {
	users, err := uc.userRepo.Index()
	if err != nil {
		return err
	}
	return c.JSON(users)
}
