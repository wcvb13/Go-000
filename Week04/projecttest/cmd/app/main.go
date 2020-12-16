package main

import (
	"log"
	"projecttest/internal/app/di"
)

func main() {
	app := di.InitApp()
	err := app.Start()
	log.Printf("app quit with err:%s",err.Error())
}