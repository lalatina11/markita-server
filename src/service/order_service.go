package service

import (
	"github.com/google/uuid"
	"github.com/lalatina11/markita.git/src/dto/order_dto"
	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/model"
	"gorm.io/gorm"
)

type OrderService struct {
	Db *gorm.DB
}

func NewOrderService(Db *gorm.DB) *OrderService {
	return &OrderService{Db}
}

func (this *OrderService) Find(id string) (*order_dto.OrderDTO, *service_error.ServiceError) {
	order := new(model.Order)
	order.ID = id

	err := this.Db.
		Preload("User").
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Store").
		Preload("Items.Product.Media").
		First(order).Error

	if err != nil {
		return nil, service_error.NotFound()
	}

	return order.ToOrderDTO(), nil
}

func (this *OrderService) Direct(order *model.Order, userID string) (*order_dto.OrderDTO, *service_error.ServiceError) {
	orderId := uuid.NewString()
	order.ID = orderId
	order.UserID = userID

	for i := range order.Items {
		order.Items[i].ID = uuid.NewString()
		order.Items[i].OrderID = orderId
		order.Items[i].Status = "pending"
	}

	err := this.Db.Create(order).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	return this.Find(order.ID)
}

func (this *OrderService) GetAllOrders(userID string) ([]order_dto.OrderDTO, *service_error.ServiceError) {
	var orders []model.Order

	err := this.Db.
		Preload("User").
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Store").
		Preload("Items.Product.Media").
		Where(&model.Order{UserID: userID}).
		Find(&orders).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	var dtos = make([]order_dto.OrderDTO, len(orders))
	for i, o := range orders {
		dtos[i] = *o.ToOrderDTO()
	}

	return dtos, nil
}
