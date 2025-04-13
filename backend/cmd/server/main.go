package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"promotion-scraper/internal/auth"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	r := gin.Default()

	// Teste
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/auth/login", auth.LoginHandler)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
