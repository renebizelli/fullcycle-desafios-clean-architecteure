go version: 1.23.0

Ports:
Web server on port :8000
gRPC on port 50051
GraphQL server on port 8080

Para rodar chamadas web:
Http file on root (api.http)

Par rodar gRPC:
> evans -r repl
127.0.0.1:50051> package pb
pb@127.0.0.1:50051> service OrderService
pb.OrderService@127.0.0.1:50051> call ListOrders

Para rodar playground GRaphql:
http://localhost:8080/