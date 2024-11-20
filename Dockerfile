FROM golang:1.23

# Instalar dependências necessárias
RUN apt-get update && apt-get install -y netcat-openbsd && rm -rf /var/lib/apt/lists/*

# Definir diretório de trabalho
WORKDIR /src

# Copiar todo o código-fonte do projeto
COPY . .

# Garantir que as dependências Go estão atualizadas
RUN go mod tidy

# Instalar o Wire
RUN go install github.com/google/wire/cmd/wire@latest

# Gerar o arquivo wire_gen.go
RUN /go/bin/wire ./cmd/ordersystem

# Instalar goose para rodar as migrações
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Definir diretório de trabalho para a aplicação principal
WORKDIR /src/cmd/ordersystem

# Copiar o script wait-for-it.sh e torná-lo executável
COPY wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

# Copiar explicitamente o diretório de migrações
COPY db/migrations /src/db/migrations

# Comando para rodar as migrações e depois iniciar a aplicação
CMD /usr/local/bin/wait-for-it.sh mysql:3306 -- goose -dir /src/db/migrations mysql 'root:root@tcp(mysql:3306)/orders' up && go run main.go wire_gen.go