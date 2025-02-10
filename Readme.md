go version: 1.23.0

Ports:
Web server on port :8000
gRPC on port 50051
GraphQL server on port 8080

Running:
\cmd\ordersystem> go run .\main.go .\wire_gen.go

Http file on root (api.http)


127.0.0.1:50051> package pb

pb@127.0.0.1:50051> service OrderService

pb.OrderService@127.0.0.1:50051> call ListOrders


http://localhost:8080/