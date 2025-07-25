package shopy

import "github.com/mark3labs/mcp-go/mcp"

func (s *MCPServer) addTools() {
	getCategories := mcp.NewTool("get_categories", mcp.WithDescription("Get all categories in store"))
	s.server.AddTool(getCategories, s.HandleGetCategories)
}
