FROM golang:latest

WORKDIR /go/src/github.com/bnmo-backend
COPY . .

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o main .

CMD ["./main"]