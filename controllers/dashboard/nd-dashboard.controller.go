package dashboard

import (
	"fmt"

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
		UUID     string  `json:"uuid"`
		Brand    string  `json:"brand"`
		Presence int     `json:"presence"`
		Visits   int     `json:"visits"`
		Pourcent float64 `json:"pourcent"`
	}

	sqlQuery := `
		SELECT 
			provinces.name AS name,
			provinces.uuid AS uuid,
			brands.name AS brand,
			COUNT(DISTINCT pos_form_items.uuid) AS presence,
			(SELECT COUNT(DISTINCT pos_forms.uuid) 
			 FROM pos_forms 
			 WHERE pos_forms.country_uuid = ? 
			 AND pos_forms.province_uuid = ?
			 AND pos_forms.created_at BETWEEN ? AND ?
			 AND pos_forms.deleted_at IS NULL
			) AS visits,
			CASE 
				WHEN (SELECT COUNT(DISTINCT pos_form_items.uuid) 
					  FROM pos_form_items
					  INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
					  WHERE pos_forms.country_uuid = ? 
					  AND pos_forms.province_uuid = ?
					  AND pos_forms.created_at BETWEEN ? AND ?
					  AND pos_forms.deleted_at IS NULL) > 0 
				THEN (COUNT(DISTINCT pos_form_items.uuid) * 100.0 / 
					  (SELECT COUNT(DISTINCT pos_form_items.uuid) 
					   FROM pos_form_items
					   INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
					   WHERE pos_forms.country_uuid = ? 
					   AND pos_forms.province_uuid = ?
					   AND pos_forms.created_at BETWEEN ? AND ?
					   AND pos_forms.deleted_at IS NULL))
				ELSE 0
			END AS pourcent
		FROM pos_form_items 
		INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid
		INNER JOIN provinces ON pos_forms.province_uuid = provinces.uuid
		WHERE pos_forms.country_uuid = ? AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY provinces.name, provinces.uuid, brands.name
		ORDER BY pourcent DESC;
	`
	rows, err := db.Raw(sqlQuery, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date).Rows()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var name, uuid, brand string
		var presence, visits int
		var pourcent float64
		if err := rows.Scan(&name, &uuid, &brand, &presence, &visits, &pourcent); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to scan data",
				"error":   err.Error(),
			})
		}
		results = append(results, struct {
			Name     string  `json:"name"`
			UUID     string  `json:"uuid"`
			Brand    string  `json:"brand"`
			Presence int     `json:"presence"`
			Visits   int     `json:"visits"`
			Pourcent float64 `json:"pourcent"`
		}{
			Name:     name,
			UUID:     uuid,
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
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name     string  `json:"name"`
		UUID     string  `json:"uuid"`
		Brand    string  `json:"brand"`
		Presence int     `json:"presence"`
		Visits   int     `json:"visits"`
		Pourcent float64 `json:"pourcent"`
	}

	sqlQuery := `
		SELECT  
			areas.name AS name,
			areas.uuid AS uuid,
			brands.name AS brand,
			COUNT(DISTINCT pos_form_items.uuid) AS presence,
			(SELECT COUNT(DISTINCT pos_forms.uuid) 
			 FROM pos_forms 
			 WHERE pos_forms.country_uuid = ? 
			 AND pos_forms.province_uuid = ? 
			 AND pos_forms.created_at BETWEEN ? AND ?
			 AND pos_forms.deleted_at IS NULL
			) AS visits,
			CASE 
				WHEN (SELECT COUNT(DISTINCT pos_form_items.uuid) 
					  FROM pos_form_items
					  INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
					  WHERE pos_forms.country_uuid = ? 
					  AND pos_forms.province_uuid = ?
					  AND pos_forms.created_at BETWEEN ? AND ?
					  AND pos_forms.deleted_at IS NULL) > 0 
				THEN (COUNT(DISTINCT pos_form_items.uuid) * 100.0 / 
					  (SELECT COUNT(DISTINCT pos_form_items.uuid) 
					   FROM pos_form_items
					   INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
					   WHERE pos_forms.country_uuid = ? 
					   AND pos_forms.province_uuid = ?
					   AND pos_forms.created_at BETWEEN ? AND ?
					   AND pos_forms.deleted_at IS NULL))
				ELSE 0
			END AS pourcent
		FROM pos_form_items
		INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid
		INNER JOIN areas ON pos_forms.area_uuid = areas.uuid
		WHERE pos_forms.country_uuid = ?
		AND pos_forms.province_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY areas.name, areas.uuid, brands.name
		ORDER BY pourcent DESC;
	`
	rows, err := db.Raw(sqlQuery, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date, country_uuid, province_uuid, start_date, end_date).Rows()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var name, uuid, brand string
		var presence, visits int
		var pourcent float64
		if err := rows.Scan(&name, &uuid, &brand, &presence, &visits, &pourcent); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to scan data",
				"error":   err.Error(),
			})
		}
		results = append(results, struct {
			Name     string  `json:"name"`
			UUID     string  `json:"uuid"`
			Brand    string  `json:"brand"`
			Presence int     `json:"presence"`
			Visits   int     `json:"visits"`
			Pourcent float64 `json:"pourcent"`
		}{
			Name:     name,
			UUID:     uuid,
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
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name     string  `json:"name"`
		UUID     string  `json:"uuid"`
		Brand    string  `json:"brand"`
		Presence int     `json:"presence"`
		Visits   int     `json:"visits"`
		Pourcent float64 `json:"pourcent"`
	}

	sqlQuery := `
		SELECT  
			sub_areas.name AS name,
			sub_areas.uuid AS uuid,
			brands.name AS brand,
			COUNT(DISTINCT pos_form_items.uuid) AS presence,
			(SELECT COUNT(DISTINCT pos_forms.uuid) 
			 FROM pos_forms 
			 WHERE pos_forms.country_uuid = ? 
			 AND pos_forms.province_uuid = ? 
			 AND pos_forms.area_uuid = ?
			 AND pos_forms.created_at BETWEEN ? AND ?
			 AND pos_forms.deleted_at IS NULL
			) AS visits,
			CASE 
				WHEN (SELECT COUNT(DISTINCT pos_forms.uuid) 
					  FROM pos_forms 
					  WHERE pos_forms.country_uuid = ? 
					  AND pos_forms.province_uuid = ? 
					  AND pos_forms.area_uuid = ?
					  AND pos_forms.created_at BETWEEN ? AND ?
					  AND pos_forms.deleted_at IS NULL) > 0 
				THEN (COUNT(DISTINCT pos_form_items.uuid) * 100.0 / 
					  (SELECT COUNT(DISTINCT pos_forms.uuid) 
					   FROM pos_forms 
					   WHERE pos_forms.country_uuid = ? 
					   AND pos_forms.province_uuid = ? 
					   AND pos_forms.area_uuid = ?
					   AND pos_forms.created_at BETWEEN ? AND ?
					   AND pos_forms.deleted_at IS NULL))
				ELSE 0
			END AS pourcent
		FROM pos_form_items 
		INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid 
		INNER JOIN sub_areas ON pos_forms.sub_area_uuid = sub_areas.uuid 
		WHERE pos_forms.country_uuid = ? 
		AND pos_forms.province_uuid = ? 
		AND pos_forms.area_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY sub_areas.name, sub_areas.uuid, brands.name
		ORDER BY pourcent DESC;
	`

	rows, err := db.Raw(sqlQuery,
		country_uuid, province_uuid, area_uuid, start_date, end_date,
		country_uuid, province_uuid, area_uuid, start_date, end_date,
		country_uuid, province_uuid, area_uuid, start_date, end_date,
		country_uuid, province_uuid, area_uuid, start_date, end_date).Rows()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var name, uuid, brand string
		var presence, visits int
		var pourcent float64
		if err := rows.Scan(&name, &uuid, &brand, &presence, &visits, &pourcent); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to scan data",
				"error":   err.Error(),
			})
		}

		results = append(results, struct {
			Name     string  `json:"name"`
			UUID     string  `json:"uuid"`
			Brand    string  `json:"brand"`
			Presence int     `json:"presence"`
			Visits   int     `json:"visits"`
			Pourcent float64 `json:"pourcent"`
		}{
			Name:     name,
			UUID:     uuid,
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
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var results []struct {
		Name     string  `json:"name"`
		UUID     string  `json:"uuid"`
		Brand    string  `json:"brand"`
		Presence int     `json:"presence"`
		Visits   int     `json:"visits"`
		Pourcent float64 `json:"pourcent"`
	}

	fmt.Println("country_uuid:", country_uuid)
	fmt.Println("province_uuid:", province_uuid)
	fmt.Println("area_uuid:", area_uuid)
	fmt.Println("sub_area_uuid:", sub_area_uuid)
	fmt.Println("start_date:", start_date)
	fmt.Println("end_date:", end_date)

	sqlQuery := `
		SELECT  
			communes.name AS name,
			communes.uuid AS uuid,
			brands.name AS brand,
			COUNT(DISTINCT pos_form_items.uuid) AS presence,
			(SELECT COUNT(DISTINCT pos_forms.uuid) 
			 FROM pos_forms 
			 WHERE pos_forms.country_uuid = ? 
			 AND pos_forms.province_uuid = ? 
			 AND pos_forms.area_uuid = ? 
			 AND pos_forms.sub_area_uuid = ?
			 AND pos_forms.created_at BETWEEN ? AND ?
			 AND pos_forms.deleted_at IS NULL
			) AS visits,
			CASE 
				WHEN (SELECT COUNT(DISTINCT pos_forms.uuid) 
					  FROM pos_forms 
					  WHERE pos_forms.country_uuid = ? 
					  AND pos_forms.province_uuid = ? 
					  AND pos_forms.area_uuid = ? 
					  AND pos_forms.sub_area_uuid = ?
					  AND pos_forms.created_at BETWEEN ? AND ?
					  AND pos_forms.deleted_at IS NULL) > 0 
				THEN (COUNT(DISTINCT pos_form_items.uuid) * 100.0 / 
					  (SELECT COUNT(DISTINCT pos_forms.uuid) 
					   FROM pos_forms 
					   WHERE pos_forms.country_uuid = ? 
					   AND pos_forms.province_uuid = ? 
					   AND pos_forms.area_uuid = ? 
					   AND pos_forms.sub_area_uuid = ?
					   AND pos_forms.created_at BETWEEN ? AND ?
					   AND pos_forms.deleted_at IS NULL))
				ELSE 0
			END AS pourcent
		FROM pos_form_items 
		INNER JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		INNER JOIN brands ON pos_form_items.brand_uuid = brands.uuid 
		INNER JOIN communes ON pos_forms.commune_uuid = communes.uuid 
		WHERE pos_forms.country_uuid = ? 
		AND pos_forms.province_uuid = ? 
		AND pos_forms.area_uuid = ? 
		AND pos_forms.sub_area_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY communes.name, communes.uuid, brands.name
		ORDER BY pourcent DESC;
	`

	rows, err := db.Raw(sqlQuery,
		country_uuid, province_uuid, area_uuid, sub_area_uuid, start_date, end_date,
		country_uuid, province_uuid, area_uuid, sub_area_uuid, start_date, end_date,
		country_uuid, province_uuid, area_uuid, sub_area_uuid, start_date, end_date,
		country_uuid, province_uuid, area_uuid, sub_area_uuid, start_date, end_date).Rows()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch data",
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var name, uuid, brand string
		var presence, visits int
		var pourcent float64
		if err := rows.Scan(&name, &uuid, &brand, &presence, &visits, &pourcent); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to scan data",
				"error":   err.Error(),
			})
		}
		results = append(results, struct {
			Name     string  `json:"name"`
			UUID     string  `json:"uuid"`
			Brand    string  `json:"brand"`
			Presence int     `json:"presence"`
			Visits   int     `json:"visits"`
			Pourcent float64 `json:"pourcent"`
		}{
			Name:     name,
			UUID:     uuid,
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
