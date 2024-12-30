FROM golang:alpine as build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /myapp ./cmd/exporter
 
FROM alpine:latest as run
COPY --from=build /myapp /myapp

ENTRYPOINT ["/myapp"]
RUN mkdir wheel-logs

