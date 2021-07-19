FROM golang:alpine AS builder
RUN apk --update add ca-certificates

WORKDIR /app
COPY main.go go.* /app/
RUN CGO_ENABLED=0 go build -o /app/demo

FROM scratch

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/demo /app/
COPY template.html /app/

EXPOSE 80
ENTRYPOINT ["/app/demo"]
