FROM 1.24rc3-alpine3.20 AS builder

WORKDIR /api

COPY go.* ./
RUN go mod download

COPY ./src ./src
WORKDIR /api/src
RUN go build -o ./app

FROM alpine:3.20
WORKDIR /api
COPY --from=builder /api/src/app ./app
RUN \
  addgroup -S api-group &&\
  adduser -S api-user -G api-group --no-create-home --disabled-password
USER api-user
CMD [ "./app" ]
EXPOSE ${SERVER_PORT}