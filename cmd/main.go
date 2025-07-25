package main

import (
	"log"
	"log/slog"
	"os"

	"shopy"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	server := server.NewMCPServer(
		"Shopy",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: false,
	}))

	srv := shopy.NewMCPServer(server, logger)
	if err := srv.Run().WithHTTP(); err != nil {
		log.Fatal(err)
	}
}
