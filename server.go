package shopy

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/mark3labs/mcp-go/server"
)

const defaultTimeout = 3 * time.Second

type MCPServer struct {
	server     *server.MCPServer
	logger     *slog.Logger
	url        string
	httpClient *http.Client
}

func NewMCPServer(server *server.MCPServer, url string, logger *slog.Logger) *MCPServer {
	return &MCPServer{
		server: server,
		logger: logger,
		url:    url,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (s *MCPServer) Run() *MCPServer {
	s.addTools()
	return s
}

func (s *MCPServer) WithStdio() error {
	return server.ServeStdio(s.server)
}

func (s *MCPServer) WithHTTP() error {
	s.logger.Info("HTTP server listening on :8080/mcp")
	httpServer := server.NewStreamableHTTPServer(s.server)
	return httpServer.Start(":8080")
}
