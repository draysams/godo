package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello, Worlds!")
	app := fiber.New()

	err := godotenv.Load(".env.development.local") // env file for local development
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		fmt.Println("Check read all!")
		return c.Status(fiber.StatusOK).JSON(todos)
	})

	app.Get("/api/todos/:id", func(c *fiber.Ctx) error {
		fmt.Println("Check read one!")
		id := c.Params("id")
		for _, todo := range todos {
			if fmt.Sprintf("%d", todo.ID) == id {
				return c.Status(fiber.StatusOK).JSON(todo)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		fmt.Println("Check create!")
		todo := &Todo{}
		if err := c.BodyParser(&todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if todo.Body == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body is required"})
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.Status(fiber.StatusCreated).JSON(todo)
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		fmt.Println("Check patch!")
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprintf("%d", todo.ID) == id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(fiber.StatusOK).JSON(todos[i])
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		fmt.Println("Check delete!")
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprintf("%d", todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(fiber.StatusNoContent).Send(nil)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	})

	log.Fatal(app.Listen(":" + PORT))
}
