# Start from a small, secure base image
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cmd ./cmd/main.go

# Create a minimal production image
FROM alpine:latest

# It's essential to regularly update the packages within the image to include security patches
RUN apk update && apk upgrade

# Reduce image size
RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

# Avoid running code as a root user
RUN adduser -D appuser
USER appuser

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /cmd/app .

# Set any environment variables required by the application
# Env for run server
ENV HTTP_ADDR=:127.0.0.1
ENV HTTP_PORT=:8080
ENV SERVICE_NAME=:'portal_documents_service'
ENV SERVICE_DESCRIPTION=:'portal_documents_service description'
ENV SERVICE_VERSION=:1.0.0


# Expose the port that the application listens on
EXPOSE 8080

# Run the binary when the container starts
CMD ["./app"]