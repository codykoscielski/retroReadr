package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Handler struct {
    log *slog.Logger
}

func NewHandler(log *slog.Logger) *Handler {
    return &Handler{log: log}
}

func (h *Handler) HandleRoutes() error {

    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("hi")
    })

    fmt.Println("Server is on port 3000")
    return http.ListenAndServe(":3000", mux)
}
