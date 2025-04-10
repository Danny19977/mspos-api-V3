package dashboard

import (
	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
)

func NdTableView(c *fiber.Ctx) error {
	province := c.Params("province")
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
		SELECT areas.name AS area, 
			ROUND(SUM(eq1) / COUNT("pos_forms"."id") * 100) AS eq,
			ROUND(SUM(dhl1) / COUNT("pos_forms"."id") * 100) AS dhl, 
			ROUND(SUM(ar1) / COUNT("pos_forms"."id") * 100) AS ar, 
			ROUND(SUM(sbl1)/ COUNT("pos_forms"."id") * 100) AS sbl, 
			ROUND(SUM(pmf1) / COUNT("pos_forms"."id") * 100) AS pmf,
			ROUND(SUM(pmm1) / COUNT("pos_forms"."id") * 100) AS pmm, 
			ROUND(SUM(ticket1) / COUNT("pos_forms"."id") * 100) AS ticket, 
			ROUND(SUM(mtc1) / COUNT("pos_forms"."id") * 100) AS mtc, 
			ROUND(SUM(ws1) / COUNT("pos_forms"."id") * 100) AS ws, 
			ROUND(SUM(mast1) / COUNT("pos_forms"."id") * 100) AS mast,
			ROUND(SUM(oris1) / COUNT("pos_forms"."id") * 100) AS oris, 
			ROUND(SUM(elite1) / COUNT("pos_forms"."id") * 100) AS elite,
			ROUND(SUM(yes1) / COUNT("pos_forms"."id") * 100) AS yes, 
			ROUND(SUM(time1) / COUNT("pos_forms"."id") * 100) AS time
		FROM pos_forms
		INNER JOIN areas ON pos_forms.area_uuid=areas.id
		INNER JOIN provinces ON pos_forms.province_uuid=provinces.id
		WHERE "pos_forms"."deleted_at" IS NULL AND "provinces"."name"= ? AND "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP 
			AND ? ::TIMESTAMP 
		GROUP BY areas.name;
	`
	var chartData []models.NDChartData
	database.DB.Raw(sql1, province, start_date, end_date).Scan(&chartData)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func NdByYear(c *fiber.Ctx) error {
	province := c.Params("province")
	sql1 := `  
	SELECT EXTRACT(MONTH FROM "pos_forms"."created_at") AS month,
		ROUND(SUM(eq1) / COUNT(*) * 100) AS eq
	FROM pos_forms
	INNER JOIN provinces ON pos_forms.province_uuid=provinces.id
	WHERE "pos_forms"."deleted_at" IS NULL AND "provinces"."name"=? AND 
    EXTRACT(YEAR FROM "pos_forms"."created_at") = EXTRACT(YEAR FROM CURRENT_DATE)
		AND EXTRACT(MONTH FROM "pos_forms"."created_at") BETWEEN 1 AND 12 
    
		GROUP BY EXTRACT(MONTH FROM "pos_forms"."created_at")
		ORDER BY EXTRACT(MONTH FROM "pos_forms"."created_at");
	`
	var chartData []models.NdByYear
	database.DB.Raw(sql1, province).Scan(&chartData)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}
