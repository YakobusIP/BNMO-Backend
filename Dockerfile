FROM golang AS builder

COPY ./src /app
WORKDIR /app

RUN go get github.com/go-redis/redis/v9
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors
RUN go get -u gorm.io/gorm

RUN go build -o /main

FROM mariadb
COPY --from=builder /main .

CMD ["./main"]