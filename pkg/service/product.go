// C:\GoProject\src\eShop\pkg\service\product.go

package service

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/repository"
	"eShop/utils"
	"errors"
	"time"
)

// GetAllProducts получает все активные продукты
func GetAllProducts() ([]models.Product, error) {
	products, err := repository.GetAllProducts()
	if err != nil {
		logger.Error.Printf("[service.GetAllProducts] error retrieving products: %v", err)
		return nil, err
	}

	return products, nil
}

// GetProductByID получает продукт по ID
func GetProductByID(id uint) (models.Product, error) {
	product, err := repository.GetProductByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return product, errs.ErrProductNotFound
		}
		logger.Error.Printf("[service.GetProductByID] error fetching product by id: %v\n", err)
		return product, err
	}

	return product, nil
}

// GetProductByBarcode получает продукт по штрих-коду
func GetProductByBarcode(barcode string) (models.Product, error) {
	product, err := repository.GetProductByBarcode(barcode)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return product, errs.ErrProductNotFound
		}
		logger.Error.Printf("[service.GetProductByBarcode] error fetching product by barcode: %v\n", err)
		return product, err
	}

	return product, nil
}

// UpdateProductByID обновляет данные продукта по его ID
func UpdateProductByID(id uint, updatedProduct models.Product) (models.Product, error) {
	// Получаем существующий продукт по ID
	product, err := repository.GetProductByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return product, errs.ErrProductNotFound
		}
		logger.Error.Printf("[service.UpdateProductByID] error retrieving product by id: %v\n", err)
		return product, err
	}

	// Обновляем только изменённые поля
	if updatedProduct.Title != "" {
		product.Title = updatedProduct.Title
	}
	if updatedProduct.Quantity != 0 {
		product.Quantity = updatedProduct.Quantity
	}
	if updatedProduct.SupplierPrice != 0 {
		product.SupplierPrice = updatedProduct.SupplierPrice
	}
	if updatedProduct.CategoryID != 0 {
		product.CategoryID = updatedProduct.CategoryID
	}
	if updatedProduct.SupplierID != 0 {
		product.SupplierID = updatedProduct.SupplierID
	}
	if updatedProduct.Discount != 0 {
		product.Discount = updatedProduct.Discount
	}
	if updatedProduct.Unit != "" {
		product.Unit = updatedProduct.Unit
	}
	if updatedProduct.StorageLocation != "" {
		product.StorageLocation = updatedProduct.StorageLocation
	}
	if updatedProduct.ExpirationDate != nil {
		product.ExpirationDate = updatedProduct.ExpirationDate
	}

	// Пересчёт цены с учётом налогов
	taxes, err := repository.GetAllTaxes()
	if err != nil {
		logger.Error.Printf("[service.UpdateProductByID] error fetching taxes: %v\n", err)
		return product, err
	}

	product.RetailPrice = calculateRetailPrice(product.SupplierPrice, taxes, product.IsVATApplicable, product.IsExciseApplicable)
	product.TotalPrice = product.RetailPrice * product.Quantity

	// Сохраняем изменения
	if err := repository.UpdateProduct(product); err != nil {
		logger.Error.Printf("[service.UpdateProductByID] error updating product: %v\n", err)
		return product, err
	}

	return product, nil
}

// SoftDeleteProductByID мягко удаляет продукт
func SoftDeleteProductByID(id uint) error {
	// Получаем продукт, включая мягко удалённые записи
	product, err := repository.GetProductIncludingSoftDeleted(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrProductNotFound
		}
		return err
	}

	// Проверяем, был ли продукт уже удалён
	if product.IsDeleted {
		return errs.ErrProductAlreadyDeleted
	}

	// Устанавливаем флаг удаления и время
	product.IsDeleted = true
	now := time.Now()
	product.DeletedAt = &now

	// Сохраняем изменения в базе данных
	if err := repository.UpdateProduct(product); err != nil {
		return err
	}

	return nil
}

