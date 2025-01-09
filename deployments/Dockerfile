FROM golang:1.23-alpine
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/build/main /app/cmd/datingapp/main.go

EXPOSE 8080

CMD ["/app/build/main"]