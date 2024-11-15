package main

import (
	"entain-golang-task/cmd/app"
)

func main() {
	application := app.NewApp()
	application.Run()
}
