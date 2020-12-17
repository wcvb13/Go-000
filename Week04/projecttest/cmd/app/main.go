package main

import (
	"log"
	"projecttest/internal/app/di"
)

func main() {
	app := di.InitApp()
	err := app.Start()
	if err != nil {
		log.Printf("app quit with err:%s\n",err.Error())
	}
}