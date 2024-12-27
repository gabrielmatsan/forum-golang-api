package main

import (
	"context"
	"log"
	"time"

	"github.com/gabrielmatsan/forum-golang-api/config"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := config.ConnectDB(ctx)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer conn.Close()
}
