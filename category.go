package shopy

import (
	"context"
	"encoding/json"
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
		method = http.MethodGet
		url    = s.url + "/v1/categories"
	)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	defer resp.Body.Close()

	var response struct {
		Categories Categories `json:"categories"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	data, err := json.Marshal(response)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}
