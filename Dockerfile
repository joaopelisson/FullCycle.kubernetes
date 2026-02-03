FROM golang:1.25.6

COPY . .

RUN go build -o server .

CMD ["./server"]