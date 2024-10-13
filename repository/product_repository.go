package repository

import (
	"coeffee/domain"
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRepository struct {
	database   mongo.Database
	collection string
}


func NewProductRepository(database mongo.Database, collection string) domain.ProductRepository {
	return &productRepository{
		database:   database,
		collection: collection}
}

func (r *productRepository) CreateProduct(ctx context.Context , product domain.Product) (domain.Product, error) {
	
	collection := r.database.Collection(r.collection)
	_, err := collection.InsertOne(ctx, product)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil

}

// get by id

func (r *productRepository) GetProductByID(ctx context.Context, id string) (domain.Product, error) {
	
	collection := r.database.Collection(r.collection)

	var product domain.Product

	// change the string in to primitive object id
	newid , err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return domain.Product{}, err
	}
	filter := bson.M{"_id": newid}

	err = collection.FindOne(ctx, filter).Decode(&product)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
	
}



func (ur *productRepository) GetAllProduct(byName, limit, page string) ([]domain.Product, error) {
	var products []domain.Product

	// Build the query filter for name search if provided
	filter := bson.M{}
	if byName != "" {
		filter = bson.M{"username": bson.M{"$regex": byName, "$options": "i"}} // Case-insensitive search
	}

	// Convert limit and page to int
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return []domain.Product{}, err
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return []domain.Product{}, err
	}

	// Set pagination options
	options := options.Find()
	options.SetLimit(int64(limitInt)) // Convert limitInt to int64
	options.SetSkip(int64((pageInt - 1) * limitInt)) // Convert pageInt to int64 for skip calculation

	// Query the database with the filter and pagination options
	collection := ur.database.Collection(ur.collection)
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return []domain.Product{}, err
	}
	defer cursor.Close(context.Background())

	// Decode the results into the products slice
	for cursor.Next(context.Background()) {
		var user domain.Product
		if err := cursor.Decode(&user); err != nil {
			return []domain.Product{}, err
		}
		products = append(products, user)
	}

	// Check for any cursor errors after iteration
	if err := cursor.Err(); err != nil {
		return []domain.Product{}, err
	}

	return products, nil
}