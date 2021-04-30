FROM golang:1.16

WORKDIR /app
COPY . .

RUN go build -o main .
CMD ["/app/main"]