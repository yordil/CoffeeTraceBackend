package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID          primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	UserID      string             `json:"user_id"`
	Name 	  string             `json:"name"`
	Image 	 string             `json:"image"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Tags        []string           `json:"tags"`
	CreatedAt   time.Time             `json:"created_at"`
	// default False is resource value 
	IsResource 	bool             `json:"is_resource"`
	VideoURL 	string             `json:"video_url"`	
}

type Reply struct {
	ID    primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	UserID string `json:"user_id"`
	Name string `json:"name"`
	Image string `json:"image"`
	Reply string `json:"reply"`
	BlogID string `json:"blog_id"`
	CreatedAt time.Time `json:"created_at"`
}

type BlogRepository interface {
	CreateBlog(ctx context.Context , blog Blog) (Blog, error)
	GetAllBlog(ctx context.Context , bytag , pagesize , limit string) ([]Blog, error)
	GetBlogByID(ctx context.Context  , id string) (Blog, error)
	CreateResource(ctx context.Context, blog Blog)(SuccessResponse, ErrorResponse)
	GetResource(ctx context.Context)([]Blog, error)
	// CreateReply(ctx context.Context, reply Reply)(Reply, error)
	
}

type BlogUseCase interface {
	CreateBlog(blog Blog) (SuccessResponse, ErrorResponse)
	GetAllBlog(bytag , pagesize , limit string) ([]Blog, error)
	GetBlogByID(id string) (Blog, error)
	CreateReply(reply Reply)(Reply, error)
	GetReply(blogID string) ([]Reply, error)
	CreateResouce(blog Blog)(SuccessResponse, ErrorResponse)
	GetResource() ([]Blog, error)	
}

type ReplyRepository interface {
	CreateReply(ctx context.Context, reply Reply)(Reply, error)
	GetReply(ctx context.Context, blogID string)([]Reply, error)
}
