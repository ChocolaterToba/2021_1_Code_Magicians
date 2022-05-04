package domain

import (
	productpb "pinterest/services/product/proto"
)

type Shop struct {
	Id          uint64   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ManagerIDs  []uint64 `json:"manager_ids"`
}

func ToShop(pbShop *productpb.Shop) Shop {
	return Shop{
		Id:          pbShop.Id,
		Title:       pbShop.Title,
		Description: pbShop.Description,
		ManagerIDs:  pbShop.ManagerIds,
	}
}

func ToPbCreateShopRequest(shop Shop) *productpb.CreateShopRequest {
	return &productpb.CreateShopRequest{
		Title:       shop.Title,
		Description: shop.Description,
		ManagerIds:  shop.ManagerIDs,
	}
}

func ToPbEditShopRequest(shop Shop) *productpb.EditShopRequest {
	return &productpb.EditShopRequest{
		Id:          shop.Id,
		Title:       shop.Title,
		Description: shop.Description,
		ManagerIds:  shop.ManagerIDs,
	}
}

type ShopIDResponse struct {
	ShopID uint64 `json:"shop_id"`
}
