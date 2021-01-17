package main

import (
	"nixLevelFour/application"
	"time"
)

func main() {
	app := application.InitApp()
	app.Start()
	time.Sleep(time.Second * 5)
}
