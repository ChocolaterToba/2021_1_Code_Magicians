package domain

import (
	pb "pinterest/services/shopProduct/proto"
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
		ShopId:       product.ShopId,
	}
}
