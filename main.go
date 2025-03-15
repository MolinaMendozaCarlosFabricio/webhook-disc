package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"webhock-disc.com/w/src/web/infrastructure"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	engine := gin.Default()

	infrastructure.Routes(engine)

	engine.Run(":" + port)
}