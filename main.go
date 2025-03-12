package main

import (
	"context"
	"os"
	"log"
	"github.com/joho/godotenv"

	"github.com/Liptor/todo_list.git/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

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