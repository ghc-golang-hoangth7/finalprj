FROM golang:1.18-alpine

WORKDIR /app
COPY . .

RUN go build -o server ./client/graph/cmd

CMD ["./server"]

EXPOSE 8080
