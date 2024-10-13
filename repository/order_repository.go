package repository

import (
	"coeffee/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	database   mongo.Database
	collection string
}

func NewOrderRepository(database mongo.Database, collection string) domain.OrderRepository {
	return &OrderRepository{
		database:   database,
		collection: collection}

}

func (or *OrderRepository) CreateOrder(order domain.Order) (domain.Order, error) {
	_, err := or.database.Collection(or.collection).InsertOne(context.Background() , order)
	if err != nil {
		return domain.Order{}, err	
	}
	return order, nil
}

func (or *OrderRepository) GetMyOrders(userID, Role string) ([]domain.Order, error) {
	var orders []domain.Order
	if Role == "merchant" {
		cursor, err := or.database.Collection(or.collection).Find(context.Background(), bson.M{"merchant_id": userID})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var order domain.Order
			cursor.Decode(&order)
			orders = append(orders, order)
		}
	} else if Role == "farmer" {
		cursor, err := or.database.Collection(or.collection).Find(context.Background(), bson.M{"farmer_id": userID})
		if err != nil {
			return nil, err

		}
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var order domain.Order
			cursor.Decode(&order)
			orders = append(orders, order)
		}
	} else if Role == "driver" {
		cursor, err := or.database.Collection(or.collection).Find(context.Background(), bson.M{"driver_id": userID})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var order domain.Order
			cursor.Decode(&order)
			orders = append(orders, order)
		}
	}


	
	return orders, nil
}

func (or *OrderRepository) AcceptOrderFarmer(orderID, Role string) (domain.Order, error) {
	var order domain.Order

	objId , err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return domain.Order{}, err
	}
	err = or.database.Collection(or.collection).FindOne(context.Background(), bson.M{"_id": objId}).Decode(&order)
	if err != nil {
		return domain.Order{}, err
	}
	order.OrderStatusFarmer = "Accepted"
	update := bson.M{"$set": bson.M{"order_status_farmer": order.OrderStatusFarmer}}
	_, err = or.database.Collection(or.collection).UpdateOne(context.Background(), bson.M{"_id": objId}, update)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (or *OrderRepository) RejectOrderFarmer(orderID, Role string) (domain.Order, error) {
	var order domain.Order

	objId, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return domain.Order{}, err
	}

	err = or.database.Collection(or.collection).FindOne(context.Background(), bson.M{"_id": objId}).Decode(&order)
	if err != nil {
		return domain.Order{}, err
	}
	order.OrderStatusFarmer = "Rejected"
	update := bson.M{"$set": bson.M{"order_status_farmer": order.OrderStatusFarmer}}
	_, err = or.database.Collection(or.collection).UpdateOne(context.Background(), bson.M{"_id": objId}, update)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}



func (or *OrderRepository) AcceptOrderDriver(orderID, Role string) (domain.Order, error) {
	var order domain.Order

	objId , err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return domain.Order{}, err
	}
	err = or.database.Collection(or.collection).FindOne(context.Background(), bson.M{"_id": objId}).Decode(&order)
	if err != nil {
		return domain.Order{}, err
	}
	order.DriverStatus = "Accepted"
	update := bson.M{"$set": bson.M{"driver_status": order.DriverStatus}}
	_, err = or.database.Collection(or.collection).UpdateOne(context.Background(), bson.M{"_id": objId}, update)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}


func (or *OrderRepository) RejectOrderDriver(orderID, Role string) (domain.Order, error) {
	var order domain.Order

	objId , err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return domain.Order{}, err
	}
	err = or.database.Collection(or.collection).FindOne(context.Background(), bson.M{"_id": objId}).Decode(&order)
	if err != nil {
		return domain.Order{}, err
	}
	order.DriverStatus = "Rejected"
	update := bson.M{"$set": bson.M{"driver_status": order.DriverStatus}}
	_, err = or.database.Collection(or.collection).UpdateOne(context.Background(), bson.M{"_id": objId}, update)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}


func (or *OrderRepository) PickupOrderDriver(orderID, Role string) (domain.Order, error) {
	var order domain.Order

	objId , err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return domain.Order{}, err
	}
	err = or.database.Collection(or.collection).FindOne(context.Background(), bson.M{"_id": objId}).Decode(&order)
	if err != nil {
		return domain.Order{}, err
	}
	order.ShippingStatus = "Picked"
	update := bson.M{"$set": bson.M{"shipping_status": order.ShippingStatus}}
	_, err = or.database.Collection(or.collection).UpdateOne(context.Background(), bson.M{"_id": objId}, update)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}


func (or *OrderRepository) PayFarmer(orderID, Role string) (domain.Order, error) { 
	var order domain.Order

	objId , err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return domain.Order{}, err
	}
	err = or.database.Collection(or.collection).FindOne(context.Background(), bson.M{"_id": objId}).Decode(&order)
	if err != nil {
		return domain.Order{}, err
	}
	order.FarmerPaymentStatus = "Paid"
	update := bson.M{"$set": bson.M{"farmer_payment_status": order.FarmerPaymentStatus}}
	_, err = or.database.Collection(or.collection).UpdateOne(context.Background(), bson.M{"_id": objId}, update)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}


func (or *OrderRepository) PayDriver(orderID, Role string) (domain.Order, error) { 
	var order domain.Order

	objId , err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return domain.Order{}, err
	}
	err = or.database.Collection(or.collection).FindOne(context.Background(), bson.M{"_id": objId}).Decode(&order)
	if err != nil {
		return domain.Order{}, err
	}
	order.DriverPaymentStatus = "Paid"
	update := bson.M{"$set": bson.M{"driver_payment_status": order.DriverPaymentStatus}}
	_, err = or.database.Collection(or.collection).UpdateOne(context.Background(), bson.M{"_id": objId}, update)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}


func (or *OrderRepository) UpdateDestinations(orderID, Role string, Destinations domain.Destinations) (domain.Order, error) {
	var order domain.Order

	objId , err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return domain.Order{}, err
	}

	err = or.database.Collection(or.collection).FindOne(context.Background(), bson.M{"_id": objId}).Decode(&order)
	if err != nil {

		return domain.Order{}, err
	}
	order.DestinationsLocation = append(order.DestinationsLocation, Destinations)
	update := bson.M{"$set": bson.M{"destinations_location": order.DestinationsLocation}}
	_, err = or.database.Collection(or.collection).UpdateOne(context.Background(), bson.M{"_id": objId}, update)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

