FROM golang:1.23.0-alpine

WORKDIR /home/shopy/mcp-server

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN apk update && apk add --no-cache build-base git zip curl
RUN make tools
