FROM golang:1.18-alpine

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o server cmd/server/main.go

CMD ["./server"]
