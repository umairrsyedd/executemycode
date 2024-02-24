# Use the official Golang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entrypoint script to the container
COPY entrypoint.sh /app/entrypoint.sh

# Set the executable bit on the entrypoint script
RUN chmod +x entrypoint.sh

# Define the default command to run when the container starts
CMD ["/app/entrypoint.sh"]
