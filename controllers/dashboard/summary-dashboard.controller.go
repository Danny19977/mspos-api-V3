package dashboard

import (
	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
)

func CycloCount(c *fiber.Ctx) error {
	sql1 := `
	 SELECT COUNT(*) FROM users WHERE "users"."deleted_at" IS NULL AND role='Cyclo' AND status=true;
	`
	var chartData models.SummaryCount
	database.DB.Raw(sql1).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func DrCount(c *fiber.Ctx) error {
	sql1 := `
	 SELECT COUNT(*) FROM users WHERE "users"."deleted_at" IS NULL AND role='DR' AND status=true;
	`
	var chartData models.SummaryCount
	database.DB.Raw(sql1).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func POSCount(c *fiber.Ctx) error {
	sql1 := `
	 SELECT COUNT(*) FROM pos WHERE "pos"."deleted_at" IS NULL AND status=true; 
	`
	var chartData models.SummaryCount
	database.DB.Raw(sql1).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func ProvinceCount(c *fiber.Ctx) error {
	sql1 := `
	 SELECT COUNT(*) FROM provinces WHERE "provinces"."deleted_at" IS NULL;
	`

	var chartData models.SummaryCount
	database.DB.Raw(sql1).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func AreaCount(c *fiber.Ctx) error {
	sql1 := `
	 SELECT COUNT(*) FROM areas WHERE "areas"."deleted_at" IS NULL;
	`

	var chartData models.SummaryCount
	database.DB.Raw(sql1).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func SOSPie(c *fiber.Ctx) error {
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
		SELECT "provinces"."name" AS province,
			ROUND(SUM(eq) / (SUM(eq) + SUM(dhl) + SUM(ar) +
			SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
			SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
			SUM(time) ) * 100) AS eq
		FROM pos_forms 
			INNER JOIN provinces ON pos_forms.province_uuid=provinces.id
				WHERE "pos_forms"."deleted_at" IS NULL AND "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP  
				AND ? ::TIMESTAMP 
		GROUP BY "provinces"."name"; 
	`

	var chartData []models.SosPieChart
	database.DB.Raw(sql1, start_date, end_date).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func TrackingVisitDRS(c *fiber.Ctx) error {
	days := c.Params("days")
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
		
	SELECT "provinces"."name" AS province,
		ROUND(SUM(eq1) / COUNT("pos_forms"."id") * 100) AS nd,
		ROUND(SUM(eq) / (SUM(eq) + SUM(dhl) + SUM(ar) +
		SUM(sbl) +
		SUM(pmf) +
		SUM(pmm) +
		SUM(ticket) +
		SUM(mtc) +
		SUM(ws) +
		SUM(mast) +
		SUM(oris) +
		SUM(elite) +
		SUM(yes) +
		SUM(time) ) * 100) AS sos,
	ROUND(100 - ROUND(SUM(eq1) / COUNT("pos_forms"."id") * 100)) AS oos,
	(SELECT COUNT(*) FROM users 
		INNER JOIN provinces ON users.province_uuid=provinces.id
		WHERE "users"."deleted_at" IS NULL AND role='DR' AND status=true AND province_uuid="provinces"."id") AS dr,
		
	COUNT("pos_forms"."id") AS visit,

	ROUND(40 * (SELECT COUNT(*) 
			FROM users 
			INNER JOIN provinces ON users.province_uuid=provinces.id
				WHERE  "users"."deleted_at" IS NULL AND role='DR' AND status=true AND province_uuid="provinces"."id") * 
	CASE 
		WHEN ? = 0 THEN 1
			ELSE ?
		END ) AS obj,
	
		ROUND(COUNT("pos_forms"."id") / (40 * (SELECT COUNT(*) 
			FROM users 
			INNER JOIN provinces ON users.province_uuid=provinces.id
				WHERE  "users"."deleted_at" IS NULL AND role='DR' AND status=true AND province_uuid="provinces"."id") * 
	CASE 
			WHEN ? = 0 THEN 1
			ELSE ?
		END )) AS perf
	
			FROM pos_forms 
					INNER JOIN provinces ON pos_forms.province_uuid=provinces.id
				INNER JOIN users ON pos_forms.user_id=users.id
					WHERE "pos_forms"."deleted_at" IS NULL AND "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP  
					AND ? ::TIMESTAMP

	GROUP BY "provinces"."name";
	`
	var chartData []models.TrackingVisitDRSChart
	database.DB.Raw(sql1, days, days, days, days, start_date, end_date).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func SummaryChartBar(c *fiber.Ctx) error {
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
		SELECT "provinces"."name" AS province,
			ROUND(SUM(eq1) / COUNT("pos_forms"."id") * 100) AS nd,
			ROUND(SUM(eq) / (SUM(eq) + SUM(dhl) + SUM(ar) +
			SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) +
			SUM(mtc) + SUM(ws) + SUM(mast) +
			SUM(oris) + SUM(elite) + SUM(yes) + SUM(time) 
		) * 100) AS sos,
		ROUND(100 - ROUND(SUM(eq1) / COUNT("pos_forms"."id") * 100)) AS oos
		FROM pos_forms 
				INNER JOIN provinces ON pos_forms.province_uuid=provinces.id
			INNER JOIN users ON pos_forms.user_id=users.id
				WHERE "pos_forms"."deleted_at" IS NULL AND "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP  
				AND ? ::TIMESTAMP

		GROUP BY "provinces"."name";
	`
	var chartData []models.SumChartBar
	database.DB.Raw(sql1, start_date, end_date).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func BetterDR(c *fiber.Ctx) error {
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
	SELECT fullname, 
		"provinces"."name" AS province, 
		"areas"."name" AS area,
	SUM(sold) AS ventes
	FROM pos_forms
	INNER JOIN users ON pos_forms.user_id=users.id
	INNER JOIN provinces ON pos_forms.province_uuid=provinces.id
	INNER JOIN areas ON pos_forms.province_uuid=areas.id
	WHERE "pos_forms"."deleted_at" IS NULL AND "users"."role"='DR' AND "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP 
			AND ? ::TIMESTAMP
	GROUP BY fullname, "provinces"."name", "areas"."name"
	ORDER BY ventes DESC
	LIMIT 10; 
	`
	var chartData []models.SummaryBetterDR
	database.DB.Raw(sql1, start_date, end_date).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func BetterSup(c *fiber.Ctx) error {
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
		SELECT fullname, 
			"provinces"."name" AS province, 
			"areas"."name" AS area,
		SUM(sold) AS ventes
		FROM pos_forms
		INNER JOIN users ON pos_forms.user_id=users.id
		INNER JOIN provinces ON pos_forms.province_uuid=provinces.id
		INNER JOIN areas ON pos_forms.province_uuid=areas.id
		WHERE "pos_forms"."deleted_at" IS NULL AND "users"."role"='Supervisor' AND "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP 
				AND ? ::TIMESTAMP
		GROUP BY fullname, "provinces"."name", "areas"."name"
		ORDER BY ventes DESC
		LIMIT 5;
	`
	var chartData []models.SummaryBetterDR
	database.DB.Raw(sql1, start_date, end_date).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func StatusEquipement(c *fiber.Ctx) error {
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
		SELECT input_group_selector AS equipement,  
		COUNT(input_group_selector) AS count
		FROM pos_forms
		INNER JOIN pos ON pos_forms.pos_id=pos.id 
		WHERE "pos_forms"."deleted_at" IS NULL AND "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP 
				AND ? ::TIMESTAMP
		GROUP BY input_group_selector;
	`

	var chartData []models.StatusEquip
	database.DB.Raw(sql1, start_date, end_date).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func GoogleMaps(c *fiber.Ctx) error {
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
		SELECT  
			pos_forms.latitude AS latitude,
			pos_forms.longitude AS longitude,
			users.fullname AS name
		FROM pos_forms
		INNER JOIN users ON pos_forms.user_id=users.id
		WHERE "pos_forms"."deleted_at" IS NULL AND latitude::FLOAT != 0 AND longitude::FLOAT != 0 AND
		"pos_forms"."created_at" BETWEEN ? ::TIMESTAMP AND ? ::TIMESTAMP;
	`
	var chartData []models.GoogleMap
	database.DB.Raw(sql1, start_date, end_date).Scan(&chartData)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func PriceSale(c *fiber.Ctx) error {
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
		SELECT price AS price,
		COUNT(*)
		FROM pos_forms  
		WHERE "pos_forms"."deleted_at" IS NULL AND created_at BETWEEN ? ::TIMESTAMP AND ? ::TIMESTAMP
		GROUP BY price;
	`
	var chartData []models.PriceChart
	database.DB.Raw(sql1, start_date, end_date).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}
