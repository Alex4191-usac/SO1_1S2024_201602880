FROM golang:latest as builder

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o grpc_server .


FROM alpine:latest


COPY --from=builder /go/src/app/grpc_server /app/grpc_server


EXPOSE 3000

CMD ["/app/grpc_server"]