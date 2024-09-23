// C:\GoProject\src\eShop\pkg\repository\barcode.go

package repository

import (
	"eShop/db"
	"eShop/errs"
	"eShop/logger"
	"eShop/models"

	"gorm.io/gorm"
)

func GetProductForBarcodeByID(productID uint) (models.Product, error) {
	var product models.Product

	err := db.GetDBConn().First(&product, productID).Error
	if err != nil {
		logger.Error.Printf("Error fetching product by ID %d: %v", productID, err)
		return product, errs.ErrProductNotFound
	}

	return product, nil
}

func FindProductByID(productID uint) (models.Product, error) {
	var product models.Product
	err := db.GetDBConn().First(&product, productID).Error
	if err != nil {
		logger.Error.Printf("Error fetching product by ID %d: %v", productID, err)
		return product, errs.ErrProductNotFound
	}
	return product, nil
}

func CheckOrderExists(orderID uint) (bool, error) {
	var order models.Order
	err := db.GetDBConn().First(&order, orderID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		logger.Error.Printf("Error checking order existence: %v", err)
		return false, err
	}
	return true, nil
}

func InsertProductIntoOrder(orderID uint, productID uint, quantity float64, price float64) error {
	// Добавляем товар в таблицу order_items
	orderItem := models.OrderItem{
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
		Total:     price * quantity,
	}

	err := db.GetDBConn().Create(&orderItem).Error
	if err != nil {
		logger.Error.Printf("Error inserting product into order: %v", err)
		return err
	}

	// Обновляем сумму заказа
	err = UpdateOrderTotal(orderID, orderItem.Total)
	if err != nil {
		logger.Error.Printf("Error updating order total: %v", err)
		return err
	}

	return nil
}

// UpdateOrderTotal обновляет общую сумму заказа
func UpdateOrderTotal(orderID uint, amountToAdd float64) error {
	// Увеличиваем сумму заказа на сумму добавленного товара
	err := db.GetDBConn().Model(&models.Order{}).Where("id = ?", orderID).
		UpdateColumn("total_amount", gorm.Expr("total_amount + ?", amountToAdd)).Error
	if err != nil {
		logger.Error.Printf("Error updating order total for order ID %d: %v", orderID, err)
		return err
	}

	return nil
}
