package main

import (
	"log"

	"github.com/joho/godotenv"
	"src/cmd"
	"src/config"
	"src/db"
)

func main() {

	if godotenv.Load(".env") != nil {
		log.Fatal("Error loading .env file")
	}

	confVars, configErr := config.New()

	if configErr != nil {
		log.Fatal(configErr)
	}

	dbErr := db.Init()

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	defer db.Close()

	app := cmd.InitApp()

	err := app.Listen(confVars.Port)
	if err != nil {
		return
	}
}
