package usecase

import (
	"coeffee/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductUseCase struct {
	ProductRepository domain.ProductRepository
	contextTimeout    time.Duration
}

func NewProductUseCase(prodrepo domain.ProductRepository, timeout time.Duration) domain.ProductUseCase {
	return &ProductUseCase{
		ProductRepository: prodrepo,
		contextTimeout: timeout,
	}
}


// create product implmentation

func (uc *ProductUseCase) CreateProduct(userID string , product domain.Product) (domain.SuccessResponse, domain.ErrorResponse) { 
	
	ctx, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()

	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	// product.Status = "active"
	product.FarmerID = userID

	_ , err := uc.ProductRepository.CreateProduct(ctx, product)

	if err != nil {
		return domain.SuccessResponse{},domain.ErrorResponse{Message: "Failed to create product", Status: 400}
	}

	return domain.SuccessResponse{Message: "Product created successfully", Status: 200}, domain.ErrorResponse{}
}




// get all product implementation

func (uc *ProductUseCase) GetAllProduct(name, limit , page  string) ([]domain.Product, error) {
	

	pro, err := uc.ProductRepository.GetAllProduct(name, limit, page)

	if err != nil {
		return []domain.Product{}, err
	}

	return pro, nil
}