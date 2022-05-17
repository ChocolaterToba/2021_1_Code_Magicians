package facade

import (
	"bytes"
	"context"
	"io"
	"pinterest/services/product/application"
	"pinterest/services/product/domain"

	pb "pinterest/services/product/proto"
	"time"

	"github.com/pkg/errors"
	_ "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductFacade struct {
	pb.UnimplementedProductServiceServer
	app application.ProductAppInterface
}

func NewProductFacade(app application.ProductAppInterface) *ProductFacade {
	return &ProductFacade{
		app: app,
	}
}

func (facade *ProductFacade) CreateShop(ctx context.Context, in *pb.CreateShopRequest) (*pb.CreateShopResponse, error) {
	shopInput := domain.Shop{
		Title:       in.Title,
		Description: in.Description,
		ManagerIDs:  in.ManagerIds,
	}

	shopID, err := facade.app.CreateShop(ctx, shopInput)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create shop")
	}

	return &pb.CreateShopResponse{Id: shopID}, nil
}

func (facade *ProductFacade) EditShop(ctx context.Context, in *pb.EditShopRequest) (*pb.Empty, error) {
	shopInput := domain.Shop{
		Id:          in.Id,
		Title:       in.Title,
		Description: in.Description,
		ManagerIDs:  in.ManagerIds,
	}

	err := facade.app.EditShop(ctx, shopInput)
	if err != nil {
		return nil, errors.Wrap(err, "Could not edit shop")
	}

	return &pb.Empty{}, nil
}

func (facade *ProductFacade) GetShopByID(ctx context.Context, in *pb.GetShopRequest) (*pb.Shop, error) {
	shop, err := facade.app.GetShopByID(ctx, in.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get shop")
	}

	return domain.ToPbShop(shop), nil
}

func (facade *ProductFacade) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	productInput := domain.Product{
		Id:           0,
		Title:        in.Title,
		Description:  in.Description,
		Price:        in.Price,
		Availability: in.Availability,
		AssemblyTime: in.AssemblyTime,
		PartsAmount:  in.PartsAmount,
		Size:         in.Size,
		Category:     in.Category,
		ShopId:       in.ShopId,
	}

	productID, err := facade.app.CreateProduct(ctx, productInput)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create product")
	}

	return &pb.CreateProductResponse{Id: productID}, nil
}

func (facade *ProductFacade) EditProduct(ctx context.Context, in *pb.EditProductRequest) (*pb.Empty, error) {
	productInput := domain.Product{
		Id:           in.Id,
		Title:        in.Title,
		Description:  in.Description,
		Price:        in.Price,
		Availability: in.Availability,
		AssemblyTime: in.AssemblyTime,
		PartsAmount:  in.PartsAmount,
		Size:         in.Size,
		Category:     in.Category,
		ShopId:       in.ShopId,
	}

	err := facade.app.EditProduct(ctx, productInput)
	if err != nil {
		return nil, errors.Wrap(err, "Could not edit product")
	}

	return &pb.Empty{}, nil
}

const maxPostAvatarsBodySize = 70 * 1024 * 1024 // 70 mB

func (facade *ProductFacade) UpdateProductAvatars(stream pb.ProductService_UpdateProductAvatarsServer) error {
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive product id")
	}

	productID := req.GetProductId()

	filesWithNames := make([]domain.FileWithName, 0)

	fileWithName := domain.FileWithName{}

	req, err = stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive filename")
	}

	fileWithName.Filename = req.GetFilename()

OuterLoop:
	for {
		imageData := bytes.Buffer{}
		imageSize := 0

	InnerLoop:
		for {
			req, err = stream.Recv()
			if err == io.EOF { // last file has been read
				fileWithName.File = &imageData
				filesWithNames = append(filesWithNames, fileWithName)
				fileWithName = domain.FileWithName{Filename: req.GetFilename()}
				break OuterLoop
			}
			if err != nil {
				return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
			}

			if req.GetFilename() != "" { // file is over, next one starts
				fileWithName.File = &imageData
				filesWithNames = append(filesWithNames, fileWithName)
				fileWithName = domain.FileWithName{Filename: req.GetFilename()}
				break InnerLoop
			}

			chunk := req.GetChunk()
			size := len(chunk)

			imageSize += size
			if imageSize > maxPostAvatarsBodySize {
				return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxPostAvatarsBodySize)
			}
			_, err = imageData.Write(chunk)
			if err != nil {
				return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
			}
		}
	}

	err = facade.app.UpdateProductAvatars(context.Background(), productID, filesWithNames)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.Empty{})
}

const maxPostVideoBodySize = 70 * 1024 * 1024 // 8 mB

