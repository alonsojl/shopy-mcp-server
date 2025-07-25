package shopy

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *MCPServer) addTools() {
	// getCategories := mcp.NewTool("get_categories", mcp.WithDescription("Get all categories in store"))
	// s.server.AddTool(getCategories, s.HandleGetCategories)

	getCategories := server.ServerTool{
		Tool:    mcp.NewTool("get_categories", mcp.WithDescription("Get all categories in store")),
		Handler: s.HandleGetCategories,
	}

	s.server.AddTools(getCategories)
}
