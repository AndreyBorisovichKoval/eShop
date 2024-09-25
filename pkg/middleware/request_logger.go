// C:\GoProject\src\eShop\pkg\middleware\request_logger.go

package middleware

// import (
// 	"eShop/db"
// 	"eShop/logger"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// // RequestLoggerMiddleware captures user request data and saves it to the database using raw SQL
// func RequestLoggerMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Extracting necessary user data from the context
// 		userID, exists := c.Get("userID")
// 		if !exists {
// 			logger.Error.Println("UserID not found in context")
// 			c.Next() // If no user ID, skip logging
// 			return
// 		}

// 		// Extracting username and role with additional checks for nil values
// 		username, exists := c.Get("username")
// 		if !exists || username == nil {
// 			logger.Error.Println("Username not found or is nil in context")
// 			c.Next() // If no username, skip logging
// 			return
// 		}

// 		role, exists := c.Get("role")
// 		if !exists || role == nil {
// 			logger.Error.Println("Role not found or is nil in context")
// 			c.Next() // If no role, skip logging
// 			return
// 		}

// 		// Capturing request details
// 		path := c.Request.URL.Path
// 		method := c.Request.Method
// 		ip := c.ClientIP()

// 		// Current time for the request
// 		createdAt := time.Now()

// 		// Construct the raw SQL query
// 		sql := `INSERT INTO request_history (user_id, username, role, path, method, ip, created_at) 
//                 VALUES (?, ?, ?, ?, ?, ?, ?)`

// 		// Execute the raw SQL query
// 		if _, err := db.GetDBConn().Exec(sql, userID, username, role, path, method, ip, createdAt); err != nil {
// 			logger.Error.Printf("Failed to save request history with SQL: %v", err)
// 			c.JSON(500, gin.H{"error": "Failed to log request history"})
// 			c.Abort() // Останавливаем выполнение, если запись не удалась
// 			return
// 		} else {
// 			logger.Info.Println("Request history saved successfully with SQL!")
// 		}

// 		// Proceed to the next handler
// 		c.Next()
// 	}
// }

// package middleware

// import (
// 	"eShop/db"
// 	"eShop/logger"
// 	"eShop/models"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// // RequestLoggerMiddleware captures user request data and saves it to the database
// func RequestLoggerMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Extracting necessary user data from the context
// 		userID, exists := c.Get("userID")
// 		if !exists {
// 			logger.Error.Println("UserID not found in context")
// 			c.Next() // If no user ID, skip logging
// 			return
// 		}

// 		// Extracting username and role with additional checks for nil values
// 		username, exists := c.Get("username")
// 		if !exists || username == nil {
// 			logger.Error.Println("Username not found or is nil in context")
// 			c.Next() // If no username, skip logging
// 			return
// 		}

// 		role, exists := c.Get("role")
// 		if !exists || role == nil {
// 			logger.Error.Println("Role not found or is nil in context")
// 			c.Next() // If no role, skip logging
// 			return
// 		}

// 		// Capturing request details
// 		path := c.Request.URL.Path
// 		method := c.Request.Method
// 		ip := c.ClientIP()

// 		// Logging for debugging purposes
// 		logger.Info.Printf("Logging request from user: %v, path: %s, method: %s, IP: %s", userID, path, method, ip)

// 		// Creating a new request history entry
// 		requestHistory := models.RequestHistory{
// 			UserID:    userID.(uint),
// 			Username:  username.(string),
// 			Role:      role.(string),
// 			Path:      path,
// 			Method:    method,
// 			IP:        ip,
// 			CreatedAt: time.Now(),
// 		}

// 		// Saving the request history to the database with detailed error handling
// 		if err := db.GetDBConn().Create(&requestHistory).Error; err != nil {
// 			logger.Error.Printf("Failed to save request history: %v", err)
// 			c.JSON(500, gin.H{"error": "Failed to log request history"})
// 			c.Abort() // Останавливаем выполнение, если запись не удалась
// 			return
// 		} else {
// 			logger.Info.Println("Request history saved successfully!")
// 		}

