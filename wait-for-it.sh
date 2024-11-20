#!/bin/bash
host="$1"
port="$2"

# Verificar se os parâmetros foram fornecidos
if [ -z "$host" ] || [ -z "$port" ]; then
  echo "Usage: $0 <host> <port> [-- command args]"
  exit 1
fi

# Aguardar o MySQL
echo "Aguardando MySQL..."
while ! nc -z "$host" "$port"; do
  echo "Esperando por $host:$port..."
  sleep 1
done

echo "$host:$port está disponível!"

# Executar a migração
echo "Rodando migrações..."
goose -dir /src/db/migrations mysql 'root:root@tcp(mysql:3306)/orders' up
if [ $? -eq 0 ]; then
  echo "Migração concluída com sucesso!"
else
  echo "Erro ao rodar migrações!"
  exit 1
fi

# Rodar o comando go run
echo "Executando go run..."
exec go run main.go wire_gen.go
