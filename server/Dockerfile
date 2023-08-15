# Use the official Go image as the base image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the Go files into the container
COPY . .

# Install the dependencies
RUN go mod download

# Build the web server
RUN go build -o main

# Expose the port the web server will run on
EXPOSE 8080

# Start the web server
CMD ["/app/main"]