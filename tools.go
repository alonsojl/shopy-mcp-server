package shopy

import "github.com/mark3labs/mcp-go/mcp"

func (s *MCPServer) addTools() {
	getCategories := mcp.NewTool("get_categories", mcp.WithDescription("Get all categories in store"))
	s.server.AddTool(getCategories, s.HandleGetCategories)

	getProductsByCategoryUUID := mcp.NewTool("get_products_by_category_uuid",
		mcp.WithDescription("Get all products in a category"),
		mcp.WithString("category_uuid",
			mcp.Required(),
			mcp.Description("Category UUID"),
		),
	)
	s.server.AddTool(getProductsByCategoryUUID, s.HandleGetProductsByCategoryUUID)
}
