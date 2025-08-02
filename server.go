package shopy

import (
	"log/slog"

	"github.com/mark3labs/mcp-go/server"
)

type MCPServer struct {
	server     *server.MCPServer
	logger     *slog.Logger
	url        string
	httpClient *HTTPClient
}

func NewMCPServer(server *server.MCPServer, url string, logger *slog.Logger) *MCPServer {
	return &MCPServer{
		server:     server,
		logger:     logger,
		url:        url,
		httpClient: NewHTTPClient(),
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
