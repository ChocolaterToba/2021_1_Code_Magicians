package domain

import (
	productpb "pinterest/services/product/proto"
)

type ProductWithQuantity struct {
	Product  ProductOutput
	Quantity uint64
}

func ToProductWithQuantity(pbProductWithQuantity *productpb.ProductWithQuantity) ProductWithQuantity {
	return ProductWithQuantity{
		Product:  ToProductOutput(ToProduct(pbProductWithQuantity.Product)),
		Quantity: pbProductWithQuantity.Quantity,
	}
}

func ToProductsWithQuantity(pbProductsWithQuantity []*productpb.ProductWithQuantity) []ProductWithQuantity {
	cart := make([]ProductWithQuantity, 0, len(pbProductsWithQuantity))
	for _, pbProductWithQuantity := range pbProductsWithQuantity {
		cart = append(cart, ToProductWithQuantity(pbProductWithQuantity))
	}

	return cart
}
