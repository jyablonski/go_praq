package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/jyablonski/go_praq/grpc_project/services/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()
	stop := make(chan os.Signal, 1)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 24,
			ProductID:  3123,
			Quantity:   2,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		res, err := c.GetOrders(ctx, &orders.GetOrdersRequest{
			CustomerID: 42,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))

		if err := t.Execute(w, res.GetOrders()); err != nil {
			log.Fatalf("template error: %v", err)
		}
	})

	log.Println("Starting server on", s.addr)

	// Run the HTTP server in a goroutine
	go func() {
		if err := http.ListenAndServe(s.addr, router); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-stop
	log.Println("Received shutdown signal, shutting down gracefully...")

	// You can add additional cleanup here if needed (e.g., closing connections)
	return nil
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