func (facade *ProductFacade) UpdateProductVideo(stream pb.ProductService_UpdateProductVideoServer) error {
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive product id")
	}

	productID := req.GetProductId()

	req, err = stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive filename")
	}

	filename := req.GetFilename()

	// filenamePrefix := uniuri.NewLen(10) // generating random filename
	// newAvatarPath := "avatars/" + filenamePrefix + req.GetExtension() // TODO: avatars folder sharding by date

	imageData := bytes.Buffer{}
	imageSize := 0
	for {
		req, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}
		chunk := req.GetChunk()
		size := len(chunk)

		imageSize += size
		if imageSize > maxPostVideoBodySize {
			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxPostVideoBodySize)
		}
		_, err = imageData.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}
	}

	err = facade.app.UpdateProductVideo(context.Background(), productID, filename, &imageData)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.Empty{})
}

func (facade *ProductFacade) GetProductByID(ctx context.Context, in *pb.GetProductRequest) (*pb.Product, error) {
	product, err := facade.app.GetProductByID(ctx, in.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get product")
	}

	return domain.ToPbProduct(product), nil
}

func (facade *ProductFacade) GetProductsByIDs(ctx context.Context, in *pb.GetProductsByIDsRequest) (*pb.GetProductsResponse, error) {
	products, err := facade.app.GetProductsByIDs(ctx, in.GetIds())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get products")
	}

	return &pb.GetProductsResponse{
		Products: domain.ToPbProducts(products),
	}, nil
}

func (facade *ProductFacade) GetProducts(ctx context.Context, in *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	products, err := facade.app.GetProducts(ctx, in.GetPageOffset(), in.GetPageSize(), in.GetCategory())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get products")
	}

	return &pb.GetProductsResponse{
		Products: domain.ToPbProducts(products),
	}, nil
}

func (facade *ProductFacade) AddToCart(ctx context.Context, in *pb.AddToCartRequest) (*pb.Empty, error) {
	err := facade.app.AddToCart(ctx, in.GetUserId(), in.GetProductId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not add product to cart")
	}

	return &pb.Empty{}, nil
}

func (facade *ProductFacade) RemoveFromCart(ctx context.Context, in *pb.RemoveFromCartRequest) (*pb.Empty, error) {
	err := facade.app.RemoveFromCart(ctx, in.GetUserId(), in.GetProductId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not remove product from cart")
	}

	return &pb.Empty{}, nil
}

func (facade *ProductFacade) GetCart(ctx context.Context, in *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	cart, err := facade.app.GetCart(ctx, in.GetUserId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get cart")
	}

	productIds := make([]uint64, 0) // is needed because batch-getting products is faster than one-by-one solution
	for id := range cart {
		productIds = append(productIds, id)
	}

	products, err := facade.app.GetProductsByIDs(ctx, productIds)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get products for cart")
	}

	result := make([]*pb.ProductWithQuantity, 0, len(products))
	for _, product := range products {
		result = append(result, &pb.ProductWithQuantity{
			Product:  domain.ToPbProduct(product),
			Quantity: cart[product.Id],
		})
	}

	return &pb.GetCartResponse{
		Products: domain.ToPbProductsWithQuantity(cart, products),
	}, nil
}

func (facade *ProductFacade) CompleteCart(ctx context.Context, in *pb.CompleteCartRequest) (*pb.Empty, error) {
	order := domain.Order{
		Id:              0,
		UserID:          in.GetUserId(),
		Items:           map[uint64]uint64{},
		CreatedAt:       time.Now(),
		TotalPrice:      0,
		PickUp:          in.GetPickUp(),
		DeliveryAddress: in.GetDeliveryAddress(),
		PaymentMethod:   in.GetPaymentMethod(),
		CallNeeded:      in.GetCallNeeded(),
		Status:          "",
	}
	err := facade.app.CompleteCart(ctx, order)
	if err != nil {
		return nil, errors.Wrap(err, "Could not complete cart")
	}

	return &pb.Empty{}, nil
}

func (facade *ProductFacade) GetOrderByID(ctx context.Context, in *pb.GetOrderByIDRequest) (*pb.Order, error) {
	order, err := facade.app.GetOrderByID(ctx, in.GetUserId(), in.GetOrderId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get order by id")
	}

	productIds := make([]uint64, 0) // is needed because batch-getting products is faster than one-by-one solution
	for id := range order.Items {
		productIds = append(productIds, id)
	}

	products, err := facade.app.GetProductsByIDs(ctx, productIds)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get products for order")
	}

	return domain.ToPbOrder(order, products), nil
}

func (facade *ProductFacade) GetOrders(ctx context.Context, in *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	orders, err := facade.app.GetOrdersByUserID(ctx, in.GetUserId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get user's orders")
	}

	productIdsSet := make(map[uint64]bool)
	for _, order := range orders {
		for id := range order.Items {
			productIdsSet[id] = true
		}
	}

	productIds := make([]uint64, 0, len(productIdsSet)) // is needed because batch-getting products is faster than one-by-one solution
	for id := range productIdsSet {
		productIds = append(productIds, id)
	}

	products, err := facade.app.GetProductsByIDs(ctx, productIds)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get products for order")
	}

	return &pb.GetOrdersResponse{
		Orders: domain.ToPbOrders(orders, products),
	}, nil
}
