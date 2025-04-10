package dashboard

import (
	"fmt"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/gofiber/fiber/v2"
)

// total visit per day 50 and per week 300 and 100%(percentage)
// total Visit per month 1400 and 100%(percentage)
func CycloVisiteByMonth(c *fiber.Ctx) error {
	db := database.DB
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	fmt.Println("db: ", db)
	fmt.Println("start_date: ", start_date)
	fmt.Println("end_date: ", end_date)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    "",
	})
}

// total visit per day 40 and per week 240 and 100%(percentage)
// total Visit per month 1220 and 100%(percentage)
func DRVisiteByMonth(c *fiber.Ctx) error {
	db := database.DB
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	fmt.Println("db: ", db)
	fmt.Println("start_date: ", start_date)
	fmt.Println("end_date: ", end_date)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    "",
	})
}

// total visit per day 30 and per week 180 and 100%(percentage)
// total Visit per month 840 and 100%(percentage)
func SupVisiteByMonth(c *fiber.Ctx) error {
	db := database.DB
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	fmt.Println("db: ", db)
	fmt.Println("start_date: ", start_date)
	fmt.Println("end_date: ", end_date)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    "",
	})
}

// total visit per day 20 and per week 120 and 100%(percentage)
// total Visit per month 560 and 100%(percentage)
func ASMVisiteByMonth(c *fiber.Ctx) error {
	db := database.DB
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	fmt.Println("db: ", db)
	fmt.Println("start_date: ", start_date)
	fmt.Println("end_date: ", end_date)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    "",
	})
}
