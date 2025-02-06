FROM golang:latest AS builder

WORKDIR /app

COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

FROM scratch

COPY --from=builder /app/server .
COPY --from=builder /app/cmd/ordersystem/.env .

EXPOSE 8000 8080 50051 5672 15672

CMD ["./server"]
