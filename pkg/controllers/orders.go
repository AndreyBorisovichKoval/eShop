// C:\GoProject\src\eShop\pkg\controllers\orders.go

package controllers

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddOrder
// @Summary Create a new order and add items to it
// @Tags orders
// @Description Creates a new order and adds items to it (Seller only)
// @ID add-order
// @Accept json
// @Produce json
// @Param input body []models.OrderItem true "List of order items"
// @Success 201 {string} string "Order created successfully!!!"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /orders [post]
func AddOrder(c *gin.Context) {
	// Получаем ID пользователя (продавца) из токена
	userID, exists := c.Get(userIDCtx)
	if !exists {
		handleError(c, errs.ErrUnauthorized)
		return
	}

	// Приводим userID к uint
	uid, ok := userID.(uint)
	if !ok {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Получаем список товаров для заказа
	var orderItems []models.OrderItem
	if err := c.BindJSON(&orderItems); err != nil {
		logger.Error.Printf("[controllers.AddOrder] error binding order items: %v\n", err)
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("User ID [%d] is creating an order\n", uid)

	// Создаем заказ через сервис
	order, err := service.CreateOrder(uid, orderItems)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("User ID [%d] successfully created an order with ID: %d\n", uid, order.ID)
	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully!!!"})
}

// MarkOrderAsPaid
// @Summary Mark an order as paid
// @Tags orders
// @Description Mark a specific order as paid
// @ID mark-order-as-paid
// @Param id path int true "Order ID"
// @Success 200 {string} string "Order marked as paid successfully"
// @Failure 404 {object} ErrorResponse "Order not found"
// @Failure 409 {object} ErrorResponse "Order already paid"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /orders/{id}/pay [patch]
func MarkOrderAsPaid(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("User requested to mark order ID [%d] as paid\n", orderID)

	// Call the service to mark the order as paid
	err = service.MarkOrderAsPaid(uint(orderID))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("Order ID [%d] successfully marked as paid\n", orderID)
	c.JSON(http.StatusOK, gin.H{"message": "Order marked as paid successfully"})
}

// // DeleteOrderItem
// // @Summary Remove item from the order
// // @Tags orders
// // @Description Removes a specific item from an order
// // @ID delete-order-item
// // @Param order_id path int true "Order ID"
// // @Param item_id path int true "Order Item ID"
// // @Success 200 {string} string "Order item deleted successfully"
// // @Failure 404 {object} ErrorResponse "Order item not found"
// // @Failure 500 {object} ErrorResponse "Server error"
// // @Router /orders/{order_id}/items/{item_id} [delete]
// func DeleteOrderItem(c *gin.Context) {
// 	orderID, err := strconv.Atoi(c.Param("order_id"))
// 	if err != nil {
// 		handleError(c, errs.ErrValidationFailed)
// 		return
// 	}

// 	itemID, err := strconv.Atoi(c.Param("item_id"))
// 	if err != nil {
// 		handleError(c, errs.ErrValidationFailed)
// 		return
// 	}

// 	logger.Info.Printf("User requested to delete item ID [%d] from order ID [%d]\n", itemID, orderID)

// 	err = service.DeleteOrderItem(uint(orderID), uint(itemID))
// 	if err != nil {
// 		handleError(c, err)
// 		return
// 	}

// 	logger.Info.Printf("Successfully deleted item ID [%d] from order ID [%d]\n", itemID, orderID)
// 	c.JSON(http.StatusOK, gin.H{"message": "Order item deleted successfully"})
// }

// DeleteOrderItem
// @Summary Remove item from the order
// @Tags orders
// @Description Removes a specific item from an order
// @ID delete-order-item
// @Param order_id path int true "Order ID"
// @Param item_id path int true "Order Item ID"
// @Success 200 {string} string "Order item deleted successfully"
// @Failure 404 {object} ErrorResponse "Order item not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /orders/{order_id}/items/{item_id} [delete]
func DeleteOrderItem(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("User requested to delete item ID [%d] from order ID [%d]\n", itemID, orderID)

	err = service.DeleteOrderItem(uint(orderID), uint(itemID))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("Successfully deleted item ID [%d] from order ID [%d]\n", itemID, orderID)
	c.JSON(http.StatusOK, gin.H{"message": "Order item deleted successfully"})
}
