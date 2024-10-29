# README

## Introdução

Este projeto é uma aplicação que utiliza a arquitetura limpa (Clean Architecture) para gerenciar pedidos. Ele implementa uma API REST, um servidor gRPC e um servidor GraphQL para interação com os dados de pedidos.

## Serviços

A aplicação possui os seguintes serviços:

1. **Web Server (API REST)**:
   - **Descrição**: Um servidor HTTP que fornece endpoints para criar e recuperar pedidos.
   - **Porta**: `webServerPort` (definida nas configurações).
   - **Endpoints**:
     - `POST /order`: Cria um novo pedido.
     - `GET /orders`: Recupera todos os pedidos.

2. **gRPC Server**:
   - **Descrição**: Um servidor gRPC que permite a criação e recuperação de pedidos através de chamadas gRPC.
   - **Porta**: `gRPCServerPort` (definida nas configurações).
   - **Serviço**: `OrderService` com métodos para criar e listar pedidos.

3. **GraphQL Server**:
   - **Descrição**: Um servidor GraphQL que permite consultas e mutações relacionadas aos pedidos.
   - **Porta**: `graphQLServerPort` (definida nas configurações).
   - **Endpoints**:
     - `GET /`: Página do playground do GraphQL, onde é possível explorar as consultas e mutações.
     - `POST /query`: Endpoint para executar consultas e mutações GraphQL.

## Configuração

Antes de executar a aplicação, é necessário configurar as seguintes variáveis no arquivo de configuração:

- `DBDriver`: O driver do banco de dados (ex: `mysql`).
- `DBUser`: O usuário do banco de dados.
- `DBPassword`: A senha do banco de dados.
- `DBHost`: O endereço do host do banco de dados.
- `DBPort`: A porta do banco de dados (ex: `3306` para MySQL).
- `DBName`: O nome do banco de dados.
- `WebServerPort`: A porta para o servidor web.
- `gRPCServerPort`: A porta para o servidor gRPC.
- `GraphQLServerPort`: A porta para o servidor GraphQL.

### Exemplo de configuração

```ini
DBDriver=mysql
DBUser=root
DBPassword=example
DBHost=localhost
DBPort=3306
DBName=mydatabase
WebServerPort=8080
gRPCServerPort=50051
GraphQLServerPort=8081
```

### Rodando o App
`docker compose up -d` Para levantar rabbitMQ/MySQL
`cd cmd/ordersystem && go run main.go wire_gen.go` Para subir server

