package main

import (
	"log"
	"net/http"

	handler "github.com/jyablonski/go_praq/grpc_project/services/orders/handler/orders"
	"github.com/jyablonski/go_praq/grpc_project/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)
	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
