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
	"net/http"
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
	fmt.Println("Environment variables loaded successfully!")

	// Чтение настроек из конфигурационного файла...
	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	fmt.Println("Configuration loaded successfully!")

	// Инициализация логгера...
	if err := logger.Init(); err != nil {
		log.Fatalf("Failed to initialize logger: %s", err)
	}
	fmt.Println("Logger initialized successfully!")

	// /

	// Подключение к базе данных с отложенным закрытием соединения...
	if err := db.ConnectToDB(); err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	defer func() { //TODO: handled closing DB connection
		err := db.CloseDBConn()
		if err != nil {
			log.Fatalf("Failed to close database connection: %s", err)
		}
	}() // Закрытие соединения при завершении функции...
	fmt.Println("Database connected successfully!")

	// Выполнение миграций базы данных...
	if err := db.MigrateDB(); err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}
	fmt.Println("Database migrated successfully!")

	// Логирование успешного запуска сервера с указанием имени сервера и времени запуска...
	log.Printf("\n\nServer '%s' started at %s...\n", configs.AppSettings.AppParams.ServerName, time.Now().Format("2006-01-02 15:04:05"))

	// Сообщение о прослушивании порта...
	fmt.Printf("Server is listening on port %v...\n\n", configs.AppSettings.AppParams.PortRun)

	// /

	// Инициализация HTTP сервера...
	// Создаём HTTP сервер...
	mainServer := new(server.Server)

	// Использование WaitGroup для синхронизации завершения работы (для ожидания завершения всех процессов)...
	var wg sync.WaitGroup
	wg.Add(1)

	// Запуск сервера в отдельной горутине...
	go func() {
		defer wg.Done() // Уменьшаем счетчик при завершении горутины...
		if err := mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			// Логируем ошибку, только если это не штатное завершение сервера...
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
	fmt.Println("Server shut down gracefully!")

	// Ожидание завершения всех горутин...
	wg.Wait()
	fmt.Println("Goodbye!!!")
}
