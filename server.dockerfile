# Step 1 

# Use a multi-stage build with Alpine Linux as the base image
FROM golang:1.22-alpine AS build

# Set the working directory
WORKDIR /app

# Copy the entire current directory into the container
COPY . .

# Download dependencies
RUN go mod download

# Build the Go application
RUN go build ./cmd/server
# Step 2

# Use a minimal Alpine image as the final base image
FROM alpine:latest AS final

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the previous stage
COPY --from=build /app/server .

# Expose the port the server listens on
EXPOSE 4549

# Command to run the server
CMD ["./server"]
