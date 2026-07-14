package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	// Keep health checks outside the logged API group to avoid noisy probe logs.
	app.Head("/api/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent) // 204 – lightweight, no body
	})
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
	})

	api := app.Group("/api", logger.New())

	// Setup all route groups
	setupAuthRoutes(api)
	setupUsersRoutes(api)
	setupGeographicRoutes(api)
	setupHierarchyRoutes(api)
	setupPosRoutes(api)
	setupRoutePlanRoutes(api)
	setupBrandRoutes(api)
	setupPosFormRoutes(api)
	setupObservationRoutes(api)
	setupUserLogsRoutes(api)
	setupDashboardRoutes(api)
}
