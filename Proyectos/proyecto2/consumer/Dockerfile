
FROM golang:latest AS builder

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consumer .

FROM alpine:latest

COPY --from=builder /go/src/app/consumer /app/consumer

EXPOSE 3000

CMD ["/app/consumer"]