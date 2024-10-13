package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Product struct { 
	ID primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	FarmerID string `bson:"farmer_id" json:"farmer_id"`
	FarmerName string `bson:"farmer_name" json:"farmer_name"`

	//  string `bson:"name" json:"name"`
	ProductName string `bson:"product_name" json:"product_name"`
	Description string `bson:"description" json:"description"`
	Price string `bson:"price" json:"price"`
	Quantity string `bson:"quantity" json:"quantity"`


	Origin string `bson:"origin" json:"origin"`	


	
	ImageURL string `bson:"image_url" json:"image_url"`	
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`

	Rating float64 `bson:"rating" json:"rating"`
}

type ProductUseCase interface {
	CreateProduct(userID string , product Product) (SuccessResponse, ErrorResponse)
	GetAllProduct(name, limit , page  string) ([]Product, error)


	// GetProductByID(id string) interface{}
	// UpdateProduct(id string, product Product) interface{}
	// DeleteProduct(id string) interface{}
}

type ProductRepository interface {
	CreateProduct(ctx context.Context ,product Product) (Product, error)
	GetProductByID(ctx context.Context , id string) (Product, error)
	GetAllProduct(name, limit , page  string) ([]Product, error)


	// UpdateProduct(id string, product Product) (Product, error)
	// DeleteProduct(id string) (Product, error)
}

