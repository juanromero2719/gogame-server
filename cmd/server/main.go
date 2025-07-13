package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/juanromero2719/gogame-server/internal/router"
)

func main() {

	_ = godotenv.Load()
	webServer := router.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor listo en :%s", port)
	if err := webServer.Start(":" + port); err != nil &&
		err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
