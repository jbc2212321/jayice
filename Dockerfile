# /Dockerfile

# Step 1: Build the Go binary
FROM golang:alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules and Sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Step 2: Run the Go binary
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8096 to the outside world
ENV PORT=8096

# Command to run the executable
CMD ["./main"]
