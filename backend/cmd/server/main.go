package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"promotion-scraper/internal/auth"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	auth.InitOAuthConfig() // ← aqui, depois do .env

	r := gin.Default()
	r.GET("/auth/login", auth.LoginHandler)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
