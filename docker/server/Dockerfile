FROM golang:alpine3.20 AS builder

WORKDIR /api

COPY go.* ./
RUN go mod download

COPY ./src ./src
WORKDIR /api/src
RUN go build -o ./app

FROM alpine:3.20
WORKDIR /api
COPY --from=builder /api/src/app ./app
CMD [ "./app" ]
EXPOSE ${SERVER_PORT}