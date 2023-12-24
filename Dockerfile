# STEP 1: Build the executable binary

# Use the official Golang image as the build stage
FROM golang:alpine3.18

# Set a label for the maintainer
LABEL maintainer="Titanio Yudista <titanioyudista98@gmail.com>"

# Set the working directory in the build stage
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.* ./
RUN go mod download

# Copy the local code to the container image.
COPY . .

# Build the Go application
RUN go build -o fundhub-api .

# Expose the port if your application listens on a specific port
EXPOSE 8000

# Define the command to run your application
CMD ["./fundhub-api"]