package cmd

import (
	"log/slog"

	"VideoCalling/internal/server"
)

func main() {
	if errServ := server.Run(); errServ != nil {
		slog.Error("server.Run: ", errServ)
	}
}
