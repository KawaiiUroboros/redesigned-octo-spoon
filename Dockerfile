
FROM golang:1.17

ENV PORT=9000
EXPOSE 9000
EXPOSE 3000
WORKDIR /app
COPY . .
RUN go get ./...
RUN go build ./cmd/server
CMD ["./server"]