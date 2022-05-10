package domain

import (
	pb "pinterest/services/product/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToShop(pbShop *pb.Shop) Shop {
	return Shop{
		Id:          pbShop.GetId(),
		Title:       pbShop.GetTitle(),
		Description: pbShop.GetDescription(),
		ManagerIDs:  pbShop.GetManagerIds(),
	}
}

func ToPbShop(shop Shop) *pb.Shop {
	return &pb.Shop{
		Id:          shop.Id,
		Title:       shop.Title,
		Description: shop.Description,
		ManagerIds:  shop.ManagerIDs,
	}
}

func ToProduct(pbProduct *pb.Product) Product {
	return Product{
		Id:           pbProduct.GetId(),
		Title:        pbProduct.GetTitle(),
		Description:  pbProduct.GetDescription(),
		Price:        pbProduct.GetPrice(),
		Availability: pbProduct.GetAvailability(),
		AssemblyTime: pbProduct.GetAssemblyTime(),
		PartsAmount:  pbProduct.GetPartsAmount(),
		Rating:       pbProduct.GetRating(),
		Size:         pbProduct.GetSize(),
		Category:     pbProduct.GetCategory(),
		ImageLinks:   pbProduct.GetImageLinks(),
		VideoLink:    pbProduct.GetVideoLink(),
		ShopId:       pbProduct.GetShopId(),
	}
}

func ToPbProduct(product Product) *pb.Product {
	return &pb.Product{
		Id:           product.Id,
		Title:        product.Title,
		Description:  product.Description,
		Price:        product.Price,
		Availability: product.Availability,
		AssemblyTime: product.AssemblyTime,
		PartsAmount:  product.PartsAmount,
		Rating:       product.Rating,
		Size:         product.Size,
		Category:     product.Category,
		ImageLinks:   product.ImageLinks,
		VideoLink:    product.VideoLink,
		ShopId:       product.ShopId,
	}
}

func ToPbProducts(products []Product) []*pb.Product {
	result := make([]*pb.Product, 0, len(products))

	for _, product := range products {
		result = append(result, ToPbProduct(product))
	}

	return result
}

func ToPbProductsWithQuantity(cart map[uint64]uint64, products []Product) []*pb.ProductWithQuantity {
	result := make([]*pb.ProductWithQuantity, 0, len(products))
	for _, product := range products {
		if cart[product.Id] > 0 { // products slice may contain excessive products
			result = append(result, &pb.ProductWithQuantity{
				Product:  ToPbProduct(product),
				Quantity: cart[product.Id],
			})
		}
	}

	return result
}

func ToPbOrder(order Order, products []Product) *pb.Order {
	return &pb.Order{
		Id:              order.Id,
		Items:           ToPbProductsWithQuantity(order.Items, products),
		CreatedAt:       timestamppb.New(order.CreatedAt),
		TotalPrice:      order.TotalPrice,
		PickUp:          order.PickUp,
		DeliveryAddress: order.DeliveryAddress,
		PaymentMethod:   order.PaymentMethod,
		CallNeeded:      order.CallNeeded,
		Status:          order.Status,
	}
}

func ToPbOrders(orders []Order, products []Product) []*pb.Order {
	result := make([]*pb.Order, 0, len(orders))

	for _, order := range orders {
		result = append(result, ToPbOrder(order, products))
	}

	return result
}
