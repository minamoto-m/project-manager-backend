FROM golang:1.24.1

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . .

CMD ["air"]