# Build stage
FROM golang:1.24 as builder

WORKDIR /app

# Copy go mod and sum files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o kubetag main.go

# Final stage: minimal runtime image
FROM gcr.io/distroless/base

WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /app/kubetag /kubetag

# Run the binary
ENTRYPOINT ["/kubetag"]