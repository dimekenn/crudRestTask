FROM golang:latest
RUN go version
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o crud-service .
CMD ["./crud-service"]