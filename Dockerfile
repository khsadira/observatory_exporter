# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.8

# Add Maintainer Info
LABEL maintainer="Khan Sadirac <khan.sadirac42@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . /app

# Build the Go app
RUN go build -o http-observatory_exporter

# Expose port 9230 to the outside world
EXPOSE 9230

# Command to run the executable
CMD ["./http-observatory_exporter"]
