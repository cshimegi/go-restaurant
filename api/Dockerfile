ARG GO_VERSION
FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /app

ARG PORT
ENV PORT=${PORT} GO111MODULE=on CGO_ENABLED=0 GOOS=linux

COPY ./api .
RUN go mod download && \
    go build -o /app/api-server ./server/server.go


FROM alpine:3.19.1 AS final

USER root

ARG PORT
ENV PORT=${PORT}

WORKDIR /app
COPY --from=builder /app/api-server .

RUN apk add --update --no-cache curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-386.tar.gz | tar xvz && \
    mkdir -p ./migrations

COPY --from=builder /app/migrations ./migrations/

EXPOSE ${PORT}
CMD ["./api-server"]
