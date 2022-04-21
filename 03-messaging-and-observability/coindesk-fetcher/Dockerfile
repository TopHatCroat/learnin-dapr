FROM golang:1.17-alpine3.15 as builder

WORKDIR /dir

COPY main.go .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#second stage
FROM alpine:3.15

WORKDIR /root/

COPY --from=builder /dir/main .

CMD ["./main"]
