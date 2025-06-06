# GitHub Actions Go Server

A simple Go web server that demonstrates GitHub Actions integration and containerization. The server provides a basic HTTP endpoint that can optionally perform CPU-intensive calculations.

## Features

- Simple HTTP server with hostname response
- Optional CPU-intensive prime number calculation
- Docker containerization
- Environment variable configuration
- GitHub Actions integration

## Prerequisites

- Go 1.24.0 or later
- Docker (for containerization)
- Git

## Project Structure

```
.
├── Dockerfile          # Container configuration
├── main.go            # Main application code
├── go.mod             # Go module definition
└── go.sum             # Go module checksums
```

## Getting Started

### Local Development

1. Clone the repository:
   ```bash
   git clone https://github.com/andymartinezot/github-actions-go-server.git
   cd github-actions-go-server
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

The server will start on port 80 by default. You can change the port by setting the `PORT` environment variable.

### Using Docker

1. Build the Docker image:
   ```bash
   docker build -t github-actions-go-server .
   ```

2. Run the container:
   ```bash
   docker run -p 80:80 github-actions-go-server
   ```

## Environment Variables

- `PORT`: Server port (default: 80)
- `HEAVY_BUILD`: Enable CPU-intensive calculation (default: false)

Example with environment variables:
```bash
docker run -p 8080:8080 -e PORT=8080 -e HEAVY_BUILD=true github-actions-go-server
```

## API Endpoints

### GET /

Returns a greeting message with the server's hostname.

- Normal response: `Hello from <hostname>!`
- With `HEAVY_BUILD=true`: `Hello from <hostname>! (calculated <count> prime numbers during build)`


## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
