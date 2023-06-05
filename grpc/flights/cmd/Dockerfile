FROM golang:1.18-alpine

WORKDIR /app
COPY . .

RUN go build -o server ./grpc/flights/cmd

CMD ["./server"]
