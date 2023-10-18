FROM golang:latest as builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o /go-sample

FROM nginx 

# The path is where nginx gives the static response
COPY --from=builder /go-sample /usr/share/nginx/html 

# for running - docker run -p 8080:80 <image-name> . (80 is the deafault port to which nginx listens to)
# This will not work as we don't have a static content here. It's a web server