# Start from a newer golang base image that supports Go 1.22
FROM golang:1.22-alpine as builder

# Install build dependencies including GCC
RUN apk add --no-cache gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -a -o main .

# Start a new stage from alpine for smaller final image
FROM alpine:latest

# Install required dependencies for SQLite
RUN apk --no-cache add ca-certificates libc6-compat sqlite

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .
# Copy the database directory
COPY --from=builder /app/database ./database

# Expose port
EXPOSE 3000

# Command to run the executable 
CMD ["./main"]