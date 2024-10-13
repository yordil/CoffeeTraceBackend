package repository

import (
	"coeffee/domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type replyRepository struct {
	database   mongo.Database
	collection string
}


func NewReplyRepository(database mongo.Database, collection string) domain.ReplyRepository {
	return &replyRepository{
		database:   database,
		collection: collection}
}

func (r *replyRepository) CreateReply(ctx context.Context , reply domain.Reply) (domain.Reply, error) {
	
	collection  := r.database.Collection(r.collection)

	_, err := collection.InsertOne(ctx, reply)

	if err != nil {
		return domain.Reply{}, err
	}

	return reply, nil

}

func (ur *replyRepository) GetReply(ctx context.Context , blogID string) ([]domain.Reply, error) {
	
	// get all the replies for a blog
	collection := ur.database.Collection(ur.collection)

	cursor, err := collection.Find(ctx, bson.M{"blogid": blogID})
	fmt.Println(err , blogID)
	if err != nil {
		return []domain.Reply{}, err
	}

	var replies []domain.Reply

	if err = cursor.All(ctx, &replies); err != nil {
		return []domain.Reply{}, err
	}



	return replies, nil
	
}

