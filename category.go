package shopy

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

type Categories []Category
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *MCPServer) HandleGetCategories(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	categories := Categories{
		Category{ID: 1, Name: "Category 1"},
		Category{ID: 2, Name: "Category 2"},
		Category{ID: 3, Name: "Category 3"},
		Category{ID: 4, Name: "Category 4"},
		Category{ID: 5, Name: "Category 5"},
		Category{ID: 6, Name: "Category 6"},
		Category{ID: 7, Name: "Category 7"},
		Category{ID: 8, Name: "Category 8"},
		Category{ID: 9, Name: "Category 9"},
		Category{ID: 10, Name: "Category 10"},
	}
	result, err := json.Marshal(categories)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(string(result)), nil
}
