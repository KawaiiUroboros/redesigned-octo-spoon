FROM golang:1.17
EXPOSE 9000
//expose wild card port

WORKDIR /app
COPY . .
RUN go get ./...
RUN go build ./cmd/server
CMD ["./server"]