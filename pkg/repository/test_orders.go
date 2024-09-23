package repository

// import (
// 	"eShop/db"
// 	"eShop/logger"
// 	"eShop/models"
// )

// // GetAllUsersTest возвращает всех пользователей
// func GetAllUsersTest() ([]models.User, error) {
// 	var users []models.User
// 	err := db.GetDBConn().Find(&users).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetAllUsersTest] error fetching users: %v\n", err)
// 		return nil, err
// 	}
// 	return users, nil
// }

// // GetAllProductsTest возвращает все товары
// func GetAllProductsTest() ([]models.Product, error) {
// 	var products []models.Product
// 	err := db.GetDBConn().Find(&products).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetAllProductsTest] error fetching products: %v\n", err)
// 		return nil, err
// 	}
// 	return products, nil
// }
