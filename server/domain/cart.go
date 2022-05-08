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
