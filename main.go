package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/francischacko/d4donline/pkg/configs"
	"github.com/francischacko/d4donline/pkg/controllers"
	"github.com/francischacko/d4donline/pkg/db"
	"github.com/francischacko/d4donline/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

const maxRequestsPerCPU = 5

func init() {
	configs.InitEnvConfigs()
	db.Connectdb()
	db.Syncdb()
}
func main() {
	app := fiber.New()
	// app.Get("/", controllers.HelloWorld)
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // adding compression middleware
	}))
	app.Get("getoffers/:country", middleware.RateLimiter(maxRequestsPerCPU), controllers.GetDataByCountry)

	ctx, cancel := context.WithCancel(context.Background())

	// Create a channel to receive OS interrupt signals (e.g., Ctrl+C)
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		if err := app.Listen(configs.EnvConf.LocalServerPort); err != nil {
			log.Fatalf("Error starting the server: %v", err)
		}
	}()

	// Wait for an interrupt signal
	<-interruptChan

	// Call the cancel function to signal the server to stop accepting new requests
	cancel()

	log.Println("Gracefully shutting down the server...")

	// Wait for a short duration to allow the server to shut down gracefully
	timeout := 5 * time.Second
	select {
	case <-ctx.Done():
	case <-time.After(timeout):
	}

	log.Println("Server shutdown complete.")
}