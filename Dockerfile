# Build
FROM golang:1.20-alpine as builder
RUN mkdir /app
WORKDIR /app
ADD go.mod .
RUN go mod download
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o dmail-server-rest-api ./cmd/dmail-server-rest-api/main.go

# Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/dmail-server-rest-api .
COPY ./config ./config
COPY ./entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

Expose 3000
CMD ["./entrypoint.sh"]