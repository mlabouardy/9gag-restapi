# Start from an official go image with the latest version installed
FROM golang:1.8.0

# Copy the application files
ADD . /go/src/github.com/mlabouardy/9gag-restapi

# Install the dependencies
RUN go get github.com/gorilla/mux github.com/mlabouardy/9gag

# Expose the application on port 3000
EXPOSE 3000

# Set application path as working directory
WORKDIR /go/src/github.com/mlabouardy/9gag-restapi

# Run the application when the container starts
CMD go run app.go
