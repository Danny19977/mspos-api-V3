package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// isAllowedOrigin returns true for origins that should be permitted.
// Using a function instead of a static list lets the server accept any
// private-network IP, which is essential when users connect through a VPN
// that may assign a different 192.168.x.x or 10.x.x.x address.
func isAllowedOrigin(origin string) bool {
	// Strip the scheme so we can check the host prefix cleanly.
	host := origin
	if strings.HasPrefix(host, "http://") {
		host = host[7:]
	} else if strings.HasPrefix(host, "https://") {
		host = host[8:]
	}

	// Allow all RFC-1918 private-network addresses (covers every VPN subnet).
	if strings.HasPrefix(host, "192.168.") ||
		strings.HasPrefix(host, "10.") ||
		strings.HasPrefix(host, "172.16.") ||
		strings.HasPrefix(host, "172.17.") ||
		strings.HasPrefix(host, "172.18.") ||
		strings.HasPrefix(host, "172.19.") ||
		strings.HasPrefix(host, "172.2") ||
		strings.HasPrefix(host, "172.30.") ||
		strings.HasPrefix(host, "172.31.") {
		return true
	}

	// Allow local development.
	if strings.HasPrefix(host, "localhost") || strings.HasPrefix(host, "127.0.0.1") {
		return true
	}

	// Allow the production deployment on Render.
	if strings.HasPrefix(origin, "https://mspos-v3.com") {
		return true
	}

	return false
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {

	database.Connect()

	app := fiber.New(fiber.Config{
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	// Initialize default config
	app.Use(logger.New())

	// Middleware
	app.Use(cors.New(cors.Config{
		// AllowOriginsFunc replaces the old hardcoded IP list.
		// This lets any phone on the VPN (192.168.x.x, 10.x.x.x, …) pass
		// CORS without needing to add every possible VPN-assigned address.
		AllowOriginsFunc: isAllowedOrigin,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	routes.Setup(app)

	log.Fatal(app.Listen(getPort()))

}
