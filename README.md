# Shopy MCP Server

Shopy MCP Server is a backend service written in Go that provides APIs for managing categories and products in an e-commerce platform. It is designed to be lightweight, efficient, and easy to deploy using Docker.

## Available Tools

The following tools are available via the MCP server:

- **get_categories**: Get all categories in store.
- **get_products_by_category_uuid**: Get all products in a category. Requires a parameter:
  - `category_uuid` (string, required): Category UUID.
