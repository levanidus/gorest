# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="levanidus"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o build cmd/main.go

# Expose port 8000 to the outside world
EXPOSE 8000

# Run the executable
CMD ["./build"]