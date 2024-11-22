# Start from the official Go image
FROM docker.io/golang:1.23.3-alpine3.20

# Set the working directory in the container
WORKDIR /usr/app

# Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the working directory
COPY . .

# Build the Go application
RUN go build -o goapp ./main.go

# Make the binary executable (optional)
RUN chmod +x goapp

# Expose the port your app runs on (e.g., 8080)
EXPOSE 8080

# Run the built application
CMD ["./goapp"]