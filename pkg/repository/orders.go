// C:\GoProject\src\eShop\pkg\repository\orders.go

package repository

import (
	"eShop/db"
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
)

// CreateOrder создает новый заказ в базе данных
func CreateOrder(order *models.Order) error {
	if err := db.GetDBConn().Create(order).Error; err != nil {
		logger.Error.Printf("[repository.CreateOrder] error creating order: %v\n", err)
		return translateError(err)
	}
	return nil
}

// CreateOrderItem создает новый товар в заказе
func CreateOrderItem(orderItem *models.OrderItem) error {
	if err := db.GetDBConn().Create(orderItem).Error; err != nil {
		logger.Error.Printf("[repository.CreateOrderItem] error creating order item: %v\n", err)
		return translateError(err)
	}
	return nil
}

// UpdateOrder обновляет данные заказа в базе данных
func UpdateOrder(order models.Order) error {
	if err := db.GetDBConn().Save(&order).Error; err != nil {
		logger.Error.Printf("[repository.UpdateOrder] error updating order: %v\n", err)
		return translateError(err)
	}
	return nil
}

// GetOrderItemByID получает товар в заказе по ID заказа и товара
func GetOrderItemByID(orderID, itemID uint) (models.OrderItem, error) {
	var orderItem models.OrderItem
	err := db.GetDBConn().Where("order_id = ? AND id = ?", orderID, itemID).First(&orderItem).Error
	if err != nil {
		if err.Error() == "record not found" {
			return orderItem, errs.ErrRecordNotFound
		}
		return orderItem, translateError(err)
	}
	return orderItem, nil
}

// DeleteOrderItem удаляет товар из заказа
func DeleteOrderItem(orderItem models.OrderItem) error {
	if err := db.GetDBConn().Delete(&orderItem).Error; err != nil {
		logger.Error.Printf("[repository.DeleteOrderItem] error deleting order item: %v\n", err)
		return translateError(err)
	}
	return nil
}
