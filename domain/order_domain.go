package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Destinations struct {
	Time string `json:"time" bson:"time"`
	Content string `json:"content" bson:"content"`
	Color string `json:"color" bson:"color"`
	Latitude float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

type Order struct {
	ID primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	ProductID   string `json:"product_id" bson:"product_id"`
	MerchantID  string `json:"merchant_id" bson:"merchant_id"`
	MerchantName string `json:"merchant_name" bson:"merchant_name"`
	FarmerID   string `json:"farmer_id" bson:"farmer_id"`
	FarmerName string `json:"farmer_name" bson:"farmer_name"`



	Quantity    int    `json:"quantity" bson:"quantity"`
	TotalPrice  int    `json:"total_price" bson:"total_price"`
	OrderStatus string `json:"order_status" bson:"order_status"`
	OrderDate   time.Time `json:"order_date" bson:"order_date"`

	OrderStatusFarmer string `json:"order_status_farmer" bson:"order_status_farmer"`
	FarmerPaymentStatus string `json:"farmer_payment_status" bson:"farmer_payment_status"`

	OrderType string `json:"order_type" bson:"order_type"`
	DriverID string `json:"driver_id" bson:"driver_id"`
	DriverStatus string `json:"driver_status" bson:"driver_status"`
	StartLocation string `json:"start_location" bson:"start_location"`
	DestinationLocation string `json:"end_location" bson:"end_location"`

	DriverPaymentStatus string `json:"driver_payment_status" bson:"driver_payment_status"`



	ShippingCoverage string `json:"shipping_coverage" bson:"shipping_coverage"`
	ShippingCost int `json:"shipping_cost" bson:"shipping_cost"`
	ShippingDate time.Time `json:"shipping_date" bson:"shipping_date"`
	ShippingStatus string `json:"shipping_status" bson:"shipping_status"`

	DestinationsLocation []Destinations `json:"destinations_location" bson:"destinations_location"`

}


	

type OrderRepository interface {
	CreateOrder (order Order) (Order, error)
	GetMyOrders (userID, Role string) ([]Order, error)


	// farmer
	AcceptOrderFarmer (orderID, Role string) (Order, error)
	RejectOrderFarmer (orderID, Role string) (Order, error)


	// driver
	AcceptOrderDriver (orderID, Role string) (Order, error)
	RejectOrderDriver (orderID, Role string) (Order, error)
	PickupOrderDriver (orderID, Role string) (Order, error)
	UpdateDestinations (orderID, Role string, Destinations Destinations) (Order, error)


	// buyer
	PayFarmer (orderID, Role string) (Order, error)
	PayDriver (orderID, Role string) (Order, error)
}

type OrderUseCase interface {
	CreateOrder(order Order) (Order, error)
	GetMyOrders(userID , Role string) ([]Order, error)


	// farmer
	AcceptOrderFarmer(orderID, Role string) (Order, error)
	RejectOrderFarmer(orderID, Role string) (Order, error)

	// driver
	AcceptOrderDriver(orderID, Role string) (Order, error)
	RejectOrderDriver(orderID, Role string) (Order, error)
	PickupOrderDriver(orderID, Role string) (Order, error)
	UpdateDestinations(orderID, Role string, Destinations Destinations) (Order, error)

	// buyer
	PayFarmer(orderID, Role string) (Order, error)
	PayDriver(orderID, Role string) (Order, error)
}
