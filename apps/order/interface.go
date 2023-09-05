package order

import (
	"context"
	"restful-api-demo/apps/order/model"
)

type OrderInterface interface {
	SaveOrder(context.Context, *model.Order) (*model.Order, error)
	QueryOrder(context.Context, *model.QueryOrderListRequest) (*model.OrderSet, error)
	DescribeOrder(context.Context, *model.DescribeOrderRequest) (*model.Order, error)
	DeleteOrder(context.Context, *model.DeleteOrderRequest) (*model.Order, error)
	UpdateOrder(context.Context, *model.UpdateOrderRequest) (*model.Order, error)
}