// 		// Proceed to the next handler
// 		c.Next()
// 	}
// }

// package middleware

// import (
// 	"eShop/db"
// 	"eShop/logger"
// 	"eShop/models"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// // RequestLoggerMiddleware captures user request data and saves it to the database
// func RequestLoggerMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Extracting necessary user data from the context
// 		userID, exists := c.Get("userID")
// 		if !exists {
// 			logger.Error.Println("UserID not found in context")
// 			c.Next() // If no user ID, skip logging
// 			return
// 		}

// 		username, _ := c.Get("username")
// 		role, _ := c.Get("role")

// 		// Capturing request details
// 		path := c.Request.URL.Path
// 		method := c.Request.Method
// 		ip := c.ClientIP()

// 		// Logging for debugging purposes
// 		logger.Info.Printf("Logging request from user: %v, path: %s, method: %s, IP: %s", userID, path, method, ip)

// 		// Creating a new request history entry
// 		requestHistory := models.RequestHistory{
// 			UserID:    userID.(uint),
// 			Username:  username.(string),
// 			Role:      role.(string),
// 			Path:      path,
// 			Method:    method,
// 			IP:        ip,
// 			CreatedAt: time.Now(),
// 		}

// 		// Saving the request history to the database with detailed error handling
// 		if err := db.GetDBConn().Create(&requestHistory).Error; err != nil {
// 			logger.Error.Printf("Failed to save request history: %v", err)
// 			c.JSON(500, gin.H{"error": "Failed to log request history"})
// 			c.Abort() // Останавливаем выполнение, если запись не удалась
// 			return
// 		} else {
// 			logger.Info.Println("Request history saved successfully!")
// 		}

// 		// Proceed to the next handler
// 		c.Next()
// 	}
// }

// // package middleware

// // import (
// // 	"eShop/db"
// // 	"eShop/logger"
// // 	"eShop/models"
// // 	"time"

// // 	"github.com/gin-gonic/gin"
// // )

// // // RequestLoggerMiddleware captures user request data and saves it to the database
// // func RequestLoggerMiddleware() gin.HandlerFunc {
// // 	return func(c *gin.Context) {
// // 		// Extracting necessary user data from the context
// // 		userID, exists := c.Get("userID")
// // 		if !exists {
// // 			logger.Error.Println("UserID not found in context")
// // 			c.Next() // If no user ID, skip logging
// // 			return
// // 		}

// // 		username, _ := c.Get("username")
// // 		role, _ := c.Get("role")

// // 		// Capturing request details
// // 		path := c.Request.URL.Path
// // 		method := c.Request.Method
// // 		ip := c.ClientIP()

// // 		// Logging for debugging purposes
// // 		logger.Info.Printf("Logging request from user: %v, path: %s, method: %s, IP: %s", userID, path, method, ip)

// // 		// Creating a new request history entry
// // 		requestHistory := models.RequestHistory{
// // 			UserID:    userID.(uint),
// // 			Username:  username.(string),
// // 			Role:      role.(string),
// // 			Path:      path,
// // 			Method:    method,
// // 			IP:        ip,
// // 			CreatedAt: time.Now(),
// // 		}

// // 		// Saving the request history to the database using db.GetDBConn()
// // 		// if err := db.GetDBConn().Create(&requestHistory).Error; err != nil {
// // 		// 	logger.Error.Printf("Failed to save request history: %v", err)
// // 		// }
// // 		if err := db.GetDBConn().Create(&requestHistory).Error; err != nil {
// // 			logger.Error.Printf("Failed to save request history: %v", err)
// // 		} else {
// // 			logger.Info.Println("Request history saved successfully!")
// // 		}

// // 		// Proceed to the next handler
// // 		c.Next()
// // 	}
// // }

// // // package middleware

// // // import (
// // // 	"eShop/db"
// // // 	"eShop/models"
// // // 	"time"

// // // 	"github.com/gin-gonic/gin"
// // // )

