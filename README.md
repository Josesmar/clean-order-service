# clean-order-service

Este projeto é um serviço de gerenciamento de pedidos, permitindo a criação, consulta por ID e listagem de todos os pedidos. Ele implementa uma API REST, GraphQL e GRPC utilizando a arquitetura hexagonal.

## Requisitos

- **Docker** e **Docker Compose** instalados.
- **Go** versão 1.20 ou superior.
- **Ferramentas adicionais**: gqlgen e protoc (para geração de código).

## Configuração do Ambiente

Clone o repositório:
   ```bash
   git clone https://github.com/Josesmar/clean-order-service.git
   cd url

2 - Inicie os serviços auxiliares com o Docker Compose:

   ```bash
docker-compose up -d
   ```
Estrutura do docker-compose.yml:

   ```docker
services:
  mysql:
    image: mysql:5.7
    container_name: mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: orders
      MYSQL_PASSWORD: root
    ports:
      - 3306:3306
    volumes:
      - mysql_data:/var/lib/mysql

  rabbitmq:
    image: rabbitmq:3-management
    container_name: rabbitmq
    restart: always
    ports:
      - 5672:5672
      - 15672:15672
    environment:
      RABBITMQ_DEFAULT_USER: guest
      RABBITMQ_DEFAULT_PASS: guest

volumes:
  mysql_data:

```

## Certifique-se de que as dependências estão instaladas:
```bash
go mod tidy
```

## Inicie o serviço:
 ```bash
 go run main.go wire_gen.go 
```

## REST API
Criar Pedido
Endpoint: POST /orders
Body:
{
  "id": "order_id",
  "price": 100.0,
  "tax": 10.0
}

Consultar Pedido por ID
Endpoint: GET /orders/{id}
Listar Todos os Pedidos
Endpoint: GET /orders


## GraphQL
Acesse o playground em http://localhost:8080/graphql. Exemplo de query:

query {
  listOrders {
    id
    price
    tax
    finalPrice
  }
}

## gRPC
# Executando gRPC no Order Service

gRPC
Gere os arquivos gRPC necessários:

 ```bash
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
```
Certifique-se de que o serviço principal esteja rodando:

 ```bash
go run main.go
```
Teste o serviço com Evans CLI. Certifique-se de que o Evans está instalado no sistema.

Inicie o cliente Evans em modo REPL:

 ```bash
evans -r repl
```
Mostre os pacotes disponíveis:

 ```bash
show package
```
Selecione o pacote:

 ```bash
package pb
```
Escolha o serviço:

 ```bash
service OrderService
```
Faça uma chamada para criar um pedido:

 ```bash
call CreateOrder
```
Insira os dados quando solicitado:

```json
{
  "id": "order_id",
  "price": 100.0,
  "tax": 10.0
}
```
Para consultar um pedido por ID:

 ```bash
call GetOrder
```
Insira os dados quando solicitado:

```json
{
  "id": "order_id"
}
```
Para listar todos os pedidos:

 ```bash
call ListOrders
```
Insira os dados quando solicitado:

json
{}


