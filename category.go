package shopy

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
)

type Categories []Category
type Category struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

func (s *MCPServer) HandleGetCategories(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var (
		method   = http.MethodGet
		url      = s.url + "/v1/categories"
		response struct {
			Categories Categories `json:"categories"`
		}
	)

	data, err := s.httpClient.NewRequest(ctx, method, url, nil).Decode(&response)
	if err != nil {
		s.logger.Error("http client request", "error", err)
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}
