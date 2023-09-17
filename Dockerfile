FROM golang:1.20

WORKDIR /app

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN cd migrations && goose sqlite3 ../database.db up

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

CMD ["./main"]