package shopy

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
)

type Products []Product
type Product struct {
	UUID   string  `json:"uuid"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	QRCode string  `json:"qrcode"`
	Image  string  `json:"image"`
	IsTop  bool    `json:"is_top"`
}

func (s *MCPServer) HandleGetProductsByCategoryUUID(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var (
		method   = http.MethodGet
		response struct {
			Products Products `json:"products"`
		}
	)

	uuid, err := request.RequireString("category_uuid")
	if err != nil {
		s.logger.Error("get category uuid", "error", err)
		return mcp.NewToolResultError(err.Error()), nil
	}

	url := s.url + "/v1/products?category_uuid=" + uuid
	data, err := s.httpClient.NewRequest(ctx, method, url, nil).Decode(&response)
	if err != nil {
		s.logger.Error("http client request", "error", err)
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}
