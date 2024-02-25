package main

import (
	"github.com/AhmedSamy16/01-Book-Directory-Go/application"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := application.New()

	app.Start()
}
