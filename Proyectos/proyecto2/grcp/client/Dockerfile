FROM golang:latest as builder

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o grpc_client .


FROM alpine:latest


COPY --from=builder /go/src/app/grpc_client /app/grpc_client


EXPOSE 3000

CMD ["/app/grpc_client"]