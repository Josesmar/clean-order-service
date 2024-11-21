
# clean-order-service

Este projeto é um serviço de gerenciamento de pedidos, permitindo a criação, consulta por ID e listagem de todos os pedidos. Ele implementa uma API REST, GraphQL e GRPC utilizando a arquitetura hexagonal.

## Requisitos

- **Docker** e **Docker Compose** instalados.
- **Go** versão 1.20 ou superior.
- **Ferramentas adicionais**: gqlgen e protoc (para geração de código).

## Ports

| Serviço      | Descrição                     | Porta  |
|--------------|-------------------------------|--------|
| **Ordersystem** | Servidor Web (API REST)       | `8000` |
| **Ordersystem** | Servidor gRPC                | `50051` |
| **Ordersystem** | Servidor GraphQL             | `8080` |

### Observações

- **Servidor Web (API REST)**: A API principal do sistema está disponível na porta `8000`.
- **Servidor gRPC**: A comunicação baseada em gRPC opera na porta `50051`.
- **Servidor GraphQL**: O endpoint GraphQL está acessível na porta `8080`.

> Certifique-se de verificar as configurações no `docker-compose.yml` caso altere alguma porta.


## Configuração do Ambiente

1. Clone o repositório:
   ```bash
   git clone https://github.com/Josesmar/clean-order-service.git
   cd clean-order-service
   ```

2. Inicie os serviços auxiliares com o Docker Compose:
   ```bash
   docker-compose up -d
   ```

3. Certifique-se de que as dependências estão instaladas:
   ```bash
   go mod tidy
   ```

4. Inicie o serviço:
   ```bash
   go run main.go wire_gen.go
   ```

## REST API

Criar Pedido  
Endpoint: `POST /orders`  
Body:
```json
{
  "id": "order_id",
  "price": 100.0,
  "tax": 10.0
}
```

Consultar Pedido por ID  
Endpoint: `GET /orders/{id}`  

Listar Todos os Pedidos  
Endpoint: `GET /orders`  

## GraphQL

Acesse o playground em `http://localhost:8080/graphql`. Exemplo de query:
```graphql
query {
  listOrders {
    id
    price
    tax
    finalPrice
  }
}
```

## gRPC

1. Gere os arquivos gRPC necessários:
   ```bash
   protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
   ```

2. Certifique-se de que o serviço principal esteja rodando:
   ```bash
   go run main.go
   ```

3. Teste o serviço com Evans CLI:

   - Inicie o cliente Evans em modo REPL:
     ```bash
     evans -r repl
     ```

   - Mostre os pacotes disponíveis:
     ```bash
     show package
     ```

   - Selecione o pacote:
     ```bash
     package pb
     ```

   - Escolha o serviço:
     ```bash
     service OrderService
     ```

   - Faça uma chamada para criar um pedido:
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

   - Para consultar um pedido por ID:
     ```bash
     call GetOrder
     ```

     Insira os dados quando solicitado:
     ```json
     {
       "id": "order_id"
     }
     ```

   - Para listar todos os pedidos:
     ```bash
     call ListOrders
     ```

     Insira os dados quando solicitado:
     ```json
     {}
     ```
