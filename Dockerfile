# Start from a base image with Go installed
FROM golang:1.21.5 as builder

# Set the Current Working Directory inside the container
WORKDIR /

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /main .

# Expose port 6234 to the outside world
EXPOSE 6234

# Command to run the executable
CMD ["./main"]
