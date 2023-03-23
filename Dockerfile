FROM golang:alpine AS go
WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . /app/
RUN go build -o stats_server

FROM node:alpine AS node
WORKDIR /app
COPY ./front/package.json /app/
RUN npm i
COPY ./front /app/
RUN npm run build

FROM alpine 
WORKDIR /app
ENV NODE_ENV=production
ENV PORT_HTTP=3000
ENV REDIS_ADDRESS="192.168.0.103:6379"
ENV REDIS_PASS=""
ENV REDIS_DB=0
ENV LOCAL_PING_ADDRESS="192.168.0.1"
ENV REMOTE_PING_ADDRESS="1.1.1.1"
ENV PING_DURATION=1
ENV PING_COUNT=1
ENV SLEEP_TIME=10
USER root:root
RUN apk update && apk add tzdata ca-certificates
COPY --from=go /app/stats_server /app/
COPY --from=node /app/out/ /app/front
ENTRYPOINT ["/app/stats_server"]