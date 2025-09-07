# Development
FROM golang:1.25.0-alpine AS development

WORKDIR /app
COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go install github.com/air-verse/air@latest

RUN go mod download 

COPY . .

CMD ["air", "-c", ".air.toml"]

# Builder
FROM golang:1.25.0-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /main /app/main.go

# Production
FROM alpine:3.20 AS production

WORKDIR /app

COPY --from=builder /app/main .

CMD [ "./main" ]