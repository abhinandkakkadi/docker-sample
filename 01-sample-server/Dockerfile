# Specify the base image
FROM golang:1.16-alpine

# Setup working directory
WORKDIR /app

# copy everything in the current directory to temporory container
COPY go.mod ./

# go mod download
RUN go mod download

# Copy go files into the temporary container
COPY *.go ./

# Create build for the application
RUN go build -o /go-sample

# Run the executable
CMD [ "/go-sample" ]