// RestoreProductByID восстанавливает продукт
func RestoreProductByID(id uint) error {
	// Получаем продукт, включая мягко удалённые записи
	product, err := repository.GetProductIncludingSoftDeleted(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrProductNotFound
		}
		return err
	}

	// Проверяем, не был ли продукт уже восстановлен
	if !product.IsDeleted {
		return errs.ErrProductNotDeleted
	}

	// Сбрасываем флаг удаления и удаляем дату удаления
	product.IsDeleted = false
	product.DeletedAt = nil

	// Сохраняем изменения
	if err := repository.UpdateProduct(product); err != nil {
		return err
	}

	return nil
}

// HardDeleteProductByID выполняет жесткое удаление продукта
func HardDeleteProductByID(id uint) error {
	// Получаем продукт по ID, включая мягко удалённые записи
	product, err := repository.GetProductIncludingSoftDeleted(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			logger.Warning.Printf("[service.HardDeleteProductByID] product with ID: %v not found", id)
			return errs.ErrProductNotFound
		}
		return err
	}

	// Выполняем жёсткое удаление продукта
	if err := repository.HardDeleteProductByID(product); err != nil {
		return err
	}

	logger.Info.Printf("[service.HardDeleteProductByID] product with ID %v hard deleted successfully", id)
	return nil
}

// Функция для расчета розничной цены
func calculateRetailPrice(supplierPrice float64, taxes []models.Taxes, isVAT, isExcise bool) float64 {
	retailPrice := supplierPrice

	for _, tax := range taxes {
		if tax.ApplyTo == "final_price" && isVAT && tax.Title == "VAT" {
			retailPrice += supplierPrice * (tax.Rate / 100)
		}

		if tax.ApplyTo == "profit" && isExcise && tax.Title == "Excise" {
			profit := retailPrice - supplierPrice
			retailPrice += profit * (tax.Rate / 100)
		}
	}

	return retailPrice
}

// // AddProduct добавляет новый продукт и рассчитывает конечную цену с налогами
// func AddProduct(product models.Product) error {
// 	// Если штрих-код не указан, генерируем его
// 	if product.Barcode == "" {
// 		for {
// 			barcode, err := utils.GenerateBarcode()
// 			if err != nil {
// 				logger.Error.Printf("[service.AddProduct] error generating barcode: %v\n", err)
// 				return err
// 			}

// 			exists, err := repository.CheckBarcodeExists(barcode)
// 			if err != nil {
// 				logger.Error.Printf("[service.AddProduct] error checking barcode existence: %v\n", err)
// 				return err
// 			}

// 			// Если штрих-код уникален, сохраняем его и выходим из цикла
// 			if !exists {
// 				product.Barcode = barcode
// 				logger.Info.Printf("[service.AddProduct] generated unique barcode for product: %s", barcode)
// 				break
// 			}
// 		}
// 	}

// 	// Проверка на уникальность штрих-кода
// 	exists, err := repository.CheckBarcodeExists(product.Barcode)
// 	if err != nil {
// 		logger.Error.Printf("[service.AddProduct] error checking barcode existence: %v\n", err)
// 		return err
// 	}
// 	if exists {
// 		logger.Warning.Printf("[service.AddProduct] duplicate barcode for product: %v\n", product.Barcode)
// 		return errs.ErrProductAlreadyExists
// 	}

// 	// Проверка поставщика
// 	supplier, err := repository.GetSupplierByID(product.SupplierID)
// 	if err != nil {
// 		if errors.Is(err, errs.ErrRecordNotFound) {
// 			return errs.ErrSupplierNotFound
// 		}
// 		logger.Error.Printf("[service.AddProduct] error fetching supplier by id: %v\n", err)
// 		return err
// 	}
// 	product.Supplier = supplier

