FROM golang:1.18-alpine

WORKDIR /app
COPY . .

RUN go build -o server ./grpc/planes/cmd

CMD ["./server"]

EXPOSE 50051
