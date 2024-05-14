# To Build 
# docker build -t umairrsyedd/execute-my-code-server -f server.dockerfile .


# Step 1 

# Use a multi-stage build with Alpine Linux as the base image
FROM golang:1.22 AS build

# Set the working directory
WORKDIR /app

# Copy the entire current directory into the container
COPY . .

# Download dependencies
RUN go mod download

# Build the Go application
RUN GOARCH=amd64 GOOS=linux go build -o /app/server ./cmd/server
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
