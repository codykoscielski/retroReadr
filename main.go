package main

import (
	"log/slog"
	"os"

	"www.github.com/gnomes/retroreadr/internal/handlers"
)

func main() {
    log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

    h := handlers.NewHandler(log)

    err := h.HandleRoutes()
    if err != nil {
        log.Error(err.Error())
    }

}
