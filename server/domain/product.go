package domain

import (
	"fmt"
	productpb "pinterest/services/product/proto"
)

type Product struct {
	Id           uint64 `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Price        uint64 `json:"price"`
	Availability bool   `json:"availability"`
	// AssemblyTime is measured in minutes
	AssemblyTime uint64   `json:"assembly_time"`
	PartsAmount  uint64   `json:"parts_amount"`
	Rating       float64  `json:"rating"`
	Size         string   `json:"size"`
	Category     string   `json:"category"`
	ImageLinks   []string `json:"image_links"`
	VideoLink    string   `json:"video_link"`
	ShopId       uint64   `json:"shop_id"`
}

type ProductOutput struct {
	Id           uint64 `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Price        uint64 `json:"price"`
	Availability bool   `json:"availability"`
	// AssemblyTime is a "%d hours %d minutes" string here
	AssemblyTime string   `json:"assembly_time"`
	PartsAmount  uint64   `json:"parts_amount"`
	Rating       float64  `json:"rating"`
	Size         string   `json:"size"`
	Category     string   `json:"category"`
	ImageLinks   []string `json:"image_links"`
	VideoLink    string   `json:"video_link"`
	ShopId       uint64   `json:"shop_id"`
}

type ProductWithQuantity struct {
	Product  ProductOutput
	Quantity uint64
}

type ProductIDResponse struct {
	ProductID uint64 `json:"product_id"`
}

func ToProductOutput(product Product) (output ProductOutput) {
	return ProductOutput{
		Id:           product.Id,
		Title:        product.Title,
		Description:  product.Description,
		Price:        product.Price,
		Availability: product.Availability,
		AssemblyTime: fmt.Sprintf("%d часов %d минут", product.AssemblyTime/60, product.AssemblyTime%60),
		PartsAmount:  product.PartsAmount,
		Rating:       product.Rating,
		Size:         product.Size,
		Category:     product.Category,
		ImageLinks:   product.ImageLinks,
		VideoLink:    product.VideoLink,
		ShopId:       product.ShopId,
	}
}

func ToProductOutputs(products []Product) (outputs []ProductOutput) {
	outputs = make([]ProductOutput, 0, len(products))

	for _, product := range products {
		outputs = append(outputs, ToProductOutput(product))
	}
	return outputs
}

func ToProduct(pbProduct *productpb.Product) Product {
	return Product{
		Id:           pbProduct.Id,
		Title:        pbProduct.Title,
		Description:  pbProduct.Description,
		Price:        pbProduct.Price,
		Availability: pbProduct.Availability,
		AssemblyTime: pbProduct.AssemblyTime,
		PartsAmount:  pbProduct.PartsAmount,
		Rating:       pbProduct.Rating,
		Size:         pbProduct.Size,
		Category:     pbProduct.Category,
		ImageLinks:   pbProduct.ImageLinks,
		VideoLink:    pbProduct.VideoLink,
		ShopId:       pbProduct.ShopId,
	}
}

func ToProducts(pbProducts []*productpb.Product) []Product {
	result := make([]Product, 0, len(pbProducts))

	for _, pbProduct := range pbProducts {
		result = append(result, ToProduct(pbProduct))
	}

	return result
}

func ToPbCreateProductRequest(product Product) *productpb.CreateProductRequest {
	return &productpb.CreateProductRequest{
		Title:        product.Title,
		Description:  product.Description,
		Price:        product.Price,
		Availability: product.Availability,
		AssemblyTime: product.AssemblyTime,
		PartsAmount:  product.PartsAmount,
		Size:         product.Size,
		Category:     product.Category,
		ShopId:       product.ShopId,
	}
}

func ToPbEditProductRequest(product Product) *productpb.EditProductRequest {
	return &productpb.EditProductRequest{
		Id:           product.Id,
		Title:        product.Title,
		Description:  product.Description,
		Price:        product.Price,
		Availability: product.Availability,
		AssemblyTime: product.AssemblyTime,
		PartsAmount:  product.PartsAmount,
		Size:         product.Size,
		Category:     product.Category,
		ShopId:       product.ShopId,
	}
}

func ToProductWithQuantity(pbProductWithQuantity *productpb.ProductWithQuantity) ProductWithQuantity {
	return ProductWithQuantity{
		Product:  ToProductOutput(ToProduct(pbProductWithQuantity.Product)),
		Quantity: pbProductWithQuantity.Quantity,
	}
}
