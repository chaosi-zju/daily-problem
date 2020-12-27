package main

import (
	"github.com/chaosi-zju/daily-problem/internal/app"
	"os"
)

const defaultPort string = "3001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := app.SetupRoutes()

	err := r.Run(":" + port)
	if err != nil {
		panic("failed to start engine: " + err.Error())
	}
}
