package dashboard

import (
	"github.com/danny19977/mspos-api-v3/database"
	"github.com/gofiber/fiber/v2"
)

func SosTableViewProvince(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name             string  `json:"name"`
		UUID             string  `json:"uuid"`
		BrandName        string  `json:"brand_name"`
		TotalFarde       float64 `json:"total_farde"`
		TotalGlobalFarde float64 `json:"total_global_farde"`
		Percentage       float64 `json:"percentage"`
		TotalPos         int64   `json:"total_pos"`
	}

	err := db.Table("pos_form_items").
		Select(`
		provinces.name AS name,
		provinces.uuid AS uuid,
		brands.name AS brand_name, 
		ROUND(SUM(pos_form_items.number_farde)::numeric, 2) AS total_farde,
		(SELECT SUM(pos_form_items.number_farde) 
		 FROM pos_form_items 
		 INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		 WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?
		) AS total_global_farde,
		ROUND((SUM(pos_form_items.number_farde) * 100.0 / (SELECT SUM(pos_form_items.number_farde) 
		 FROM pos_form_items 
		 INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		 WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?))::numeric, 2) AS percentage,
		(SELECT COUNT(DISTINCT pos_forms.pos_uuid) 
		 FROM pos_form_items 
		 INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		 WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?
		 ) AS total_pos
	`, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date).
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?", country_uuid, province_uuid).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos_forms.deleted_at IS NULL").
		Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
		Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
		Joins("INNER JOIN provinces ON pos_forms.province_uuid = provinces.uuid").
		Group("provinces.name, provinces.uuid, brands.name").
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

func SosTableViewArea(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name             string  `json:"name"`
		UUID             string  `json:"uuid"`
		BrandName        string  `json:"brand_name"`
		TotalFarde       float64 `json:"total_farde"`
		TotalGlobalFarde float64 `json:"total_global_farde"`
		Percentage       float64 `json:"percentage"`
		TotalPos         int64   `json:"total_pos"`
	}

	err := db.Table("pos_form_items").
		Select(`
			areas.name AS name,
			areas.uuid AS uuid,
			brands.name AS brand_name, 
			ROUND(SUM(pos_form_items.number_farde)::numeric, 2) AS total_farde,
			(SELECT SUM(pos_form_items.number_farde) 
			 FROM pos_form_items 
			 INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
			 WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?
			) AS total_global_farde,
			ROUND((SUM(pos_form_items.number_farde) * 100.0 / (SELECT SUM(pos_form_items.number_farde) 
			 FROM pos_form_items 
			 INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
			 WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?))::numeric, 2) AS percentage,
			(SELECT COUNT(DISTINCT pos_forms.pos_uuid) 
			FROM pos_form_items 
			INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
			WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?
			) AS total_pos
		`, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date).
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?", country_uuid, province_uuid).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos_forms.deleted_at IS NULL").
		Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
		Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
		Joins("INNER JOIN areas ON pos_forms.area_uuid = areas.uuid").
		Group("areas.name, areas.uuid, brands.name").
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

func SosTableViewSubArea(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	area_uuid := c.Query("area_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name             string  `json:"name"`
		UUID             string  `json:"uuid"`
		BrandName        string  `json:"brand_name"`
		TotalFarde       float64 `json:"total_farde"`
		TotalGlobalFarde float64 `json:"total_global_farde"`
		Percentage       float64 `json:"percentage"`
		TotalPos         int64   `json:"total_pos"`
	}

	err := db.Table("pos_form_items").
		Select(`
			sub_areas.name AS name,
			sub_areas.uuid AS uuid,
			brands.name AS brand_name, 
			ROUND(SUM(pos_form_items.number_farde)::numeric, 2) AS total_farde,
			(SELECT SUM(pos_form_items.number_farde) 
			 FROM pos_form_items 
			 INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
			 WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?
			) AS total_global_farde,
			ROUND((SUM(pos_form_items.number_farde) * 100.0 / (SELECT SUM(pos_form_items.number_farde) 
			 FROM pos_form_items 
			 INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
			 WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?))::numeric, 2) AS percentage,
			(SELECT COUNT(DISTINCT pos_forms.pos_uuid) 
			FROM pos_form_items 
			INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
			WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?
			) AS total_pos
		`, country_uuid, province_uuid, area_uuid, start_date, end_date, country_uuid, province_uuid, area_uuid, start_date, end_date, country_uuid, province_uuid, area_uuid, start_date, end_date).
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ?", country_uuid, province_uuid, area_uuid).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos_forms.deleted_at IS NULL").
		Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
		Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
		Joins("INNER JOIN sub_areas ON pos_forms.sub_area_uuid = sub_areas.uuid").
		Group("sub_areas.name, sub_areas.uuid, brands.name").
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

