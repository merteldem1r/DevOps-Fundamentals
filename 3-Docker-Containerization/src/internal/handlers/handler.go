package handlers

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/models"
)

type GlobalHandler struct {
	message string
	pool    *pgxpool.Pool
	logger  *slog.Logger
}

type HandlerResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewGlobalHandler(msg string, pg *pgxpool.Pool, lg *slog.Logger) *GlobalHandler {
	return &GlobalHandler{message: msg, pool: pg, logger: lg}
}

// global
func (h *GlobalHandler) Get(w http.ResponseWriter, r *http.Request) {
	res := HandlerResponse{
		Status: "Success",
		Data:   h.message,
	}
	SuccessJSON(w, r, res)
}

func (h *GlobalHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	res := HandlerResponse{
		Status: "Success",
		Data:   "OK",
	}
	SuccessJSON(w, r, res)
}

// db interaction
func (h *GlobalHandler) GetTodos(w http.ResponseWriter, r *http.Request) {

	rows, err := h.pool.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var t models.Todo
		err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Priority)
		if err != nil {
			h.logger.Error("Error on todo row scan", "error", err)
		}
		todos = append(todos, t)
	}

	SuccessJSON(w, r, todos)
}

func (h *GlobalHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {

}
