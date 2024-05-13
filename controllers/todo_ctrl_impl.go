package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbalmahad/todolist-app-golang.git/models"
	"github.com/iqbalmahad/todolist-app-golang.git/repositories"
)

type todoController struct {
	todoRepo repositories.TodoRepo
}

func NewTodoController(todoRepo repositories.TodoRepo) TodoController {
	return &todoController{
		todoRepo: todoRepo,
	}
}

func (tc *todoController) Create(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return err
	}
	if err := tc.todoRepo.Create(todo); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(todo)
}

func (tc *todoController) Read(c *fiber.Ctx) error {
	todoID := c.Params("id")
	todo, err := tc.todoRepo.Read(todoID)
	if err != nil {
		return err
	}
	return c.JSON(todo)
}

func (tc *todoController) Update(c *fiber.Ctx) error {
	todoID := c.Params("id")
	id, err := strconv.Atoi(todoID)
	if err != nil {
		return err
	}
	var updateTodo models.Todo
	if err := c.BodyParser(&updateTodo); err != nil {
		return err
	}
	updateTodo.ID = uint(id)
	if err := tc.todoRepo.Update(updateTodo); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(updateTodo)
}

func (tc *todoController) Delete(c *fiber.Ctx) error {
	todoID := c.Params("id")
	if err := tc.todoRepo.Delete(todoID); err != nil {
		return err
	}
	return c.Status(fiber.StatusNoContent).SendString("")
}

func (tc *todoController) Index(c *fiber.Ctx) error {
	todos, err := tc.todoRepo.Index()
	if err != nil {
		return err
	}
	return c.JSON(todos)
}