func SosTableViewCommune(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	area_uuid := c.Query("area_uuid")
	sub_area_uuid := c.Query("sub_area_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name             string  `json:"name"`
		UUID             string  `json:"uuid"`
		BrandName        string  `json:"brand_name"`
		TotalFarde       float64 `json:"total_farde"`
		TotalGlobalFarde float64 `json:"total_global_farde"`
		Percentage       float64 `json:"percentage"`
		TotalPos         int64   `json:"total_pos"`
	}

	err := db.Table("pos_form_items").
		Select(`
			communes.name AS name,
			communes.uuid AS uuid,
			brands.name AS brand_name, 
			ROUND(SUM(pos_form_items.number_farde)::numeric, 2) AS total_farde,
			(SELECT SUM(pos_form_items.number_farde) 
			 FROM pos_form_items 
			 INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
			 WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.sub_area_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?
			) AS total_global_farde,
			ROUND((SUM(pos_form_items.number_farde) * 100.0 / (SELECT SUM(pos_form_items.number_farde) 
			 FROM pos_form_items 
			 INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
			 WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.sub_area_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?))::numeric, 2) AS percentage,
			(SELECT COUNT(DISTINCT pos_forms.pos_uuid) 
			FROM pos_form_items 
			INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
			WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.sub_area_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?
			) AS total_pos
		`, country_uuid, province_uuid, area_uuid, sub_area_uuid, start_date, end_date, country_uuid, province_uuid, area_uuid, sub_area_uuid, start_date, end_date, country_uuid, province_uuid, area_uuid, sub_area_uuid, start_date, end_date).
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.sub_area_uuid = ?", country_uuid, province_uuid, area_uuid, sub_area_uuid).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos_forms.deleted_at IS NULL").
		Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
		Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
		Joins("INNER JOIN communes ON pos_forms.commune_uuid = communes.uuid").
		Group("communes.name, communes.uuid, brands.name").
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

func SosTotalByBrandByMonth(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	year := c.Query("year")

	var results []struct {
		BrandName        string  `json:"brand_name"`
		Month            int     `json:"month"`
		TotalFarde       float64 `json:"total_farde"`
		TotalGlobalFarde float64 `json:"total_global_farde"`
		Percentage       float64 `json:"percentage"`
		TotalPos         int64   `json:"total_pos"`
	}

	err := db.Table("pos_form_items").
		Select(`
		brands.name AS brand_name,
		EXTRACT(MONTH FROM pos_forms.created_at) AS month, 
		ROUND(SUM(pos_form_items.number_farde)::numeric, 2) AS total_farde,
		(SELECT SUM(pos_form_items.number_farde) 
		 FROM pos_form_items 
		 INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		 WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND EXTRACT(YEAR FROM pos_forms.created_at) = ?
		) AS total_global_farde,
		ROUND((SUM(pos_form_items.number_farde) * 100.0 / (SELECT SUM(pos_form_items.number_farde) 
		 FROM pos_form_items 
		 INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		 WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND EXTRACT(YEAR FROM pos_forms.created_at) = ?))::numeric, 2) AS percentage,
		(SELECT COUNT(DISTINCT pos_forms.pos_uuid) 
		 FROM pos_form_items 
		 INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		 WHERE pos_form_items.deleted_at IS NULL AND pos_forms.country_uuid = ? AND EXTRACT(YEAR FROM pos_forms.created_at) = ?
		 ) AS total_pos
	`, country_uuid, year, country_uuid, year, country_uuid, year).
		Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
		Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
		Where("pos_forms.country_uuid = ? AND EXTRACT(YEAR FROM pos_forms.created_at) = ?", country_uuid, year).
		Where("pos_forms.deleted_at IS NULL").
		Group("brands.name, month").
		Order("brands.name, month ASC").
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
		"message": "results data",
		"data":    results,
	})
}
