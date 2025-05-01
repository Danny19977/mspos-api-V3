package dashboard

import (
	"github.com/danny19977/mspos-api-v3/database"
	"github.com/gofiber/fiber/v2"
)

// calculate the ND by Country and Province
func NdTableViewProvince(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name       string  `json:"name"`
		BrandName  string  `json:"brand_name"`
		TotalCount int     `json:"total_count"`
		Percentage float64 `json:"percentage"`
		TotalPos   int     `json:"total_pos"`
	}

	err := db.Table("pos_form_items").
		Select(`
		provinces.name AS name, 
		brands.name AS brand_name,
		SUM(pos_form_items.counter) AS total_count,
		ROUND((SUM(pos_form_items.counter) / (SELECT SUM(counter) FROM pos_form_items INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid WHERE pos_forms.country_uuid = ? AND  pos_forms.province_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?)) * 100, 2) AS percentage,
		(SELECT SUM(counter) FROM pos_form_items INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?) AS total_pos
		`, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date).
		Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
		Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
		Joins("INNER JOIN provinces ON pos_forms.province_uuid = provinces.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?", country_uuid, province_uuid).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Group("provinces.name, brands.name").
		Order("provinces.name, total_count DESC").
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

// calculate the ND by Area Found here
func NdTableViewArea(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name       string  `json:"name"`
		BrandName  string  `json:"brand_name"`
		TotalCount int     `json:"total_count"`
		Percentage float64 `json:"percentage"`
		TotalPos   int     `json:"total_pos"`
	}
	err := db.Table("pos_form_items").
		Select(`
		areas.name AS name, 
		brands.name AS brand_name, 
		SUM(pos_form_items.counter) AS total_count,
		ROUND((SUM(pos_form_items.counter) / (SELECT SUM(counter) FROM pos_form_items INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?)) * 100, 2) AS percentage,
		(SELECT SUM(counter) FROM pos_form_items INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?) AS total_pos
		`, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date).
		Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
		Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
		Joins("INNER JOIN areas ON pos_forms.area_uuid = areas.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?", country_uuid, province_uuid).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Group("areas.name, brands.name").
		Order("areas.name, total_count DESC").
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

// calculate the ND by Subarea Found here
func NdTableViewSubArea(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	area_uuid := c.Query("area_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name       string  `json:"name"`
		BrandName  string  `json:"brand_name"`
		TotalCount int     `json:"total_count"`
		Percentage float64 `json:"percentage"`
	}
	err := db.Table("pos_form_items").
		Select(`sub_areas.name AS name, 
		brands.name AS brand_name,
		 SUM(pos_form_items.counter) AS total_count,
		 ROUND((SUM(pos_form_items.counter) / (SELECT SUM(counter) FROM pos_form_items WHERE country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?)) * 100, 2) AS percentage,
		 (SELECT SUM(counter) FROM pos_form_items INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?) AS total_pos
		`, country_uuid, province_uuid, area_uuid, start_date, end_date, country_uuid, province_uuid, area_uuid, start_date, end_date).
		Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
		Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
		Joins("INNER JOIN sub_areas ON pos_forms.sub_area_uuid = sub_areas.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ?", country_uuid, province_uuid, area_uuid).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Group("sub_areas.name, brands.name").
		Order("sub_areas.name, total_count DESC").
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

// calculate the ND by Commune Found here
func NdTableViewCommune(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	area_uuid := c.Query("area_uuid")
	sub_area_uuid := c.Query("sub_area_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name       string  `json:"name"`
		BrandName  string  `json:"brand_name"`
		TotalCount int     `json:"total_count"`
		Percentage float64 `json:"percentage"`
	}

	err := db.Table("pos_form_items").
		Select(`
		communes.name AS name,
		brands.name AS brand_name,
		SUM(pos_form_items.counter) AS total_count,
		ROUND((SUM(pos_form_items.counter) / (SELECT SUM(counter) FROM pos_form_items WHERE country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.sub_area_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?)) * 100, 2) AS percentage,
		(SELECT SUM(counter) FROM pos_form_items INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.created_at BETWEEN ? AND ?) AS total_pos
		`, country_uuid, province_uuid, area_uuid, sub_area_uuid, start_date, end_date, country_uuid, province_uuid, area_uuid, sub_area_uuid, start_date, end_date).
		Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
		Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
		Joins("INNER JOIN communes ON pos_forms.commune_uuid = communes.uuid").
		Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.sub_area_uuid = ?", country_uuid, province_uuid, area_uuid, sub_area_uuid).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Group("communes.name, brands.name").
		Order("communes.name, total_count DESC").
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

// Line chart for sum brand by month
func NdTotalByBrandByMonth(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	year := c.Query("year")

	var results []struct {
		BrandName  string  `json:"brand_name"`
		Month      int     `json:"month"`
		TotalCount int     `json:"total_count"`
		Percentage float64 `json:"percentage"`
	}

	err := db.Table("pos_form_items").
		Select(`
		brands.name AS brand_name,
		EXTRACT(MONTH FROM pos_forms.created_at) AS month, 
		SUM(pos_form_items.counter) AS total_count,
       ROUND((SUM(pos_form_items.counter) / (SELECT SUM(counter) FROM pos_form_items INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid WHERE pos_forms.country_uuid = ? AND EXTRACT(YEAR FROM pos_forms.created_at) = ?)) * 100, 2) AS percentage
		`, country_uuid, year).
		Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
		Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
		Where("pos_forms.country_uuid = ? AND EXTRACT(YEAR FROM pos_forms.created_at) = ?", country_uuid, year).
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
		"message": "Total count by brand grouped by month for the year",
		"data":    results,
	})
}
