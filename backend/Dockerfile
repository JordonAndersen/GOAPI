# Use the official Golang image as a base image
FROM golang:1.19.0

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app

# Install air for live-reloading
RUN go install github.com/cosmtrek/air@latest

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the application
CMD ["air"]
