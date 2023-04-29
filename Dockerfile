# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Ahamd Syaifudin"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /server
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Production Dev
# Build the Go server
RUN go build -o /build

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "/build" ]