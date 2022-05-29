ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine

WORKDIR /app

ARG SERVICE_NAME

COPY ../../api/${SERVICE_NAME} /app/

RUN go mod download && \
    go build -o /${SERVICE_NAME}

EXPOSE 8080

CMD ["/${SERVICE_NAME}" ]
