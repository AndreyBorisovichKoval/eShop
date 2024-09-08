// C:\GoProject\src\eShop\main.go

package main

import (
	app "eShop/cmd"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearConsole() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clearConsole()
	fmt.Println(time.Now())
	app.RunApp()
}
