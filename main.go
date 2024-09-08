// C:\GoProject\src\eShop\main.go

package main

import (
	app "eShop/cmd"
	"eShop/utils"
	"fmt"
	"time"
)

func main() {
	utils.ClearConsole()
	fmt.Printf("Starting server: %s...\n", time.Now().Format("2006-01-02 15:04:05"))

	app.RunApp()
}
