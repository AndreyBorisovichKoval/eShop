// C:\GoProject\src\eShop\main.go

package main

import (
	app "eShop/cmd"
	"eShop/utils"
	"log"
	"time"
)

func main() {
	utils.ClearConsole()
	log.Printf("Starting server: %s...\n", time.Now().Format("2006-01-02 15:04:05"))

	app.RunApp()
}
