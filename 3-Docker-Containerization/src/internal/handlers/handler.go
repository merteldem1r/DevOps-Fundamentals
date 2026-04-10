package handlers

import "net/http"

type GlobalHandler struct{}

type HealthResponse struct {
	status string
	msg    string
}

func NewGlobalHandler() *GlobalHandler {
	return &GlobalHandler{}
}

func (h *GlobalHandler) Get(w http.ResponseWriter, r *http.Request) {
	msg := "<3"
	SuccessJSON(w, r, msg)
}

func (h *GlobalHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	health := HealthResponse{status: "Success", msg: "OK"}
	SuccessJSON(w, r, health)
}
