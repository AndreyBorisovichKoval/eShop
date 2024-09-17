// C:\GoProject\src\eShop\pkg\service\orders.go

package service

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/repository"
	"time"
)

// CreateOrder создает заказ и добавляет товары в заказ
func CreateOrder(userID uint, orderItems []models.OrderItem) (models.Order, error) {
	// Создаем новый заказ
	order := models.Order{
		UserID:    userID,
		IsPaid:    false,
		CreatedAt: time.Now(),
	}

	// Вставляем заказ в базу
	err := repository.CreateOrder(&order)
	if err != nil {
		logger.Error.Printf("[service.CreateOrder] error creating order: %v\n", err)
		return order, err
	}

	var totalAmount float64

	// Обрабатываем каждый товар в заказе
	for i, item := range orderItems {
		product, err := repository.GetProductByID(item.ProductID)
		if err != nil {
			if err == errs.ErrRecordNotFound {
				logger.Warning.Printf("[service.CreateOrder] product with ID [%d] not found\n", item.ProductID)
				return order, errs.ErrProductNotFound
			}
			return order, err
		}

		// Проверяем количество на складе
		if product.Stock < float64(item.Quantity) {
			logger.Warning.Printf("[service.CreateOrder] insufficient stock for product ID [%d]\n", item.ProductID)
			return order, errs.ErrInsufficientStock
		}

		// Обновляем цену на момент заказа
		item.OrderID = order.ID
		item.Price = product.RetailPrice
		item.Total = product.RetailPrice * float64(item.Quantity)

		// Вычитаем количество товара из склада
		product.Stock += float64(item.Quantity)

		// Обновляем данные продукта в БД
		err = repository.UpdateProduct(product)
		if err != nil {
			return order, err
		}

		// Вставляем OrderItem в базу
		err = repository.CreateOrderItem(&item)
		if err != nil {
			return order, err
		}

		// Считаем общую сумму заказа
		totalAmount += item.Total
		orderItems[i] = item
	}

	// Обновляем общую сумму заказа
	order.TotalAmount = totalAmount

	// Обновляем заказ в базе
	err = repository.UpdateOrder(order)
	if err != nil {
		return order, err
	}

	return order, nil
}

// DeleteOrderItem удаляет товар из заказа
func DeleteOrderItem(orderID, itemID uint) error {
	orderItem, err := repository.GetOrderItemByID(orderID, itemID)
	if err != nil {
		if err == errs.ErrRecordNotFound {
			return errs.ErrOrderItemNotFound
		}
		return err
	}

	// Возвращаем товар на склад
	product, err := repository.GetProductByID(orderItem.ProductID)
	if err != nil {
		return err
	}

	product.Stock += float64(orderItem.Quantity)

	// Обновляем количество товара на складе
	err = repository.UpdateProduct(product)
	if err != nil {
		return err
	}

	// Удаляем товар из заказа
	err = repository.DeleteOrderItem(orderItem)
	if err != nil {
		return err
	}

	return nil
}
