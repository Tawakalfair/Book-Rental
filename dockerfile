# Dockerfile

# --- Stage 1: The Builder ---
# Use an official Go image to build our app
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy dependency files and download them
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application.
# CGO_ENABLED=0 is important for a static binary.
# -o /main builds the output file named 'main' in the root directory.
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .


# --- Stage 2: The Final Image ---
# Use a minimal base image for a small and secure final container
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the assets your application needs
COPY views ./views
COPY public ./public

# Copy only the compiled binary from the 'builder' stage
COPY --from=builder /main .

# Expose the port your Fiber app listens on
EXPOSE 3000

# The command to run when the container starts
CMD ["./main"]