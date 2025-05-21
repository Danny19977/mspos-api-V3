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
		Name     string  `json:"name"`
		Brand    string  `json:"brand"`
		Presence int     `json:"presence"`
		Visits   int     `json:"visits"`
		Pourcent float64 `json:"pourcent"`
	}

	// err := db.Table("pos_form_items").
	// 	Select(`
	// 		provinces.name AS name,
	// 		brands.name AS brand,
	// 		COUNT(brands.name) AS presence,
	// 		(SELECT COUNT(pos_forms.uuid) FROM pos_forms
	// 		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
	// 		AND pos_forms.created_at BETWEEN ? AND ?
	// 		AND pos_forms.deleted_at IS NULL
	// 	"pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?", country_uuid, province_uuid).
	// 		) AS visits,
	// 		(COUNT(brands.name) * 100 / (SELECT COUNT(pos_forms.uuid) FROM pos_forms)) AS pourcent
	// 	`, country_uuid, province_uuid, start_date, end_date).
	// 	Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
	// 	Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
	// 	Joins("INNER JOIN provinces ON pos_forms.province_uuid = provinces.uuid").
	// 	Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
	// 	Where("pos_forms.deleted_at IS NULL").
	// 	Group("provinces.name, brands.name").
	// 	Order("pourcent DESC").
	// 	Scan(&results).Error

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"status":  "error",
	// 		"message": "Failed to fetch data",
	// 		"error":   err.Error(),
	// 	})
	// }

	sqlQuery := `
	   
		SELECT 
		provinces.name AS name,

		brands.name AS brand,

		COUNT(brands.name) AS presence,

		(SELECT COUNT(pos_forms.uuid) FROM pos_forms 
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		) AS visits,

		(COUNT(brands.name) * 100 / (
		SELECT COUNT(pos_forms.uuid) FROM pos_forms 
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		)) AS pourcent
		
		FROM pos_form_items 
		INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid
		INNER JOIN provinces ON pos_forms.province_uuid = provinces.uuid
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY provinces.name, brands.name
		ORDER BY pourcent DESC;
	`
	rows, err := db.Raw(sqlQuery, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date).Rows()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var name, brand string
		var presence, visits int
		var pourcent float64
		if err := rows.Scan(&name, &brand, &presence, &visits, &pourcent); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to scan data",
				"error":   err.Error(),
			})
		}
		results = append(results, struct {
			Name     string  `json:"name"`
			Brand    string  `json:"brand"`
			Presence int     `json:"presence"`
			Visits   int     `json:"visits"`
			Pourcent float64 `json:"pourcent"`
		}{
			Name:     name,
			Brand:    brand,
			Presence: presence,
			Visits:   visits,
			Pourcent: pourcent,
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
	area_uuid := c.Query("area_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name     string  `json:"name"`
		Brand    string  `json:"brand"`
		Presence int     `json:"presence"`
		Visits   int     `json:"visits"`
		Pourcent float64 `json:"pourcent"`
	}
	// err := db.Table("pos_form_items").
	// 	Select(`
	// 	provinces.name AS name,
	// 		brands.name AS brand,
	// 		COUNT(brands.name) AS presence,
	// 		(SELECT COUNT(pos_forms.uuid) FROM pos_forms) AS visits,
	// 		(COUNT(brands.name) * 100 / (SELECT COUNT(pos_forms.uuid) FROM pos_forms)) AS pourcent
	// 	`).
	// 	Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
	// 	Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
	// 	Joins("INNER JOIN areas ON pos_forms.area_uuid = areas.uuid").
	// 	Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?", country_uuid, province_uuid).
	// 	Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
	// 	Where("pos_forms.deleted_at IS NULL").
	// 	Group("areas.name, brands.name").
	// 	Order("pourcent DESC").
	// 	Scan(&results).Error

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"status":  "error",
	// 		"message": "Failed to fetch data",
	// 		"error":   err.Error(),
	// 	})
	// }
	sqlQuery := `
	   
		SELECT 
		provinces.name AS name,

		brands.name AS brand,

		COUNT(brands.name) AS presence,

		(SELECT COUNT(pos_forms.uuid) FROM pos_forms 
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		) AS visits,

		(COUNT(brands.name) * 100 / (
		SELECT COUNT(pos_forms.uuid) FROM pos_forms 
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		)) AS pourcent
		
		FROM pos_form_items 
		INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid
		INNER JOIN areas ON pos_forms.area_uuid = areas.uuid
		INNER JOIN provinces ON pos_forms.province_uuid = provinces.uuid
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY provinces.name, brands.name
		ORDER BY pourcent DESC;
	`

	rows, err := db.Raw(sqlQuery, country_uuid, province_uuid, area_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date).Rows()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var name, brand string
		var presence, visits int
		var pourcent float64
		if err := rows.Scan(&name, &brand, &presence, &visits, &pourcent); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to scan data",
				"error":   err.Error(),
			})
		}
		results = append(results, struct {
			Name     string  `json:"name"`
			Brand    string  `json:"brand"`
			Presence int     `json:"presence"`
			Visits   int     `json:"visits"`
			Pourcent float64 `json:"pourcent"`
		}{
			Name:     name,
			Brand:    brand,
			Presence: presence,
			Visits:   visits,
			Pourcent: pourcent,
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
	subarea_uuid := c.Query("sub_area_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name     string  `json:"name"`
		Brand    string  `json:"brand"`
		Presence int     `json:"presence"`
		Visits   int     `json:"visits"`
		Pourcent float64 `json:"pourcent"`
	}

	// err := db.Table("pos_form_items").
	// 	Select(`
	// 		provinces.name AS name,
	// 		brands.name AS brand,
	// 		COUNT(brands.name) AS presence,
	// 		(SELECT COUNT(pos_forms.uuid) FROM pos_forms) AS visits,
	// 		(COUNT(brands.name) * 100 / (SELECT COUNT(pos_forms.uuid) FROM pos_forms)) AS pourcent
	// 	`).
	// 	Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
	// 	Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
	// 	Joins("INNER JOIN sub_areas ON pos_forms.sub_area_uuid = sub_areas.uuid").
	// 	Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ?", country_uuid, province_uuid, area_uuid).
	// 	Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
	// 	Where("pos_forms.deleted_at IS NULL").
	// 	Group("sub_areas.name, brands.name").
	// 	Order("pourcent DESC").
	// 	Scan(&results).Error

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"status":  "error",
	// 		"message": "Failed to fetch data",
	// 		"error":   err.Error(),
	// 	})
	// }

	sqlQuery := `
	   
		SELECT 
		provinces.name AS name,

		brands.name AS brand,

		COUNT(brands.name) AS presence,

		(SELECT COUNT(pos_forms.uuid) FROM pos_forms 
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		) AS visits,

		(COUNT(brands.name) * 100 / (
		SELECT COUNT(pos_forms.uuid) FROM pos_forms 
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		)) AS pourcent
		
		FROM pos_form_items 
		INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid
		INNER JOIN areas ON pos_forms.area_uuid = areas.uuid
		INNER JOIN sub_areas ON pos_forms.sub_area_uuid = sub_areas.uuid
		INNER JOIN provinces ON pos_forms.province_uuid = provinces.uuid
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY provinces.name, brands.name
		ORDER BY pourcent DESC;
	`

	rows, err := db.Raw(sqlQuery, country_uuid, province_uuid, area_uuid, subarea_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date).Rows()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var name, brand string
		var presence, visits int
		var pourcent float64
		if err := rows.Scan(&name, &brand, &presence, &visits, &pourcent); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to scan data",
				"error":   err.Error(),
			})
		}
		results = append(results, struct {
			Name     string  `json:"name"`
			Brand    string  `json:"brand"`
			Presence int     `json:"presence"`
			Visits   int     `json:"visits"`
			Pourcent float64 `json:"pourcent"`
		}{
			Name:     name,
			Brand:    brand,
			Presence: presence,
			Visits:   visits,
			Pourcent: pourcent,
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
	commune_uuid := c.Query("commune_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name     string  `json:"name"`
		Brand    string  `json:"brand"`
		Presence int     `json:"presence"`
		Visits   int     `json:"visits"`
		Pourcent float64 `json:"pourcent"`
	}

	// err := db.Table("pos_form_items").
	// 	Select(`
	// 		provinces.name AS name,
	// 		brands.name AS brand,
	// 		COUNT(brands.name) AS presence,
	// 		(SELECT COUNT(pos_forms.uuid) FROM pos_forms) AS visits,
	// 		(COUNT(brands.name) * 100 / (SELECT COUNT(pos_forms.uuid) FROM pos_forms)) AS pourcent
	// 	`).
	// 	Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
	// 	Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
	// 	Joins("INNER JOIN communes ON pos_forms.commune_uuid = communes.uuid").
	// 	Where("pos_forms.country_uuid = ? AND pos_forms.province_uuid = ? AND pos_forms.area_uuid = ? AND pos_forms.sub_area_uuid = ?", country_uuid, province_uuid, area_uuid, sub_area_uuid).
	// 	Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
	// 	Where("pos_forms.deleted_at IS NULL").
	// 	Group("communes.name, brands.name").
	// 	Order("pourcent DESC").
	// 	Scan(&results).Error

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"status":  "error",
	// 		"message": "Failed to fetch data",
	// 		"error":   err.Error(),
	// 	})
	// }

	sqlQuery := `
	   
		SELECT 
		provinces.name AS name,

		brands.name AS brand,

		COUNT(brands.name) AS presence,

		(SELECT COUNT(pos_forms.uuid) FROM pos_forms 
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		) AS visits,

		(COUNT(brands.name) * 100 / (
		SELECT COUNT(pos_forms.uuid) FROM pos_forms 
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		)) AS pourcent
		
		FROM pos_form_items 
		INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid
		INNER JOIN areas ON pos_forms.area_uuid = areas.uuid
		INNER JOIN sub_areas ON pos_forms.sub_area_uuid = sub_areas.uuid
		INNER JOIN communes ON pos_forms.commune_uuid = communes.uuid
		INNER JOIN provinces ON pos_forms.province_uuid = provinces.uuid
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY provinces.name, brands.name
		ORDER BY pourcent DESC;
	`

	rows, err := db.Raw(sqlQuery, country_uuid, province_uuid, area_uuid, sub_area_uuid, commune_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date).Rows()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var name, brand string
		var presence, visits int
		var pourcent float64
		if err := rows.Scan(&name, &brand, &presence, &visits, &pourcent); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to scan data",
				"error":   err.Error(),
			})
		}
		results = append(results, struct {
			Name     string  `json:"name"`
			Brand    string  `json:"brand"`
			Presence int     `json:"presence"`
			Visits   int     `json:"visits"`
			Pourcent float64 `json:"pourcent"`
		}{
			Name:     name,
			Brand:    brand,
			Presence: presence,
			Visits:   visits,
			Pourcent: pourcent,
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
		Brand    string  `json:"brand"`
		Month    int     `json:"month"`
		Presence int     `json:"presence"`
		Pourcent float64 `json:"pourcent"`
	}

	// err := db.Table("pos_form_items").
	// 	Select(`
	// 	brands.name AS brand,
	// 	EXTRACT(MONTH FROM pos_forms.created_at) AS month,
	// 	SUM(pos_form_items.counter) AS presence,
	//    ROUND((SUM(pos_form_items.counter) / (SELECT SUM(counter) FROM pos_form_items
	//    INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid WHERE pos_forms.country_uuid = ?
	//    AND EXTRACT(YEAR FROM pos_forms.created_at) = ?)) * 100, 2) AS percentage
	// 	`, country_uuid, year).
	// 	Joins("INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid").
	// 	Joins("INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid").
	// 	Where("pos_forms.country_uuid = ? AND EXTRACT(YEAR FROM pos_forms.created_at) = ?", country_uuid, year).
	// 	Where("pos_forms.deleted_at IS NULL").
	// 	Group("brands.name, month").
	// 	Order("brands.name, month ASC").
	// 	Scan(&results).Error

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"status":  "error",
	// 		"message": "Failed to fetch data",
	// 		"error":   err.Error(),
	// 	})
	// }

	sqlQuery := `
		SELECT
			brands.name AS brand,
			EXTRACT(MONTH FROM pos_forms.created_at) AS month,
			COUNT(brands.name) AS presence,
			(COUNT(brands.name) * 100 / (
				SELECT COUNT(pos_forms.uuid) FROM pos_forms 
				WHERE pos_forms.country_uuid = ? 
				AND EXTRACT(YEAR FROM pos_forms.created_at) = ?
				AND pos_forms.deleted_at IS NULL
			)) AS pourcent
		FROM pos_form_items 
		INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid
		INNER JOIN provinces ON pos_forms.province_uuid = provinces.uuid
		WHERE pos_forms.country_uuid = ? AND EXTRACT(YEAR FROM pos_forms.created_at) = ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY brands.name, month
		ORDER BY brands.name, month ASC;
	`
	rows, err := db.Raw(sqlQuery, country_uuid, year, country_uuid, year).Rows()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}
	defer rows.Close()
	for rows.Next() {
		var brand string
		var month, presence int
		var pourcent float64
		if err := rows.Scan(&brand, &month, &presence, &pourcent); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to scan data",
				"error":   err.Error(),
			})
		}
		results = append(results, struct {
			Brand    string  `json:"brand"`
			Month    int     `json:"month"`
			Presence int     `json:"presence"`
			Pourcent float64 `json:"pourcent"`
		}{
			Brand:    brand,
			Month:    month,
			Presence: presence,
			Pourcent: pourcent,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Total count by brand grouped by month for the year",
		"data":    results,
	})
}
