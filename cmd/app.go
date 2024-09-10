// C:\GoProject\src\eShop\cmd\app.go

package app

import (
	"context"
	"eShop/configs"
	"eShop/db"
	"eShop/logger"
	"eShop/pkg/controllers"
	"eShop/server"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func RunApp() {
	// Запуск сервера...
	fmt.Printf("Starting server...\n\n")

	// Загружаем переменные окружения из файла .env...
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env file: %s", err)
	}
	fmt.Println("Environment variables loaded successfully.")

	// Чтение настроек из конфигурационного файла...
	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	fmt.Println("Configuration loaded successfully.")

	// Инициализация логгера...
	if err := logger.Init(); err != nil {
		log.Fatalf("Failed to initialize logger: %s", err)
	}
	fmt.Println("Logger initialized successfully.")

	// Подключение к базе данных с отложенным закрытием соединения...
	if err := db.ConnectToDB(); err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	defer db.CloseDBConn() // Закрытие соединения при завершении функции...
	fmt.Println("Database connected successfully.")

	// Выполнение миграций базы данных...
	if err := db.MigrateDB(); err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}
	fmt.Println("Database migrated successfully.")

	// Логирование успешного запуска сервера с указанием имени сервера и времени запуска...
	log.Printf("\n\nServer '%s' started at %s\n", configs.AppSettings.AppParams.ServerName, time.Now().Format("2006-01-02 15:04:05"))

	// Сообщение о прослушивании порта...
	fmt.Printf("Server is listening on port %v\n\n", configs.AppSettings.AppParams.PortRun)

	// Инициализация HTTP сервера...
	mainServer := new(server.Server)

	// Использование WaitGroup для синхронизации завершения работы...
	var wg sync.WaitGroup
	wg.Add(1)

	// Запуск сервера в отдельной горутине...
	go func() {
		defer wg.Done() // Уменьшаем счетчик при завершении горутины...
		if err := mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.InitRoutes()); err != nil {
			log.Fatalf("HTTP server failed to start: %v", err)
		}
	}()

	// Ожидание сигнала завершения работы (например, от операционной системы)...
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// Начало процедуры завершения работы сервера...
	fmt.Printf("\nShutting down server...\n")

	// Остановка HTTP сервера...
	if err := mainServer.Shutdown(context.Background()); err != nil {
		fmt.Println(err.Error())
		log.Fatalf("HTTP server shutdown failed: %v", err)
	}
	fmt.Println("Server shut down gracefully.")

	// Ожидание завершения всех горутин...
	wg.Wait()
	fmt.Println("Goodbye.")
}

// =================================================================

// // C:\GoProject\src\eShop\cmd\app.go

// package app

// import (
// 	"context"
// 	"eShop/configs"
// 	"eShop/db"
// 	"eShop/logger"
// 	"eShop/pkg/controllers"
// 	"eShop/server"
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/joho/godotenv"
// )

// func RunApp() {
// 	// Запуск сервера...
// 	fmt.Printf("Starting server launch...\n\n")

// 	// Загружаем переменные окружения из файла .env...
// 	if err := godotenv.Load(".env"); err != nil {
// 		log.Fatalf("Error loading .env file. Errors is %s...", err)
// 	}
// 	fmt.Println("Environment variables loaded successfully!!!")

// 	// Чтение настроек из конфигурационного файла...
// 	if err := configs.ReadSettings(); err != nil {
// 		log.Fatalf("Error loading configuration. Errors is %s...", err)
// 	}
// 	fmt.Println("Settings loaded successfully!!!")

// 	// Инициализация логгера...
// 	if err := logger.Init(); err != nil {
// 		log.Fatalf("Error initializing logger. Errors is %s...", err)
// 	}
// 	fmt.Println("Logger initialized!!!")

// 	// Подключение к базе данных...
// 	var err error
// 	if err = db.ConnectToDB(); err != nil {
// 		log.Fatalf("Error connecting to database. Errors is %s...", err)
// 	}
// 	fmt.Println("Connected to the database Successfully!!!")

// 	// Выполнение миграций для базы данных...
// 	if err = db.MigrateDB(); err != nil {
// 		log.Fatalf("Error migrating database. Errors is %s...", err)
// 	}
// 	fmt.Println("Database migration Successful!!!")

// 	// Логирование успешного запуска сервера с указанием имени сервера и времени запуска...
// 	log.Printf("\n\nServer '%s' Launched at %s!!!\n", configs.AppSettings.AppParams.ServerName, time.Now().Format("2006-01-02 15:04:05"))

// 	// Сообщение о прослушивании порта...
// 	fmt.Printf("Server is Listening on port %v...\n\n", configs.AppSettings.AppParams.PortRun)

// 	// Запуск HTTP сервера...
// 	mainServer := new(server.Server)
// 	go func() {
// 		if err = mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.InitRoutes()); err != nil {
// 			log.Fatalf("HTTP Server failed to start: %v", err)
// 		}
// 	}()

// 	// Ожидание сигнала завершения работы (например, от операционной системы)...
// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
// 	<-quit

// 	// Начало процедуры завершения работы сервера...
// 	fmt.Printf("\nStarting to shut down the server...\n")

// 	// Закрытие соединения с базой данных, если оно активно...
// 	db.CloseDBConn()
// 	fmt.Println("The connection to the database was closed successfully!!!")

// 	// Остановка HTTP сервера...
// 	if err = mainServer.Shutdown(context.Background()); err != nil {
// 		fmt.Println(err.Error())
// 		log.Fatalf("HTTP Server shutdown failed: %v...", err)
// 	}
// 	fmt.Println("Server shut down Gracefully...")

// 	// Финальное сообщение перед завершением работы...
// 	fmt.Println("Goodbye and good luck!!!")
// }
