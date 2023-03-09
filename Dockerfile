FROM golang:alpine AS go
WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download
COPY . /app/
RUN go build -o stats_server

FROM alpine 
WORKDIR /app
RUN apk update && apk add tzdata ca-certificates
COPY --from=go /app/stats_server /app/
ENV PORT_HTTP=3000
USER root:root
ENTRYPOINT ["/app/stats_server"]