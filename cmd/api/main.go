package main

import (
	"context"
	"log"
	"time"

	"github.com/gabrielmatsan/forum-golang-api/config"
	"github.com/gin-gonic/gin"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := config.ConnectDB(ctx)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer conn.Close()

	r := gin.Default() // Cria uma instância padrão do Gin

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // Inicia o servidor na porta 8080
}
