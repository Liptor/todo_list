package main

import (
	"context"
	"os"

	"github.com/Liptor/todo_list.git/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	app := fiber.New()

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		os.Exit(1)
	}

	defer dbpool.Close()

	h := handlers.NewHandler(dbpool)

	app.Post("/tasks", h.CreateTaskHandler)
	app.Get("/tasks", h.GetTaskHandler)
	app.Put("/tasks/:id", h.UpdateTaskHandler)
	app.Delete("/tasks/:id", h.DeleteTaskHandler)
	

	app.Listen(os.Getenv("PORT"))
}