FROM golang:1.16-alpine AS builder

RUN go version
WORKDIR /build
COPY . .

RUN go build -o main .

FROM alpine:latest AS final
RUN apk update && apk add ca-certificates
RUN update-ca-certificates
WORKDIR /app
COPY --from=builder /build .
ENTRYPOINT ["/app/main"]