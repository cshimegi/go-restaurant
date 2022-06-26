ARG GO_VERSION
FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /go/src/service

ARG SERVICE_NAME
COPY ../../api /go/src/service
RUN cd  /go/src/service/${SERVICE_NAME} && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o /bin/server


FROM scratch AS final
COPY --from=builder /bin/server /bin/server
EXPOSE 8000
ENTRYPOINT ["/bin/server" ]
