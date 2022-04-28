FROM golang:1.17-alpine3.15

WORKDIR /dir

COPY . .

RUN go get -d -v

CMD ["go", "run", "main.go"]
