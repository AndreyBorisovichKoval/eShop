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

// MarkOrderAsPaid updates the order status to 'paid'
func MarkOrderAsPaid(orderID uint) error {
	// Get the order by ID
	order, err := repository.GetOrderByID(orderID)
	if err != nil {
		if err == errs.ErrRecordNotFound {
			logger.Warning.Printf("[service.MarkOrderAsPaid] order with ID [%d] not found\n", orderID)
			return errs.ErrOrderNotFound
		}
		return err
	}

	// Check if the order is already paid
	if order.IsPaid {
		logger.Warning.Printf("[service.MarkOrderAsPaid] order with ID [%d] is already paid\n", orderID)
		return errs.ErrOrderAlreadyPaid
	}

	// Mark the order as paid
	order.IsPaid = true

	// Update the order in the database
	err = repository.UpdateOrder(order)
	if err != nil {
		return err
	}

	logger.Info.Printf("Order with ID [%d] has been marked as paid\n", orderID)
	return nil
}

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

	// Обновляем общую сумму заказа
	order, err := repository.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	order.TotalAmount -= orderItem.Total
	err = repository.UpdateOrder(order)
	if err != nil {
		return err
	}

	logger.Info.Printf("Order item ID [%d] deleted from order ID [%d]\n", itemID, orderID)
	return nil
}

func DeleteOrder(orderID uint) error {
	logger.Info.Printf("[service.DeleteOrder] Attempting to delete order ID: %d", orderID)

	// Проверяем, существует ли заказ
	_, err := repository.GetOrderByID(orderID)
	if err != nil {
		if err == errs.ErrRecordNotFound {
			logger.Warning.Printf("[service.DeleteOrder] Order with ID [%d] not found", orderID)
			return errs.ErrOrderNotFound
		}
		return err
	}

	// Удаление всех товаров в заказе
	err = repository.DeleteOrderItemsByOrderID(orderID)
	if err != nil {
		logger.Error.Printf("[service.DeleteOrder] Error deleting order items for order ID [%d]: %v", orderID, err)
		return err
	}

	// Удаление самого заказа
	err = repository.DeleteOrder(orderID)
	if err != nil {
		logger.Error.Printf("[service.DeleteOrder] Error deleting order with ID [%d]: %v", orderID, err)
		return err
	}

	logger.Info.Printf("Order with ID [%d] and all its items have been deleted\n", orderID)
	return nil
}

// // GetOrderByID получает заказ по ID
// func GetOrderByID(orderID uint) (models.Order, error) {
// 	// Получаем заказ через репозиторий
// 	order, err := repository.GetOrderByID(orderID)
// 	if err != nil {
// 		if err == errs.ErrRecordNotFound {
// 			logger.Warning.Printf("[service.GetOrderByID] order with ID [%d] not found\n", orderID)
// 			return order, errs.ErrOrderNotFound
// 		}
// 		return order, err
// 	}

// 	// Получаем связанные товары (items) через репозиторий
// 	orderItems, err := repository.GetOrderItemsByOrderID(orderID)
// 	if err != nil {
// 		return order, err
// 	}

// 	// Присваиваем список товаров заказу
// 	order.OrderItems = orderItems

// 	return order, nil
// }

// GetOrderByID получает заказ по ID с нужными полями
func GetOrderByID(orderID uint) (map[string]interface{}, error) {
	// Получаем заказ через репозиторий
	order, err := repository.GetOrderByID(orderID)
	if err != nil {
		if err == errs.ErrRecordNotFound {
			logger.Warning.Printf("[service.GetOrderByID] order with ID [%d] not found\n", orderID)
			return nil, errs.ErrOrderNotFound
		}
		return nil, err
	}

	// Получаем связанные товары (items) через репозиторий
	orderItems, err := repository.GetOrderItemsByOrderID(orderID)
	if err != nil {
		return nil, err
	}

	// Составляем минимизированный ответ с нужными полями
	var items []map[string]interface{}
	for _, item := range orderItems {
		product, err := repository.GetProductByID(item.ProductID)
		if err != nil {
			return nil, err
		}
		items = append(items, map[string]interface{}{
			"product_title": product.Title,
			"quantity":      item.Quantity,
			"price":         item.Price,
			"total":         item.Total,
		})
	}

	result := map[string]interface{}{
		"order_id":     order.ID,
		"total_amount": order.TotalAmount,
		"order_items":  items,
	}

	return result, nil
}
