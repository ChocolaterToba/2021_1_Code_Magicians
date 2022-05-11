package domain

import (
	"time"

	productpb "pinterest/services/product/proto"
)

type Order struct {
	Id              uint64                `json:"id"`
	Items           []ProductWithQuantity `json:"items"`
	CreatedAt       time.Time             `json:"created_at"`
	TotalPrice      uint64                `json:"total_price"`
	PickUp          bool                  `json:"pick_up"`
	DeliveryAddress string                `json:"delivery_address"`
	PaymentMethod   string                `json:"payment_method"`
	CallNeeded      bool                  `json:"call_needed"`
	Status          string                `json:"status"`
}

func ToOrder(pbOrder *productpb.Order) Order {
	return Order{
		Id:              pbOrder.Id,
		Items:           ToProductsWithQuantity(pbOrder.Items),
		CreatedAt:       pbOrder.CreatedAt.AsTime(),
		TotalPrice:      pbOrder.TotalPrice,
		PickUp:          pbOrder.PickUp,
		DeliveryAddress: pbOrder.DeliveryAddress,
		PaymentMethod:   pbOrder.PaymentMethod,
		CallNeeded:      pbOrder.CallNeeded,
		Status:          pbOrder.Status,
	}
}

func ToOrders(pbOrders []*productpb.Order) []Order {
	result := make([]Order, 0, len(pbOrders))
	for _, pbOrder := range pbOrders {
		result = append(result, ToOrder(pbOrder))
	}

	return result
}
