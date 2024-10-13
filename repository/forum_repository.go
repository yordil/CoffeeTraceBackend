package repository

import (
	"coeffee/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type forumRepository struct {
	database   mongo.Database
	collection string
}


func NewForumRepository(database mongo.Database, collection string) domain.BlogRepository {
	return &forumRepository{
		database:   database,
		collection: collection}
}


//  create blog

func (r *forumRepository) CreateBlog(ctx context.Context , blog domain.Blog) (domain.Blog, error) {

	blog.IsResource = false
	collection  := r.database.Collection(r.collection)

	_, err := collection.InsertOne(ctx, blog)

	if err != nil {

		return domain.Blog{}, err

}

	return blog, nil

}


// get all blog 
func (ur *forumRepository) GetAllBlog(ctx context.Context, bytag, limit, page string) ([]domain.Blog, error) {
		var blogs []domain.Blog

	// Get all blogs where is_resource is true
	filter := bson.M{"isresource": false}

	// Use Find to retrieve multiple documents
	cursor, err := ur.database.Collection(ur.collection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode each document into a Blog
	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	// Check if any errors occurred during iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

	


// get blog by id

func (ur *forumRepository) GetBlogByID(ctx context.Context , id string) (domain.Blog, error) {
	var blog domain.Blog

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}

	err := ur.database.Collection(ur.collection).FindOne(ctx, filter).Decode(&blog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

// create resource

func (ur *forumRepository) CreateResource(ctx context.Context, blog domain.Blog) (domain.SuccessResponse, domain.ErrorResponse) {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(ctx, blog)
	if err != nil {
		return domain.SuccessResponse{}, domain.ErrorResponse{Message: err.Error(), Status: 500}
	}

	return domain.SuccessResponse{Message: "Resource created successfully", Status: 200}, domain.ErrorResponse{}

}

// get resource

func (ur *forumRepository) GetResource(ctx context.Context) ([]domain.Blog, error) {
	var blogs []domain.Blog

	// Get all blogs where is_resource is true
	filter := bson.M{"isresource": true}

	// Use Find to retrieve multiple documents
	cursor, err := ur.database.Collection(ur.collection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode each document into a Blog
	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	// Check if any errors occurred during iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}
