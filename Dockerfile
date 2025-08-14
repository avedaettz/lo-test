FROM golang:1.24-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .

 
RUN go build -o task-api ./cmd

EXPOSE 8080

CMD ["./task-api"]