// C:\GoProject\src\eShop\pkg\controllers\routes.go

package controllers

import (
	"eShop/configs"
	_ "eShop/docs"
	"eShop/logger"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// localhost:8585/swagger/index.html

// PingPong handles the ping request and responds with a pong message...
func PingPong(c *gin.Context) {
	logger.Info.Printf("Route '%s' called with method '%s'", c.FullPath(), c.Request.Method)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func InitRoutes() *gin.Engine {
	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", PingPong)

	authG := router.Group("/auth")
	{
		authG.POST("/sign-in", SignIn)
	}

	// userG := router.Group("/users")
	userG := router.Group("/users", checkUserAuthentication, checkUserBlocked())
	{
		userG.POST("", CreateUser)                        // Регистрация нового пользователя...
		userG.GET("", GetAllUsers)                        // Получение списка всех активных пользователей...
		userG.GET("/deleted", GetAllDeletedUsers)         // Получение списка всех удалённых пользователей...
		userG.GET("/:id", GetUserByID)                    // Получение данных пользователя по его ID...
		userG.PATCH("/:id", UpdateUserByID)               // Обновление данных пользователя по его ID...
		userG.PATCH("/:id/block", BlockUserByID)          // Блокировка пользователя по его ID...
		userG.PATCH("/:id/unblock", UnblockUserByID)      // Разблокировка пользователя по его ID...
		userG.PATCH("/:id/restore", RestoreUserByID)      // Восстановление удалённого пользователя...
		userG.DELETE("/:id/soft", SoftDeleteUserByID)     // Софт удаление пользователя по его ID...
		userG.DELETE("/:id/hard", HardDeleteUserByID)     // Полное удаление пользователя по его ID (хард удаление)...
		userG.PATCH("/password", ChangeOwnPassword)       // Обновление пароля текущего пользователя...
		userG.PATCH("/:id/reset-password", ResetPassword) // Сброс пароля пользователя (доступно только администратору)...
		userG.GET("/settings", GetUserSettingsByID)       // Получение настроек пользователя...
		userG.PATCH("/settings", UpdateUserSettings)      // Обновление настроек пользователя...
	}

	taxesG := router.Group("/taxes", checkUserAuthentication, checkUserBlocked())
	{
		taxesG.GET("", GetAllTaxes)         // Получение списка всех текущих налогов...
		taxesG.PATCH("/:id", UpdateTaxByID) // Обновление налоговой ставки по ID...
	}

	supplierG := router.Group("/suppliers", checkUserAuthentication, checkUserBlocked())
	{
		supplierG.POST("", CreateSupplier)                    // Регистрация нового поставщика...
		supplierG.GET("", GetAllSuppliers)                    // Получение списка всех активных поставщиков...
		supplierG.GET("/deleted", GetAllDeletedSuppliers)     // Получение списка всех удалённых поставщиков...
		supplierG.GET("/:id", GetSupplierByID)                // Получение данных поставщика по его ID...
		supplierG.PATCH("/:id", UpdateSupplierByID)           // Обновление данных поставщика по его ID...
		supplierG.PATCH("/:id/restore", RestoreSupplierByID)  // Восстановление удалённого поставщика...
		supplierG.DELETE("/:id/soft", SoftDeleteSupplierByID) // Софт удаление поставщика по его ID...
		supplierG.DELETE("/:id/hard", HardDeleteSupplierByID) // Полное удаление поставщика...
	}

	categoryG := router.Group("/categories", checkUserAuthentication, checkUserBlocked())
	{
		categoryG.POST("", CreateCategory)                    // Регистрация новой категории...
		categoryG.GET("", GetAllCategories)                   // Получение списка всех активных категорий...
		categoryG.GET("/deleted", GetAllDeletedCategories)    // Получение списка всех удалённых категорий...
		categoryG.GET("/:id", GetCategoryByID)                // Получение данных категории по её ID...
		categoryG.PATCH("/:id", UpdateCategoryByID)           // Обновление данных категории по её ID...
		categoryG.PATCH("/:id/restore", RestoreCategoryByID)  // Восстановление удалённой категории...
		categoryG.DELETE("/:id/soft", SoftDeleteCategoryByID) // Софт удаление категории по её ID...
		categoryG.DELETE("/:id/hard", HardDeleteCategoryByID) // Полное удаление категории...
	}

	productG := router.Group("/products", checkUserAuthentication, checkUserBlocked())
	{
		productG.POST("", AddProduct)                          // Добавление нового товара...
		productG.GET("", GetAllProducts)                       // Просмотр всех товаров...
		productG.GET("/:id", GetProductByID)                   // Просмотр товара по ID...
		productG.GET("/barcode/:barcode", GetProductByBarcode) // Просмотр товара по штрих-коду...
		productG.PATCH("/:id", UpdateProductByID)              // Обновление товара по ID...
		productG.PATCH("/:id/restore", RestoreProductByID)     // Восстановление удаленного продукта...
		productG.DELETE("/:id/soft", SoftDeleteProductByID)    // Мягкое удаление продукта...
		productG.DELETE("/:id/hard", HardDeleteProductByID)    // Полное удаление продукта...
	}

	orderG := router.Group("/orders", checkUserAuthentication, checkUserBlocked())
	{
		orderG.POST("", AddOrder)                                   // Создание нового заказа...
		orderG.PATCH("/:id/paid", MarkOrderAsPaid)                  // Отметка об оплате заказа...
		orderG.DELETE("/:order_id/items/:item_id", DeleteOrderItem) // Удаление товара из заказа...
		// orderG.GET("/:id", GetOrderByID)   // Получение заказа по ID...
		// orderG.DELETE("/:id", DeleteOrder) // Удаление заказа по ID...

		// orderG.POST("/:id/items", AddOrderItems)                    // Добавление товаров в заказ...
	}

	return router
}
