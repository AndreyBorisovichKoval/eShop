// C:\GoProject\src\eShop\main.go

package main

import (
	app "eShop/cmd"
	"eShop/utils"
)

// @title eShop API
// @version 0.0001
// @description API Server for eShop Application

// @host localhost:8585
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Очищаем консоль от старых сообщений...
	utils.ClearConsole()

	app.RunApp()
}

// swag init -g cmd\app.go
// go build -o swag.exe cmd/app.go
