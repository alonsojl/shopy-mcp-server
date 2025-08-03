# Shopy MCP Server

Shopy MCP Server is a backend service written in Go that provides APIs for managing categories and products in an e-commerce platform. It is designed to be lightweight, efficient, and easy to deploy using Docker.

## Available Tools

The following tools are available via the MCP server:

- **get_categories**: Get all categories in store.
- **get_products_by_category_uuid**: Get all products in a category. Requires a parameter:
  - `category_uuid` (string, required): Category UUID.

## Makefile Commands

The `Makefile` provides several commands to help with development and deployment:

- `make help` — Show all available make targets and their descriptions.
- `make up` — Start container services using Docker Compose.
- `make down` — Stop container services.
- `make linter` — Lint the source code using golangci-lint. Output is saved to `linter.txt`.
- `make clean` — Clean build files and cache.
- `make build` — Build the application binary to `./bin/mcp-server` (runs `clean` first).
- `make run` — Build and run the application.
- `make live` — Start live reload for Go applications using Air (requires Air to be installed).
- `make inspector` — Open the ModelContext inspector using npx.
- `make tools` — Install development tools (Air and golangci-lint).

Refer to the `Makefile` for more details on each command.