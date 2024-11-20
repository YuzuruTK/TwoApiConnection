# Start from the official Go image
FROM golang:1.23.3-alpine3.20

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the working directory
COPY . .

# Set the default command to build and run the Go app
CMD ["go", "run", "."]