// // // // RequestLoggerMiddleware captures user request data and saves it to the database
// // // func RequestLoggerMiddleware() gin.HandlerFunc {
// // // 	return func(c *gin.Context) {
// // // 		// Extracting necessary user data from the context
// // // 		userID, exists := c.Get("userID")
// // // 		if !exists {
// // // 			c.Next() // If no user ID, skip logging
// // // 			return
// // // 		}

// // // 		username, _ := c.Get("username")
// // // 		role, _ := c.Get("role")

// // // 		// Capturing request details
// // // 		path := c.Request.URL.Path
// // // 		method := c.Request.Method
// // // 		ip := c.ClientIP()

// // // 		// Creating a new request history entry
// // // 		requestHistory := models.RequestHistory{
// // // 			UserID:    userID.(uint),
// // // 			Username:  username.(string),
// // // 			Role:      role.(string),
// // // 			Path:      path,
// // // 			Method:    method,
// // // 			IP:        ip,
// // // 			CreatedAt: time.Now(),
// // // 		}

// // // 		// Saving the request history to the database using db.GetDBConn()
// // // 		db.GetDBConn().Create(&requestHistory)

// // // 		// Proceed to the next handler
// // // 		c.Next()
// // // 	}
// // // }

// // // // package middleware

// // // // import (
// // // // 	"eShop/models"
// // // // 	"time"

// // // // 	"github.com/gin-gonic/gin"
// // // // )

// // // // // RequestLoggerMiddleware captures user request data and saves it to the database
// // // // func RequestLoggerMiddleware() gin.HandlerFunc {
// // // // 	return func(c *gin.Context) {
// // // // 		// Extracting necessary user data from the context
// // // // 		userID, exists := c.Get("userID")
// // // // 		if !exists {
// // // // 			c.Next() // If no user ID, skip logging
// // // // 			return
// // // // 		}

// // // // 		username, _ := c.Get("username")
// // // // 		role, _ := c.Get("role")

// // // // 		// Capturing request details
// // // // 		path := c.Request.URL.Path
// // // // 		method := c.Request.Method
// // // // 		ip := c.ClientIP()

// // // // 		// Creating a new request history entry
// // // // 		requestHistory := models.RequestHistory{
// // // // 			UserID:    userID.(uint),
// // // // 			Username:  username.(string),
// // // // 			Role:      role.(string),
// // // // 			Path:      path,
// // // // 			Method:    method,
// // // // 			IP:        ip,
// // // // 			CreatedAt: time.Now(),
// // // // 		}

// // // // 		// Saving the request history to the database using dbConn
// // // // 		db.dbConn.Create(&requestHistory)

// // // // 		// Proceed to the next handler
// // // // 		c.Next()
// // // // 	}
// // // // }

// // // // // // RequestLoggerMiddleware captures user request data and saves it to the database
// // // // // func RequestLoggerMiddleware() gin.HandlerFunc {
// // // // // 	return func(c *gin.Context) {
// // // // // 		// Extracting necessary user data from the context
// // // // // 		userID, exists := c.Get("userID")
// // // // // 		if !exists {
// // // // // 			c.Next() // If no user ID, skip logging
// // // // // 			return
// // // // // 		}

// // // // // 		username, _ := c.Get("username")
// // // // // 		role, _ := c.Get("role")

// // // // // 		// Capturing request details
// // // // // 		path := c.Request.URL.Path
// // // // // 		method := c.Request.Method
// // // // // 		ip := c.ClientIP()

// // // // // 		// Creating a new request history entry
// // // // // 		requestHistory := models.RequestHistory{
// // // // // 			UserID:    userID.(uint),
// // // // // 			Username:  username.(string),
// // // // // 			Role:      role.(string),
// // // // // 			Path:      path,
// // // // // 			Method:    method,
// // // // // 			IP:        ip,
// // // // // 			CreatedAt: time.Now(),
// // // // // 		}

// // // // // 		// Saving the request history to the database
// // // // // 		db.DB.Create(&requestHistory)

// // // // // 		// Proceed to the next handler
// // // // // 		c.Next()
// // // // // 	}
// // // // // }
