package types

import (
	"context"

	"github.com/jyablonski/go_praq/grpc_project/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
