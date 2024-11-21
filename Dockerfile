FROM golang:1.23

RUN apt-get update && apt-get install -y netcat-openbsd && rm -rf /var/lib/apt/lists/*

# Instalar o Wire
RUN go install github.com/google/wire/cmd/wire@latest

# Instalar Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /src

COPY . .

RUN go mod tidy

RUN /go/bin/wire ./cmd/ordersystem

WORKDIR /src/cmd/ordersystem

COPY db/migrations /src/db/migrations

EXPOSE 8000
EXPOSE 8080

# Rodar migrações e iniciar a aplicação
CMD goose -dir /src/db/migrations mysql 'root:root@tcp(mysql:3306)/orders' up && go run main.go wire_gen.go
