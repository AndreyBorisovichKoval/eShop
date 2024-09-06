// C:\GoProject\src\eShop\cmd\app.go

package app

import (
	"eShop/configs"
	"eShop/db"
	"eShop/logger"
	"eShop/pkg/controllers"
	"errors"
	"fmt"

	"github.com/joho/godotenv"
)

func RunApp() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(errors.New(fmt.Sprintf("Error loading .env file. Errors is %s.", err)))
		// log.Fatalf("Error loading .env file. Errors is %s.", err)
	}

	err = configs.ReadSettings()
	if err != nil {
		panic(err)
	}
	fmt.Println("Settings loaded successfully!!!")

	err = logger.Init()
	if err != nil {
		panic(err)
	}
	fmt.Println("Logger initialized!!!")

	err = db.ConnectToDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database Successfully!!!")

	defer func() {
		if err := db.CloseDB(); err != nil {
			fmt.Printf("Error closing database: %v\n", err)
		}
	}()

	err = db.MigrateDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database migration Successful!!!")
	// fmt.Println("Server is Listening on port 8585...")
	fmt.Printf("Server is Listening on port %v\n", configs.AppSettings.AppParams.PortRun)

	err = controllers.RunRoutes()
	if err != nil {
		panic(err)
	}
}
