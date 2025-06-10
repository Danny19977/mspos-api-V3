package dashboard

import (
	"github.com/danny19977/mspos-api-v3/database"
	"github.com/gofiber/fiber/v2"
)

// total visit per day 50 and per week 300 and 100%(percentage)
// total Visit per month 1400 and 100%(percentage)
func TotalVisitsByProvince(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name        string  `json:"name"`
		Signature   string  `json:"signature"`
		Title       string  `json:"title"`
		TotalVisits int     `json:"total_visits"`
		Objectif    float64 `json:"objectif"`
		Target      int     `json:"target"`
	}

	err := db.Table("pos_forms").
		Select(`
		provinces.name AS name,
		pos_forms.signature AS signature,
		users.title AS title, 
		COUNT(pos_forms.signature) AS total_visits,
		(COUNT(pos_forms.signature) / (
			CASE
					WHEN users.title = 'ASM'  THEN 10 
					WHEN users.title = 'Supervisor'  THEN 20 
					WHEN users.title = 'DR'   THEN 40 
					WHEN users.title = 'Cyclo' THEN 40
					ELSE 1 
			END
		) ::numeric) * 100 AS objectif,
		(
			CASE
					WHEN users.title = 'ASM'  THEN 10 
					WHEN users.title = 'Supervisor'  THEN 20 
					WHEN users.title = 'DR'   THEN 40 
					WHEN users.title = 'Cyclo' THEN 40
					ELSE 1 
			END
		) AS target
		`).
		Joins("JOIN pos ON pos.uuid = pos_forms.pos_uuid").
		Joins("JOIN users ON users.uuid = pos.user_uuid").
		Joins("JOIN provinces ON provinces.uuid = pos_forms.province_uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?", country_uuid, province_uuid).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos_forms.deleted_at IS NULL").
		Group("provinces.name, pos_forms.signature, users.title").
		Order("provinces.name, pos_forms.signature").
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

func TotalVisitsByArea(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name        string  `json:"name"`
		Signature   string  `json:"signature"`
		Title       string  `json:"title"`
		TotalVisits int     `json:"total_visits"`
		Objectif    float64 `json:"objectif"`
		Target      int     `json:"target"`
	}

	err := db.Table("pos_forms").
		Select(`
		areas.name AS name,
		pos_forms.signature AS signature,
		users.title AS title, 
		COUNT(pos_forms.signature) AS total_visits,
		(COUNT(pos_forms.signature) / (
			CASE
					WHEN users.title = 'ASM'  THEN 10 
					WHEN users.title = 'Supervisor'  THEN 20 
					WHEN users.title = 'DR'   THEN 40 
					WHEN users.title = 'Cyclo' THEN 40
					ELSE 1 
			END
		) ::numeric) * 100 AS objectif,
		(
			CASE
					WHEN users.title = 'ASM'  THEN 10 
					WHEN users.title = 'Supervisor'  THEN 20 
					WHEN users.title = 'DR'   THEN 40 
					WHEN users.title = 'Cyclo' THEN 40
					ELSE 1 
			END
		) AS target
		`).
		Joins("JOIN pos ON pos.uuid = pos_forms.pos_uuid").
		Joins("JOIN users ON users.uuid = pos.user_uuid").
		Joins("JOIN areas ON pos_forms.area_uuid = areas.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?", country_uuid, province_uuid).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos_forms.deleted_at IS NULL").
		Group("areas.name, pos_forms.signature, users.title").
		Order("areas.name, pos_forms.signature").
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

func TotalVisitsBySubArea(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	area_uuid := c.Query("area_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name        string  `json:"name"`
		Signature   string  `json:"signature"`
		Title       string  `json:"title"`
		TotalVisits int     `json:"total_visits"`
		Objectif    float64 `json:"objectif"`
		Target      int     `json:"target"`
	}

	err := db.Table("pos_forms").
		Select(`
		sub_areas.name AS name,
		pos_forms.signature AS signature,
		users.title AS title, 
		COUNT(pos_forms.signature) AS total_visits,
		(COUNT(pos_forms.signature) / (
			CASE
					WHEN users.title = 'ASM'  THEN 10 
					WHEN users.title = 'Supervisor'  THEN 20 
					WHEN users.title = 'DR'   THEN 40 
					WHEN users.title = 'Cyclo' THEN 40
					ELSE 1 
			END
		) ::numeric) * 100 AS objectif,
		(
			CASE
					WHEN users.title = 'ASM'  THEN 10 
					WHEN users.title = 'Supervisor'  THEN 20 
					WHEN users.title = 'DR'   THEN 40 
					WHEN users.title = 'Cyclo' THEN 40
					ELSE 1 
			END
		) AS target
		`).
		Joins("JOIN pos ON pos.uuid = pos_forms.pos_uuid").
		Joins("JOIN users ON users.uuid = pos.user_uuid").
		Joins("JOIN sub_areas ON pos_forms.sub_area_uuid = sub_areas.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ?", country_uuid, province_uuid, area_uuid).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos_forms.deleted_at IS NULL").
		Group("sub_areas.name, pos_forms.signature, users.title").
		Order("sub_areas.name, pos_forms.signature").
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

func TotalVisitsByCommune(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	area_uuid := c.Query("area_uuid")
	sub_area_uuid := c.Query("sub_area_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name        string  `json:"name"`
		Signature   string  `json:"signature"`
		Title       string  `json:"title"`
		TotalVisits int     `json:"total_visits"`
		Objectif    float64 `json:"objectif"`
		Target      int     `json:"target"`
	}

	err := db.Table("pos_forms").
		Select(`
		communes.name AS name,
		pos_forms.signature AS signature,
		users.title AS title, 
		COUNT(pos_forms.signature) AS total_visits,
		(COUNT(pos_forms.signature) / (
			CASE
					WHEN users.title = 'ASM'  THEN 10 
					WHEN users.title = 'Supervisor'  THEN 20 
					WHEN users.title = 'DR'   THEN 40 
					WHEN users.title = 'Cyclo' THEN 40
					ELSE 1 
			END
		) ::numeric) * 100 AS objectif,
		(
			CASE
					WHEN users.title = 'ASM'  THEN 10 
					WHEN users.title = 'Supervisor'  THEN 20 
					WHEN users.title = 'DR'   THEN 40 
					WHEN users.title = 'Cyclo' THEN 40
					ELSE 1 
			END
		) AS target
		`).
		Joins("JOIN pos ON pos.uuid = pos_forms.pos_uuid").
		Joins("JOIN users ON users.uuid = pos.user_uuid").
		Joins("JOIN communes ON pos_forms.sub_area_uuid = communes.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.sub_area_uuid = ?", country_uuid, province_uuid, area_uuid, sub_area_uuid).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos_forms.deleted_at IS NULL").
		Group("communes.name, pos_forms.signature, users.title").
		Order("communes.name, pos_forms.signature").
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
