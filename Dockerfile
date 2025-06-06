# Start from the official Go image
FROM golang:1.24 as builder

# Set working directory
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire app source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Use a minimal base image for final stage
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Expose port (optional)
EXPOSE 80

# Set environment variables (optional defaults)
ENV PORT=80
ENV HEAVY_BUILD=false

# Run the application
CMD ["./app"]
