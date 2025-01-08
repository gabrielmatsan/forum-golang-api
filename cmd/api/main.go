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

	// Database connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := config.ConnectDB(ctx)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer conn.Close()
	db := pgstore.New(conn)

	// Modulos
	cryptoModule := modules.NewCryptographyModule()
	studentModule := modules.NewStudentsModule(db, cryptoModule)

	// API Config
	r := gin.Default()

	routes.RegisterStudentsRoutes(r, studentModule.CreateStudentController, studentModule.AuthenticateStudentController)

	// Endpoint Test (JUST IN DEVELOPMENT)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping-pong",
		})
	})

	r.Run(":8080") // Inicia o servidor na porta 8080
}
