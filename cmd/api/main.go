package main

import (
	"context"
	"log"
	"time"

	"github.com/gabrielmatsan/forum-golang-api/config"
	pgstore "github.com/gabrielmatsan/forum-golang-api/internal/infra/db/sqlc"
	"github.com/gabrielmatsan/forum-golang-api/internal/interface/modules"
	"github.com/gabrielmatsan/forum-golang-api/internal/interface/routes"
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

	db := pgstore.New(conn)

	// Modulos
	studentModule := modules.NewStudentsModule(db)

	r := gin.Default() // Cria uma instância padrão do Gin

	routes.RegisterStudentsRoutes(r, studentModule.CreateStudentController)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // Inicia o servidor na porta 8080
}
