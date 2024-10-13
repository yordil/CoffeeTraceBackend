package domain

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



type LocationType struct {
	Longtiude float64 `bson:"longtiude" json:"longtiude"`
	Latitude float64 `bson:"latitude" json:"latitude"`
	Description string `bson:"description" json:"description"`
}

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Name string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Role string `bson:"role" json:"role"`
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
	Address string `bson:"address" json:"address"`
	CreatedAt	time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time`bson:"updated_at" json:"updated_at"`

	Image string `bson:"image" json:"image"`

	// Notifications
	Notifications []Notification `bson:"notifications" json:"notifications"`
	OrderConfirmation bool `bson:"order_confirmation" json:"order_confirmation"`
	OrderStatusConfirmation bool `bson:"order_status" json:"order_status"`
	EmailNotification bool `bson:"email_notification" json:"email_notification"`


	// Transporter
	TruckNumber string `bson:"truck_number" json:"truck_number"`
	TruckType string `bson:"truck_type" json:"truck_type"`
	TrailerNumber string `bson:"trailer_number" json:"trailer_number"`
	StartLocation string `bson:"start_location" json:"start_location"`
	EndLocation string `bson:"end_location" json:"end_location"`


	



	// comunity
	Followers []string `bson:"followers" json:"followers"`
	Following []string `bson:"following" json:"following"`
	PostsId []string `bson:"posts" json:"posts"`	
}



type TokenGenerator interface {
	GenerateToken(user User) (string, error)
	GenerateRefreshToken(user User) (string, error)
	RefreshToken(token string) (string, error)
}
type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
type JwtCustomClaims struct {
	UserID string `json:"user_id"`
	Role string `json:"role"`
	Name string `json:"name"`

	jwt.StandardClaims
}



type UserUseCase interface {

	CreateAccount(user User) (User, ErrorResponse)
	Login(user User) (LoginSucessResponse, ErrorResponse)
	GetByID(id string) (User, ErrorResponse)

	UpdateProfile(id string, user User) (User, ErrorResponse)
	AddDriver(user User) (User, ErrorResponse)


	GetAllUser() ([]User, error)

	GetUserByID(id string) (User, error)
}


type UserRepository interface {

	CreateAccount(user User) (User, error)
	Login(user User) (User, error)
	GetAllUserByEmial(email string) (User, error)
	GetByID(id string) (User, error)
	UpdateProfile(id string, user User) (User, error)
	AddDriver(user User) (User, error)

	GetAllUser() ([]User, error)



	




	// GetAllUser(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context ,id string) (User, error)
	// CreateUser(ctx context.Context , user User) (error)
	// // UpdateUser(ctx context.Context ,id string, user User) (User, error)
	// // DeleteUser(ctx context.Context ,id string) (User, error)
	// // Login(ctx context.Context ,user UserLogin) (User, error)
	// FindUserByEmail(ctx context.Context, email string) (User, error)
}



type LocationUpdateRequest struct {
	// Longtiude float64 `json:"longtiude"`
	// Latitude float64 `json:"latitude"`
	// Description string `json:"description"`
	// Color string `json:"color"`
	// Time time.Time `json:"time"`
}