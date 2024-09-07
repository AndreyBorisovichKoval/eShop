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
	"syscall"

	"github.com/joho/godotenv"
)

func RunApp() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file. Errors is %s...", err)
	}
	fmt.Println("Environment variables loaded successfully!!!")

	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Error loading configuration. Errors is %s...", err)
	}
	fmt.Println("Settings loaded successfully!!!")

	if err := logger.Init(); err != nil {
		log.Fatalf("Error initializing logger. Errors is %s...", err)
	}
	fmt.Println("Logger initialized!!!")

	// Connect to the database...
	var err error
	if err = db.ConnectToDB(); err != nil {
		log.Fatalf("Error connecting to database. Errors is %s...", err)
	}
	fmt.Println("Connected to the database Successfully!!!")

	if err = db.MigrateDB(); err != nil {
		log.Fatalf("Error migrating database. Errors is %s...", err)
	}
	fmt.Println("Database migration Successful!!!")

	fmt.Printf("Server is Listening on port %v\n", configs.AppSettings.AppParams.PortRun)

	mainServer := new(server.Server)
	go func() {
		if err = mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.InitRoutes()); err != nil {
			log.Fatalf("HTTP Server failed to start: %v", err)
		}
	}()

	// Ожидание сигнала завершения работы...
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Printf("\nStarting to shut down the server...\n")

	// Закрытие соединения с БД, если необходимо...
	db.CloseDBConn()
	fmt.Println("The connection to the database was closed successfully!!!")

	if err = mainServer.Shutdown(context.Background()); err != nil {
		log.Fatalf("HTTP Server shutdown failed: %v", err)
	}
	fmt.Println("Server shut down Gracefully...")
	fmt.Println("Goodbye and good luck!!!")
}
