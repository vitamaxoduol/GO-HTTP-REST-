# Use an official Golang runtime as a parent image
FROM golang:1.22.5-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Set Go Proxy
ENV GOPROXY=https://proxy.golang.org

# Install dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]