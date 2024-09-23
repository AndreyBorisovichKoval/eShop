package service

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/repository"
	"math/rand"
	"time"
)

// GenerateRandomOrders генерирует случайные заказы для тестирования
func GenerateRandomOrders(count int) error {
	// Получаем всех доступных пользователей (продавцов)
	users, err := repository.GetAllUsers()
	if err != nil {
		return err
	}

	// Получаем все доступные товары
	products, err := repository.GetAllProducts()
	if err != nil {
		return err
	}

	// Если нет продавцов или товаров, невозможно создать заказы
	if len(users) == 0 || len(products) == 0 {
		return errs.ErrSomethingWentWrong
	}

	// Генерация случайных заказов
	for i := 0; i < count; i++ {
		// Выбираем случайного пользователя (продавца)
		user := users[rand.Intn(len(users))]

		// Создаем новый заказ
		order := models.Order{
			UserID:    user.ID,
			IsPaid:    true, // По умолчанию заказы считаем оплаченными
			CreatedAt: time.Now(),
		}

		// Добавляем заказ в базу данных
		err := repository.CreateOrder(&order)
		if err != nil {
			logger.Error.Printf("[service.GenerateRandomOrders] error creating order: %v", err)
			return err
		}

		// Генерируем случайное количество товаров в заказе (от 1 до 6)
		productCount := rand.Intn(6) + 1
		var totalAmount float64

		for j := 0; j < productCount; j++ {
			// Выбираем случайный товар
			product := products[rand.Intn(len(products))]

			// Создаем элемент заказа
			orderItem := models.OrderItem{
				OrderID:   order.ID,
				ProductID: product.ID,
				Quantity:  float64(rand.Intn(6) + 1), // Случайное количество от 1 до 6
				Price:     product.RetailPrice,
				Total:     product.RetailPrice * float64(rand.Intn(6)+1),
			}

			// Добавляем элемент заказа в базу данных
			err := repository.CreateOrderItem(&orderItem)
			if err != nil {
				logger.Error.Printf("[service.GenerateRandomOrders] error creating order item: %v", err)
				return err
			}

			totalAmount += orderItem.Total
		}

		// Обновляем общую сумму заказа
		order.TotalAmount = totalAmount
		err = repository.UpdateOrder(order)
		if err != nil {
			return err
		}
	}

	return nil
}
