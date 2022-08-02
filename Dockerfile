FROM golang:latest AS builder

COPY ./src /app
WORKDIR /app

COPY /src/go.mod .
COPY /src/go.sum .

RUN go mod download
COPY . .

RUN go build -o /main

FROM mariadb
COPY --from=builder /main .

CMD ["./main"]