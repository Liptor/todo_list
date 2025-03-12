package handlers

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gofiber/fiber/v2"
)

type Handler struct  {
	DB *pgxpool.Pool
}	

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{DB: db}
}

type Task struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Status 	string `json:"status"`
}

func (h *Handler) CreateTaskHandler(c *fiber.Ctx) error {
	taskdata := new(Task)

	if err := c.BodyParser(taskdata); err != nil {
		return err
	}

	conn, err := h.DB.Acquire(context.Background())

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Unable to acquire a database connection: %v\n")
	}
	
	defer conn.Release()

	conn.QueryRow(context.Background(), 
		`INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)`, 
		taskdata.Title,
		taskdata.Description,
		taskdata.Status,
	)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Task added seccessfully",
		"data": taskdata,
	})
}

func (h *Handler) GetTaskHandler(c *fiber.Ctx) error {
	query := `SELECT id, title, description, status FROM tasks`

	rows, err := h.DB.Query(context.Background(), query)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprint(err))
	}

	defer rows.Close()

	items := []map[string]interface{}{} 
 
	for rows.Next() {
		var id int
		var title, description, status string

		if err := rows.Scan(&id, &title, &description, &status); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to scan items")
		}

		items = append(items, map[string]interface{}{
			"id": id,
			"title": title,
			"description": description,
			"status": status,
		})
	}

	return c.JSON(fiber.Map{
		"tasks": items,
	})
}

func (h *Handler) UpdateTaskHandler(c *fiber.Ctx) error {
	var taskdata Task
	id := c.Params("id")

	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Id is necessary")
	} 


	if err := c.BodyParser(&taskdata); err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	conn, err := h.DB.Acquire(context.Background())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), 
		"UPDATE tasks SET title = $2, description = $3, status = $4 WHERE id = $1",
		id, taskdata.Title, taskdata.Description, taskdata.Status,
	)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprint(err))
	}

	if ct.RowsAffected() == 0 {
		return fiber.ErrNotFound
	}

	return c.Status(fiber.StatusAccepted).SendString("Item Changed Successfully!")
}

func (h *Handler) DeleteTaskHandler(c *fiber.Ctx) error {
	res := c.Params("id")

	if res == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Id is necessary")
	} 

	conn, err := h.DB.Acquire(context.Background())

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	defer conn.Release()


	dbquery := `DELETE FROM tasks WHERE id=$1`

	ct, err := h.DB.Exec(context.Background(), dbquery, res)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete task")
	}

	if ct.RowsAffected() == 0 {
		return fiber.ErrNotFound
	}



	return c.Status(fiber.StatusAccepted).SendString("Item deleted successfully")
}

