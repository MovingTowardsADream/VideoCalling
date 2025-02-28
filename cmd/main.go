package main

import (
	"log/slog"
	"net/http"

	"VideoCalling/server"
)

func main() {
	http.Handle("/create", http.HandlerFunc(server.CreateRoomHandler))
	http.Handle("/join", http.HandlerFunc(server.JoinRoomHandler))

	slog.Info("starting server on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("failed server", err)
	}
}
