# STEP 1: Build the executable binary

# Use the official Golang image as the build stage
FROM golang:alpine3.18 AS build

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
RUN go build -o fundhub-api-dev .

# STEP 2: Create a smaller image for the final application

# Use a minimal Alpine Linux image as the final stage
FROM alpine:latest

# Set the working directory in the final stage
WORKDIR /app

# Copy only the necessary artifacts from the build stage
COPY --from=build /app/fundhub-api-dev .

# Expose the port if your application listens on a specific port
EXPOSE 8000

# Define the command to run your application
CMD ["./fundhub-api-dev"]
