package dashboard

import (
	"fmt"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/gofiber/fiber/v2"
)

// Total POS Grosseste & and Detaillant per Area and SubArea
func CategoriesChart(c *fiber.Ctx) error {
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

// Total times a Cyclo/DR/Sup/ASM visted a POS per month in Pie Chart formate
func POSPie(c *fiber.Ctx) error {
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

// Total times a Cyclo/DR/Sup/ASM visted a POS per month in Table formate
func POSTableView(c *fiber.Ctx) error {
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

// Total Stock Found per POS Fruads numbers filtered by date
func StockTableView(c *fiber.Ctx) error {
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

// Aerage Stock Found at a POS numbers filtered by date
func AverageStockTableView(c *fiber.Ctx) error {
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

// Total stock Cyclo,DR,Sup and ASM sold to POS per month
func SoldTableView(c *fiber.Ctx) error {
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

func SoldPie(c *fiber.Ctx) error {
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
