# Specify the base image
FROM golang:1.20-alpine

# copy everything in the current directory to temporory container
COPY go.mod ./

# go mod download
RUN go mod download


# 