// 	// Проверка категории
// 	category, err := repository.GetCategoryByID(product.CategoryID)
// 	if err != nil {
// 		if errors.Is(err, errs.ErrRecordNotFound) {
// 			return errs.ErrCategoryNotFound
// 		}
// 		logger.Error.Printf("[service.AddProduct] error fetching category by id: %v\n", err)
// 		return err
// 	}
// 	product.Category = category

// 	// Получаем все налоги
// 	taxes, err := repository.GetAllTaxes()
// 	if err != nil {
// 		logger.Error.Printf("[service.AddProduct] error fetching taxes: %v\n", err)
// 		return err
// 	}

// 	// Расчет розничной цены
// 	retailPrice := calculateRetailPrice(product.SupplierPrice, taxes, product.IsVATApplicable, product.IsExciseApplicable)

// 	product.RetailPrice = retailPrice
// 	product.TotalPrice = retailPrice * product.Quantity

// 	// Добавление продукта в базу данных
// 	if err := repository.CreateProduct(product); err != nil {
// 		logger.Error.Printf("[service.AddProduct] error creating product: %v\n", err)
// 		return err
// 	}

// 	logger.Info.Printf("Product %s successfully added with calculated retail price: %.2f", product.Title, retailPrice)
// 	return nil
// }

// AddProduct добавляет новый продукт и рассчитывает конечную цену с налогами
func AddProduct(product models.Product) error {
	// Если штрих-код не указан, генерируем его
	if product.Barcode == "" {
		for {
			barcode, err := utils.GenerateBarcode()
			if err != nil {
				logger.Error.Printf("[service.AddProduct] error generating barcode: %v\n", err)
				return err
			}

			exists, err := repository.CheckBarcodeExists(barcode)
			if err != nil {
				logger.Error.Printf("[service.AddProduct] error checking barcode existence: %v\n", err)
				return err
			}

			if !exists {
				product.Barcode = barcode
				logger.Info.Printf("[service.AddProduct] generated unique barcode for product: %s", barcode)
				break
			}
		}
	}

	// Проверка на уникальность штрих-кода
	exists, err := repository.CheckBarcodeExists(product.Barcode)
	if err != nil {
		logger.Error.Printf("[service.AddProduct] error checking barcode existence: %v\n", err)
		return err
	}
	if exists {
		logger.Warning.Printf("[service.AddProduct] duplicate barcode for product: %v\n", product.Barcode)
		return errs.ErrProductAlreadyExists
	}

	// Проверка поставщика
	supplier, err := repository.GetSupplierByID(product.SupplierID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrSupplierNotFound
		}
		logger.Error.Printf("[service.AddProduct] error fetching supplier by id: %v\n", err)
		return err
	}
	product.Supplier = supplier

	// Проверка категории
	category, err := repository.GetCategoryByID(product.CategoryID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrCategoryNotFound
		}
		logger.Error.Printf("[service.AddProduct] error fetching category by id: %v\n", err)
		return err
	}
	product.Category = category

	// Получаем все налоги
	taxes, err := repository.GetAllTaxes()
	if err != nil {
		logger.Error.Printf("[service.AddProduct] error fetching taxes: %v\n", err)
		return err
	}

	// Расчет розничной цены
	retailPrice := calculateRetailPrice(product.SupplierPrice, taxes, product.IsVATApplicable, product.IsExciseApplicable)

	product.RetailPrice = retailPrice
	// product.TotalPrice = retailPrice * product.Quantity
	product.TotalPrice = product.SupplierPrice * product.Quantity
	product.Stock += product.Quantity // Обновляем поле stock

	// Добавление продукта в базу данных
	if err := repository.CreateProduct(product); err != nil {
		logger.Error.Printf("[service.AddProduct] error creating product: %v\n", err)
		return err
	}

	logger.Info.Printf("Product %s successfully added with calculated retail price: %.2f", product.Title, retailPrice)
	return nil
}
