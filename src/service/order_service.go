package service

import (
	"github.com/google/uuid"
	"github.com/lalatina11/markita.git/src/config"
	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/model"
	"gorm.io/gorm"
)

type OrderService struct {
	Db *gorm.DB
}

func NewOrderService() *OrderService {
	return &OrderService{Db: config.NewDatabaseConfig().Connect()}
}

func (this *OrderService) Direct(order *model.Order, userID string) (*model.Order, *service_error.ServiceError) {
	orderId := uuid.NewString()
	order.ID = orderId
	order.UserID = userID

	for i := range order.Items {
		order.Items[i].ID = uuid.NewString()
		order.Items[i].OrderID = orderId
		order.Items[i].Status = "pending"
	}

	err := this.Db.Create(&order).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	return order, nil
}

func (this *OrderService) GetAllOrders(userID string) ([]model.Order, *service_error.ServiceError) {
	var orders []model.Order

	err := this.Db.Where(&model.Order{UserID: userID}).Find(&orders).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	return orders, nil
}
