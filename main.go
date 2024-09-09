// C:\GoProject\src\eShop\main.go

package main

import (
	app "eShop/cmd"
	"eShop/utils"
)

func main() {
	// Очищаем консоль от старых сообщений...
	utils.ClearConsole()

	app.RunApp()
}
