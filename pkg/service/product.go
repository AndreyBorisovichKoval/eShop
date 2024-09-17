// C:\GoProject\src\eShop\pkg\service\product.go

package service

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/repository"
	"eShop/utils"
	"errors"
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

// // AddProduct добавляет новый продукт и рассчитывает конечную цену с налогами
// func AddProduct(product models.Product) error {
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
// 		if errors.Is(err, errs.ErrUniquenessViolation) {
// 			logger.Warning.Printf("[service.AddProduct] duplicate barcode for product: %v\n", product.Barcode)
// 			return errs.ErrUniquenessViolation
// 		}
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
		barcode, err := utils.GenerateBarcode()
		if err != nil {
			logger.Error.Printf("[service.AddProduct] error generating barcode: %v\n", err)
			return err
		}
		product.Barcode = barcode
		logger.Info.Printf("[service.AddProduct] generated barcode for product: %s", barcode)
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
	product.TotalPrice = retailPrice * product.Quantity

	// Добавление продукта в базу данных
	if err := repository.CreateProduct(product); err != nil {
		logger.Error.Printf("[service.AddProduct] error creating product: %v\n", err)
		return err
	}

	logger.Info.Printf("Product %s successfully added with calculated retail price: %.2f", product.Title, retailPrice)
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
