FROM golang:1.23.4

WORKDIR /usr/src/app

# Copy the Go modules manifests and install dependencies (to leverage caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o app .

# Expose the port your application listens on
EXPOSE 3020

# Command to run the application
CMD ["./app"]