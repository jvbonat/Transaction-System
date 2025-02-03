package main

import (
	"desafio-transacoes/db"
	"desafio-transacoes/routes"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when creating .ENV file")
	}

	err = db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

    db.Migrate()

    router := gin.Default()
    routes.SetupRoutes(router)

    router.Run(":8080")
}
