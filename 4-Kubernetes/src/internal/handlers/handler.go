package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/merteldem1r/DevOps-Fundamentals/4-Kubernetes/src/internal/config"
	"github.com/merteldem1r/DevOps-Fundamentals/4-Kubernetes/src/internal/models"
)

type GlobalHandler struct {
	cfg    *config.Config
	pool   *pgxpool.Pool
	logger *slog.Logger
}

type HandlerResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewGlobalHandler(cfg *config.Config, pg *pgxpool.Pool, lg *slog.Logger) *GlobalHandler {
	return &GlobalHandler{cfg: cfg, pool: pg, logger: lg}
}

// global
func (h *GlobalHandler) Get(w http.ResponseWriter, r *http.Request) {
	res := HandlerResponse{
		Status: "Success",
		Data:   h.cfg.Message,
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
	query := "SELECT id, title, description, priority FROM todos ORDER BY id"

	rows, err := h.pool.Query(r.Context(), query)
	if err != nil {
		h.logger.Error("failed to fetch todos", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var t models.Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Priority)
		if err != nil {
			h.logger.Error("error on todo row scan", "error", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		todos = append(todos, t)
	}

	if err := rows.Err(); err != nil {
		h.logger.Error("error iterating todos", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res := HandlerResponse{
		Status: "Success",
		Data:   todos,
	}
	SuccessJSON(w, r, res)
}

func (h *GlobalHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	if h.pool == nil {
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}

	var req models.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("failed to decode todo request", "error", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)
	req.Priority = strings.ToUpper(strings.TrimSpace(req.Priority))

	if req.Title == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	switch req.Priority {
	case "HIGH", "MEDIUM", "LOW":
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO todos (title, description, priority)
		VALUES ($1, $2, $3)
		RETURNING id, title, description, priority
	`

	var todo models.Todo
	if err := h.pool.QueryRow(r.Context(), query, req.Title, req.Description, req.Priority).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Priority); err != nil {
		h.logger.Error("failed to create todo", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		h.logger.Error("failed to encode todo response", "error", err)
	}
}
