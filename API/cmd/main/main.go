package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vidacalura/CS-STV/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	router := routes.NewRouter()

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err.Error())
	}
}
