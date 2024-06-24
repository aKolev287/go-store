package main

import (
	"go-store-server/db"
	"go-store-server/routes"

	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	keys := os.Getenv("DB_CREDENTIALS")
	db.InitDB(keys)
	routes.Migrate()

	s := gin.Default()
	routes.ServerRouter(s)
	s.Run(":8000")
}
