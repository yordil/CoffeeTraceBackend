package usecase

import (
	"coeffee/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderUseCase struct {
	OrderRepository  domain.OrderRepository
	ProductRepository domain.ProductRepository
	NofiticationRepository domain.NotificationRepository
	UserRepository domain.UserRepository
	TransactionsRepo domain.TransactionsRepository
	contextTimeout    time.Duration

}

func NewOrderUseCase(order domain.OrderRepository, notify domain.NotificationRepository ,  userrepo domain.UserRepository , productrepo domain.ProductRepository, transactionrepo  domain.TransactionsRepository , timeout time.Duration) domain.OrderUseCase {
	return &OrderUseCase{
		OrderRepository: order,
		ProductRepository: productrepo,
		UserRepository: userrepo,
		NofiticationRepository: notify,
		TransactionsRepo: transactionrepo,
		contextTimeout: timeout,
	}

}

func (uc *OrderUseCase) CreateOrder(order domain.Order) (domain.Order, error) {
	order.ID = primitive.NewObjectID()

	order.OrderDate = time.Now()
	order.OrderStatus = "Pending"
	order.ShippingStatus = "Pending"
	order.DriverStatus = "Pending"
	order.OrderStatusFarmer = "Pending"
	order.FarmerPaymentStatus = "Pending"
	order.DriverPaymentStatus = "Pending"

	_, err := uc.OrderRepository.CreateOrder(order)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (uc *OrderUseCase) GetMyOrders(userID, Role string) ([]domain.Order, error) {
	orders, err := uc.OrderRepository.GetMyOrders(userID, Role)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (uc *OrderUseCase) AcceptOrderFarmer(orderID, Role string) (domain.Order, error) {
	order, err := uc.OrderRepository.AcceptOrderFarmer(orderID, Role)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (uc *OrderUseCase) RejectOrderFarmer(orderID, Role string) (domain.Order, error) {
	order, err := uc.OrderRepository.RejectOrderFarmer(orderID, Role)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}


func (uc *OrderUseCase) AcceptOrderDriver(orderID, Role string) (domain.Order, error) {
	order, err := uc.OrderRepository.AcceptOrderDriver(orderID, Role)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (uc *OrderUseCase) RejectOrderDriver(orderID, Role string) (domain.Order, error) {
	order, err := uc.OrderRepository.RejectOrderDriver(orderID, Role)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (uc *OrderUseCase) PickupOrderDriver(orderID, Role string) (domain.Order, error) { 
	order, err := uc.OrderRepository.PickupOrderDriver(orderID, Role)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (uc *OrderUseCase) PayFarmer(orderID, Role string) (domain.Order, error) {
	order, err := uc.OrderRepository.PayFarmer(orderID, Role)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}


func (uc *OrderUseCase) PayDriver(orderID, Role string) (domain.Order, error) {
	order, err := uc.OrderRepository.PayDriver(orderID, Role)
	if err != nil {
		return domain.Order{}, err
	}

	newTransaction := domain.Transactions{
		OrderID: orderID,
		Amount: order.ShippingCost + order.TotalPrice,
		Description: "Payment for order",
		FarmerID: order.FarmerID,
		FarmerName: order.FarmerName,
		MerchantID: order.MerchantID,
		MerchantName: order.MerchantName,
		Quantity: order.Quantity,
		TotalPrice: order.TotalPrice,
		Datetime: time.Now().String(),
	}

	_, err = uc.TransactionsRepo.CreateTransaction(newTransaction)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (uc *OrderUseCase) UpdateDestinations(orderID, Role string, Destinations domain.Destinations) (domain.Order, error) {
	order, err := uc.OrderRepository.UpdateDestinations(orderID, Role, Destinations)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
