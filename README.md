# clean-order-service

Este projeto é um serviço de gerenciamento de pedidos, permitindo a criação, consulta por ID e listagem de todos os pedidos. Ele implementa uma API REST, GraphQL e GRPC utilizando a arquitetura hexagonal.

## Requisitos

- **Docker** e **Docker Compose** instalados.
- **Go** versão 1.20 ou superior.
- **Ferramentas adicionais**: gqlgen e protoc (para geração de código).

## Configuração do Ambiente

1. Clone o repositório:
   ```bash
   git clone https://github.com/Josesmar/clean-order-service.git
   cd url

2 - Inicie os serviços auxiliares com o Docker Compose:

   ```bash
docker-compose up -d

Estrutura do docker-compose.yml:
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

3 - Certifique-se de que as dependências estão instaladas:
```bash
go mod tidy
```

4 - Inicie o serviço:
 ```bash
 go run main.go wire_gen.go 
```

REST API
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


GraphQL
Acesse o playground em http://localhost:8080/graphql. Exemplo de query:

query {
  listOrders {
    id
    price
    tax
    finalPrice
  }
}

gRPC
Implemente um cliente gRPC baseado no order.proto disponível no projeto.

