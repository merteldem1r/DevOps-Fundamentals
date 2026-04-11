package handlers

import "net/http"

type GlobalHandler struct {
	message string
}

type HandlerResponse struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func NewGlobalHandler(msg string) *GlobalHandler {
	return &GlobalHandler{message: msg}
}

func (h *GlobalHandler) Get(w http.ResponseWriter, r *http.Request) {
	res := HandlerResponse{
		Status: "Success",
		Msg:    h.message,
	}
	SuccessJSON(w, r, res)
}

func (h *GlobalHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	res := HandlerResponse{
		Status: "Success",
		Msg:    "OK",
	}
	SuccessJSON(w, r, res)
}
