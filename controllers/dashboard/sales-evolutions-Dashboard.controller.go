package dashboard

import (
	"fmt"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/gofiber/fiber/v2"
)

// Total POS Grosseste & and Detaillant per Area and SubArea
func TypePosTableProvince(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		TypePos string `json:"type_pos"`
		Total   int    `json:"total"`
	}

	err := db.Table("pos").
		Select(`
		provinces.name AS name,
		pos.type AS type_pos, 
		COUNT(*) as total
		`).
		Joins("INNER JOIN provinces ON pos_forms.province_uuid = provinces.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?", country_uuid, province_uuid).
		Where("pos.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.deleted_at IS NULL").
		Group("provinces.name, pos.type").
		Order("provinces.name, pos.type DESC").
		Scan(&results).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    results,
	})
}

func TypePosTableArea(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		TypePos string `json:"type_pos"`
		Total   int    `json:"total"`
	}

	err := db.Table("pos").
		Select(`
		areas.name AS name,
		pos.type AS type_pos, 
		COUNT(*) as total
		`).
		Joins("INNER JOIN areas ON pos_forms.area_uuid = areas.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?", country_uuid, province_uuid).
		Where("pos.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.deleted_at IS NULL").
		Group("areas.name, pos.type").
		Order("areas.name, pos.type DESC").
		Scan(&results).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    results,
	})
}

func TypePosTableSubArea(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	area_uuid := c.Query("area_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		TypePos string `json:"type_pos"`
		Total   int    `json:"total"`
	}

	err := db.Table("pos").
		Select(`
		sub_areas.name AS name, 
		pos.type AS type_pos, 
		COUNT(*) as total
		`).
		Joins("INNER JOIN sub_areas ON pos_forms.sub_area_uuid = sub_areas.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ?", country_uuid, province_uuid, area_uuid).
		Where("pos.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.deleted_at IS NULL").
		Group("sub_areas.name, pos.type").
		Order("sub_areas.name, pos.type DESC").
		Scan(&results).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    results,
	})
}

func TypePosTableCommune(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	area_uuid := c.Query("area_uuid")
	sub_area_uuid := c.Query("sub_area_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		TypePos string `json:"type_pos"`
		Total   int    `json:"total"`
	}

	err := db.Table("pos").
		Select(`
		communes.name AS name,
		pos.type AS type_pos, 
		COUNT(*) as total
		`).
		Joins("INNER JOIN communes ON pos_forms.commune_uuid = communes.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.sub_area_uuid = ?", country_uuid, province_uuid, area_uuid, sub_area_uuid).
		Where("pos.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.deleted_at IS NULL").
		Group("communes.name, pos.type").
		Order("communes.name, pos.type DESC").
		Scan(&results).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    results,
	})
}

// Price table for POS per tige
func PriceTableProvince(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Price string `json:"price"`
	}

	err := db.Table("pos_forms").
		Select(`
		provinces.name AS name,
		SELECT price AS price,
		COUNT(*)
		FROM pos_forms  
		`).
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?", country_uuid, province_uuid).
		Where("pos.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.deleted_at IS NULL").
		Group("pos_forms.price").
		Order("provinces.name, pos_forms.price DESC").
		Scan(&results).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    results,
	})
}

func PriceTableArea(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Price string `json:"price"`
	}

	err := db.Table("pos_forms").
		Select(`
		areas.name AS name,
		SELECT price AS price,
		COUNT(*)
		FROM pos_forms  
		`).
		Joins("INNER JOIN areas ON pos_forms.area_uuid = areas.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?", country_uuid, province_uuid).
		Where("pos.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.deleted_at IS NULL").
		Group("areas.name, pos_forms.price").
		Order("areas.name, pos_forms.price DESC").
		Scan(&results).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    results,
	})
}

func PriceTableSubArea(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	area_uuid := c.Query("area_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Price string `json:"price"`
	}

	err := db.Table("pos_forms").
		Select(`
		sub_areas.name AS name,
		SELECT price AS price,
		COUNT(*)
		FROM pos_forms  
		`).
		Joins("INNER JOIN sub_areas ON pos_forms.sub_area_uuid = sub_areas.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ?", country_uuid, province_uuid, area_uuid).
		Where("pos.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.deleted_at IS NULL").
		Group("sub_areas.name, pos_forms.price").
		Order("sub_areas.name, pos_forms.price DESC").
		Scan(&results).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    results,
	})
}

func PriceTableCommune(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	area_uuid := c.Query("area_uuid")
	sub_area_uuid := c.Query("sub_area_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Price string `json:"price"`
	}

	err := db.Table("pos_forms").
		Select(`
		communes.name AS name,
		SELECT price AS price,
		COUNT(*)
		FROM pos_forms  
		`).
		Joins("INNER JOIN communes ON pos_forms.commune_uuid = communes.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.sub_area_uuid = ?", country_uuid, province_uuid, area_uuid, sub_area_uuid).
		Where("pos.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.deleted_at IS NULL").
		Group("communes.name, pos_forms.price").
		Order("communes.name, pos_forms.price DESC").
		Scan(&results).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    results,
	})
}


// Total Stock Found per POS Fruads numbers filtered by date KPI Summary
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
