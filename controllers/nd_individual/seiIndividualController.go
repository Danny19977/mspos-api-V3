package ndindividual

import (
	"github.com/danny19977/mspos-api-v3/database"
	"github.com/gofiber/fiber/v2"
)

// ╔══════════════════════════════════════════════════════════════════════════╗
// ║           SALES EVOLUTION INDIVIDUEL — PAR AGENT                        ║
// ╠══════════════════════════════════════════════════════════════════════════╣
// ║  Chaque agent peut consulter l'évolution de ses ventes (fardes,         ║
// ║  vendus, prix) par marque, par type de POS et par période.              ║
// ╠══════════════════════════════════════════════════════════════════════════╣
// ║  GET /sales-evolution-individual/summary-kpi/:user_uuid                 ║
// ║  GET /sales-evolution-individual/by-pos-type/:user_uuid                 ║
// ║  GET /sales-evolution-individual/price-by-brand/:user_uuid              ║
// ║  GET /sales-evolution-individual/evolution-by-month/:user_uuid          ║
// ║  GET /sales-evolution-individual/growth-rate/:user_uuid                 ║
// ║  GET /sales-evolution-individual/brand-competition/:user_uuid           ║
// ║  GET /sales-evolution-individual/top-pos/:user_uuid                     ║
// ║  GET /sales-evolution-individual/heatmap-day-of-week/:user_uuid         ║
// ║  GET /sales-evolution-individual/price-pie-chart/:user_uuid             ║
// ╚══════════════════════════════════════════════════════════════════════════╝

// ─────────────────────────────────────────────────────────────────────────────
// 1. SUMMARY KPI — chiffres globaux de l'agent pour la période
// ─────────────────────────────────────────────────────────────────────────────

func GetSEISummaryKPI(c *fiber.Ctx) error {
	db := database.DB

	userUUID := c.Params("user_uuid")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if userUUID == "" || startDate == "" || endDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "user_uuid, start_date et end_date sont obligatoires",
		})
	}

	type SummaryRow struct {
		UserUUID      string  `json:"user_uuid"`
		Fullname      string  `json:"fullname"`
		TotalFarde    float64 `json:"total_farde"`
		TotalSold     float64 `json:"total_sold"`
		TotalVisits   int64   `json:"total_visits"`
		ActivePos     int64   `json:"active_pos"`
		BrandsCovered int64   `json:"brands_covered"`
		AvgPrice      float64 `json:"avg_price"`
		ActiveDays    int64   `json:"active_days"`
	}

	sqlQuery := `
		SELECT
			u.uuid                                                   AS user_uuid,
			u.fullname                                               AS fullname,
			ROUND(COALESCE(SUM(pfi.number_farde), 0)::numeric, 2)   AS total_farde,
			ROUND(COALESCE(SUM(pfi.sold), 0)::numeric, 2)            AS total_sold,
			COUNT(DISTINCT pf.uuid)                                  AS total_visits,
			COUNT(DISTINCT pf.pos_uuid)                              AS active_pos,
			COUNT(DISTINCT pfi.brand_uuid)                           AS brands_covered,
			ROUND(COALESCE(AVG(pf.price), 0)::numeric, 2)            AS avg_price,
			COUNT(DISTINCT DATE(pf.created_at))                      AS active_days
		FROM users u
		LEFT JOIN pos_forms pf ON pf.user_uuid = u.uuid
			AND pf.created_at BETWEEN @start_date AND @end_date
			AND pf.deleted_at IS NULL
		LEFT JOIN pos_form_items pfi ON pfi.pos_form_uuid = pf.uuid
			AND pfi.deleted_at IS NULL
		WHERE u.uuid = @user_uuid AND u.deleted_at IS NULL
		GROUP BY u.uuid, u.fullname
	`

	var result SummaryRow
	if err := db.Raw(sqlQuery, map[string]interface{}{
		"user_uuid":  userUUID,
		"start_date": startDate,
		"end_date":   endDate,
	}).Scan(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   result,
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// 2. BY POS TYPE — répartition des ventes par type de POS
// ─────────────────────────────────────────────────────────────────────────────

func GetSEIByPosType(c *fiber.Ctx) error {
	db := database.DB

	userUUID := c.Params("user_uuid")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if userUUID == "" || startDate == "" || endDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "user_uuid, start_date et end_date sont obligatoires",
		})
	}

	type PosTypeRow struct {
		PosType          string  `json:"pos_type"`
		TotalVisits      int64   `json:"total_visits"`
		TotalPos         int64   `json:"total_pos"`
		TotalFarde       float64 `json:"total_farde"`
		TotalSold        float64 `json:"total_sold"`
		AvgFardePerVisit float64 `json:"avg_farde_per_visit"`
		AvgSoldPerVisit  float64 `json:"avg_sold_per_visit"`
		MarketShareFarde float64 `json:"market_share_farde"`
		MarketShareSold  float64 `json:"market_share_sold"`
	}

	sqlQuery := `
		WITH global AS (
			SELECT
				COALESCE(SUM(pfi.number_farde), 0) AS g_farde,
				COALESCE(SUM(pfi.sold), 0)          AS g_sold
			FROM pos_form_items pfi
			INNER JOIN pos_forms pf ON pfi.pos_form_uuid = pf.uuid
			WHERE pf.user_uuid = @user_uuid
			  AND pf.created_at BETWEEN @start_date AND @end_date
			  AND pf.deleted_at IS NULL AND pfi.deleted_at IS NULL
		)
		SELECT
			COALESCE(NULLIF(p.postype, ''), 'Non défini')            AS pos_type,
			COUNT(DISTINCT pf.uuid)                                   AS total_visits,
			COUNT(DISTINCT pf.pos_uuid)                               AS total_pos,
			ROUND(SUM(pfi.number_farde)::numeric, 2)                  AS total_farde,
			ROUND(SUM(pfi.sold)::numeric, 2)                           AS total_sold,
			ROUND((SUM(pfi.number_farde) / NULLIF(COUNT(DISTINCT pf.uuid), 0))::numeric, 2) AS avg_farde_per_visit,
			ROUND((SUM(pfi.sold)         / NULLIF(COUNT(DISTINCT pf.uuid), 0))::numeric, 2) AS avg_sold_per_visit,
			ROUND((SUM(pfi.number_farde) * 100.0 / NULLIF((SELECT g_farde FROM global), 0))::numeric, 2) AS market_share_farde,
			ROUND((SUM(pfi.sold)         * 100.0 / NULLIF((SELECT g_sold  FROM global), 0))::numeric, 2) AS market_share_sold
		FROM pos_form_items pfi
		INNER JOIN pos_forms pf ON pfi.pos_form_uuid = pf.uuid
		INNER JOIN pos p        ON pf.pos_uuid = p.uuid
		WHERE pf.user_uuid = @user_uuid
		  AND pf.created_at BETWEEN @start_date AND @end_date
		  AND pf.deleted_at IS NULL AND pfi.deleted_at IS NULL
		GROUP BY p.postype
		ORDER BY total_farde DESC
	`

	var rows []PosTypeRow
	if err := db.Raw(sqlQuery, map[string]interface{}{
		"user_uuid":  userUUID,
		"start_date": startDate,
		"end_date":   endDate,
	}).Scan(&rows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   rows,
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// 3. PRICE BY BRAND — prix moyen / min / max et volume par marque
// ─────────────────────────────────────────────────────────────────────────────

func GetSEIPriceByBrand(c *fiber.Ctx) error {
	db := database.DB

	userUUID := c.Params("user_uuid")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if userUUID == "" || startDate == "" || endDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "user_uuid, start_date et end_date sont obligatoires",
		})
	}

	type PriceBrandRow struct {
		BrandUUID    string  `json:"brand_uuid"`
		BrandName    string  `json:"brand_name"`
		TotalVisits  int64   `json:"total_visits"`
		TotalPos     int64   `json:"total_pos"`
		AvgPrice     float64 `json:"avg_price"`
		MinPrice     float64 `json:"min_price"`
		MaxPrice     float64 `json:"max_price"`
		TotalFarde   float64 `json:"total_farde"`
		TotalSold    float64 `json:"total_sold"`
		RevenueShare float64 `json:"revenue_share"`
	}

	sqlQuery := `
		WITH global_rev AS (
			SELECT COALESCE(SUM(pf.price), 0) AS g_rev
			FROM pos_forms pf
			WHERE pf.user_uuid = @user_uuid
			  AND pf.created_at BETWEEN @start_date AND @end_date
			  AND pf.deleted_at IS NULL
		)
		SELECT
			b.uuid                                                         AS brand_uuid,
			b.name                                                         AS brand_name,
			COUNT(DISTINCT pf.uuid)                                        AS total_visits,
			COUNT(DISTINCT pf.pos_uuid)                                    AS total_pos,
			ROUND(AVG(pf.price)::numeric, 2)                               AS avg_price,
			ROUND(MIN(pf.price)::numeric, 2)                               AS min_price,
			ROUND(MAX(pf.price)::numeric, 2)                               AS max_price,
			ROUND(SUM(pfi.number_farde)::numeric, 2)                       AS total_farde,
			ROUND(SUM(pfi.sold)::numeric, 2)                                AS total_sold,
			ROUND((SUM(pf.price) * 100.0 / NULLIF((SELECT g_rev FROM global_rev), 0))::numeric, 2) AS revenue_share
		FROM pos_form_items pfi
		INNER JOIN pos_forms pf ON pfi.pos_form_uuid = pf.uuid
		INNER JOIN brands b     ON pfi.brand_uuid = b.uuid
		WHERE pf.user_uuid = @user_uuid
		  AND pf.created_at BETWEEN @start_date AND @end_date
		  AND pf.deleted_at IS NULL AND pfi.deleted_at IS NULL
		GROUP BY b.uuid, b.name
		ORDER BY total_farde DESC
	`

	var rows []PriceBrandRow
	if err := db.Raw(sqlQuery, map[string]interface{}{
		"user_uuid":  userUUID,
		"start_date": startDate,
		"end_date":   endDate,
	}).Scan(&rows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   rows,
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// 4. EVOLUTION BY MONTH — tendance mensuelle des fardes/vendus par marque
// ─────────────────────────────────────────────────────────────────────────────

func GetSEIEvolutionByMonth(c *fiber.Ctx) error {
	db := database.DB

	userUUID := c.Params("user_uuid")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	brandUUID := c.Query("brand_uuid")

	if userUUID == "" || startDate == "" || endDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "user_uuid, start_date et end_date sont obligatoires",
		})
	}

	type MonthlyRow struct {
		YearMonth   string  `json:"year_month"`
		BrandName   string  `json:"brand_name"`
		TotalVisits int64   `json:"total_visits"`
		TotalPos    int64   `json:"total_pos"`
		TotalFarde  float64 `json:"total_farde"`
		TotalSold   float64 `json:"total_sold"`
		GrowthFarde float64 `json:"growth_farde_pct"`
		GrowthSold  float64 `json:"growth_sold_pct"`
	}

	brandFilter := ""
	params := map[string]interface{}{
		"user_uuid":  userUUID,
		"start_date": startDate,
		"end_date":   endDate,
	}
	if brandUUID != "" {
		brandFilter = " AND pfi.brand_uuid = @brand_uuid"
		params["brand_uuid"] = brandUUID
	}

	sqlQuery := `
		WITH monthly AS (
			SELECT
				TO_CHAR(pf.created_at, 'YYYY-MM')            AS year_month,
				b.name                                         AS brand_name,
				COUNT(DISTINCT pf.uuid)                        AS total_visits,
				COUNT(DISTINCT pf.pos_uuid)                    AS total_pos,
				ROUND(SUM(pfi.number_farde)::numeric, 2)       AS total_farde,
				ROUND(SUM(pfi.sold)::numeric, 2)                AS total_sold
			FROM pos_form_items pfi
			INNER JOIN pos_forms pf ON pfi.pos_form_uuid = pf.uuid
			INNER JOIN brands b     ON pfi.brand_uuid = b.uuid
			WHERE pf.user_uuid = @user_uuid
			  AND pf.created_at BETWEEN @start_date AND @end_date
			  AND pf.deleted_at IS NULL AND pfi.deleted_at IS NULL` + brandFilter + `
			GROUP BY year_month, b.name
		)
		SELECT
			year_month,
			brand_name,
			total_visits,
			total_pos,
			total_farde,
			total_sold,
			ROUND(((total_farde - LAG(total_farde) OVER (PARTITION BY brand_name ORDER BY year_month))
				* 100.0 / NULLIF(LAG(total_farde) OVER (PARTITION BY brand_name ORDER BY year_month), 0))::numeric, 2) AS growth_farde_pct,
			ROUND(((total_sold  - LAG(total_sold)  OVER (PARTITION BY brand_name ORDER BY year_month))
				* 100.0 / NULLIF(LAG(total_sold)  OVER (PARTITION BY brand_name ORDER BY year_month), 0))::numeric, 2)  AS growth_sold_pct
		FROM monthly
		ORDER BY brand_name, year_month
	`

	var rows []MonthlyRow
	if err := db.Raw(sqlQuery, params).Scan(&rows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   rows,
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// 5. GROWTH RATE — comparaison periode courante vs période précédente
// ─────────────────────────────────────────────────────────────────────────────

func GetSEIGrowthRate(c *fiber.Ctx) error {
	db := database.DB

	userUUID := c.Params("user_uuid")
	currStart := c.Query("curr_start")
	currEnd := c.Query("curr_end")
	prevStart := c.Query("prev_start")
	prevEnd := c.Query("prev_end")

	if userUUID == "" || currStart == "" || currEnd == "" || prevStart == "" || prevEnd == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "user_uuid, curr_start, curr_end, prev_start, prev_end sont obligatoires",
		})
	}

	type GrowthRow struct {
		BrandName      string  `json:"brand_name"`
		CurrFarde      float64 `json:"curr_farde"`
		PrevFarde      float64 `json:"prev_farde"`
		DeltaFarde     float64 `json:"delta_farde"`
		GrowthFardePct float64 `json:"growth_farde_pct"`
		CurrSold       float64 `json:"curr_sold"`
		PrevSold       float64 `json:"prev_sold"`
		DeltaSold      float64 `json:"delta_sold"`
		GrowthSoldPct  float64 `json:"growth_sold_pct"`
		CurrVisits     int64   `json:"curr_visits"`
		PrevVisits     int64   `json:"prev_visits"`
		Trend          string  `json:"trend"`
	}

	sqlQuery := `
		WITH curr AS (
			SELECT
				b.name                                    AS brand_name,
				ROUND(SUM(pfi.number_farde)::numeric, 2)  AS farde,
				ROUND(SUM(pfi.sold)::numeric, 2)           AS sold,
				COUNT(DISTINCT pf.uuid)                    AS visits
			FROM pos_form_items pfi
			INNER JOIN pos_forms pf ON pfi.pos_form_uuid = pf.uuid
			INNER JOIN brands b     ON pfi.brand_uuid = b.uuid
			WHERE pf.user_uuid = @user_uuid
			  AND pf.created_at BETWEEN @curr_start AND @curr_end
			  AND pf.deleted_at IS NULL AND pfi.deleted_at IS NULL
			GROUP BY b.name
		),
		prev AS (
			SELECT
				b.name                                    AS brand_name,
				ROUND(SUM(pfi.number_farde)::numeric, 2)  AS farde,
				ROUND(SUM(pfi.sold)::numeric, 2)           AS sold,
				COUNT(DISTINCT pf.uuid)                    AS visits
			FROM pos_form_items pfi
			INNER JOIN pos_forms pf ON pfi.pos_form_uuid = pf.uuid
			INNER JOIN brands b     ON pfi.brand_uuid = b.uuid
			WHERE pf.user_uuid = @user_uuid
			  AND pf.created_at BETWEEN @prev_start AND @prev_end
			  AND pf.deleted_at IS NULL AND pfi.deleted_at IS NULL
			GROUP BY b.name
		)
		SELECT
			COALESCE(c.brand_name, p.brand_name)                                AS brand_name,
			COALESCE(c.farde, 0)                                                 AS curr_farde,
			COALESCE(p.farde, 0)                                                 AS prev_farde,
			ROUND((COALESCE(c.farde, 0) - COALESCE(p.farde, 0))::numeric, 2)    AS delta_farde,
			ROUND(((COALESCE(c.farde, 0) - COALESCE(p.farde, 0)) * 100.0 /
			       NULLIF(COALESCE(p.farde, 0), 0))::numeric, 2)                AS growth_farde_pct,
			COALESCE(c.sold, 0)                                                  AS curr_sold,
			COALESCE(p.sold, 0)                                                  AS prev_sold,
			ROUND((COALESCE(c.sold, 0) - COALESCE(p.sold, 0))::numeric, 2)      AS delta_sold,
			ROUND(((COALESCE(c.sold, 0) - COALESCE(p.sold, 0)) * 100.0 /
			       NULLIF(COALESCE(p.sold, 0), 0))::numeric, 2)                 AS growth_sold_pct,
			COALESCE(c.visits, 0)                                                AS curr_visits,
			COALESCE(p.visits, 0)                                                AS prev_visits,
			CASE
				WHEN COALESCE(c.farde, 0) > COALESCE(p.farde, 0) THEN 'UP'
				WHEN COALESCE(c.farde, 0) < COALESCE(p.farde, 0) THEN 'DOWN'
				ELSE 'STABLE'
			END AS trend
		FROM curr c
		FULL OUTER JOIN prev p ON c.brand_name = p.brand_name
		ORDER BY curr_farde DESC
	`

	var rows []GrowthRow
	if err := db.Raw(sqlQuery, map[string]interface{}{
		"user_uuid":  userUUID,
		"curr_start": currStart,
		"curr_end":   currEnd,
		"prev_start": prevStart,
		"prev_end":   prevEnd,
	}).Scan(&rows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   rows,
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// 6. BRAND COMPETITION — parts de marché par marque pour l'agent
// ─────────────────────────────────────────────────────────────────────────────

func GetSEIBrandCompetition(c *fiber.Ctx) error {
	db := database.DB

	userUUID := c.Params("user_uuid")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if userUUID == "" || startDate == "" || endDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "user_uuid, start_date et end_date sont obligatoires",
		})
	}

	type BrandCompRow struct {
		BrandUUID   string  `json:"brand_uuid"`
		BrandName   string  `json:"brand_name"`
		TotalFarde  float64 `json:"total_farde"`
		TotalSold   float64 `json:"total_sold"`
		MarketShare float64 `json:"market_share"`
		BrandRank   int     `json:"brand_rank"`
		TotalVisits int64   `json:"total_visits"`
	}

	sqlQuery := `
		WITH base AS (
			SELECT
				b.uuid                                    AS brand_uuid,
				b.name                                    AS brand_name,
				COUNT(DISTINCT pf.uuid)                   AS total_visits,
				ROUND(SUM(pfi.number_farde)::numeric, 2)  AS total_farde,
				ROUND(SUM(pfi.sold)::numeric, 2)           AS total_sold
			FROM pos_form_items pfi
			INNER JOIN pos_forms pf ON pfi.pos_form_uuid = pf.uuid
			INNER JOIN brands b     ON pfi.brand_uuid = b.uuid
			WHERE pf.user_uuid = @user_uuid
			  AND pf.created_at BETWEEN @start_date AND @end_date
			  AND pf.deleted_at IS NULL AND pfi.deleted_at IS NULL
			GROUP BY b.uuid, b.name
		),
		total AS (
			SELECT COALESCE(SUM(total_farde), 0) AS g_farde FROM base
		)
		SELECT
			b.brand_uuid,
			b.brand_name,
			b.total_farde,
			b.total_sold,
			ROUND((b.total_farde * 100.0 / NULLIF(t.g_farde, 0))::numeric, 2) AS market_share,
			RANK() OVER (ORDER BY b.total_farde DESC)::int                      AS brand_rank,
			b.total_visits
		FROM base b, total t
		ORDER BY brand_rank
	`

	var rows []BrandCompRow
	if err := db.Raw(sqlQuery, map[string]interface{}{
		"user_uuid":  userUUID,
		"start_date": startDate,
		"end_date":   endDate,
	}).Scan(&rows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   rows,
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// 7. TOP POS — classement des meilleurs POS de l'agent
// ─────────────────────────────────────────────────────────────────────────────

func GetSEITopPos(c *fiber.Ctx) error {
	db := database.DB

	userUUID := c.Params("user_uuid")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	limit := c.Query("limit")

	if userUUID == "" || startDate == "" || endDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "user_uuid, start_date et end_date sont obligatoires",
		})
	}
	if limit == "" {
		limit = "10"
	}

	type TopPosRow struct {
		Rank        int     `json:"rank"`
		PosUUID     string  `json:"pos_uuid"`
		PosName     string  `json:"pos_name"`
		Shop        string  `json:"shop"`
		Postype     string  `json:"postype"`
		CommuneName string  `json:"commune_name"`
		TotalVisits int64   `json:"total_visits"`
		TotalFarde  float64 `json:"total_farde"`
		TotalSold   float64 `json:"total_sold"`
		AvgPrice    float64 `json:"avg_price"`
		FardeShare  float64 `json:"farde_share"`
	}

	sqlQuery := `
		WITH global_farde AS (
			SELECT COALESCE(SUM(pfi.number_farde), 0) AS g_farde
			FROM pos_form_items pfi
			INNER JOIN pos_forms pf ON pfi.pos_form_uuid = pf.uuid
			WHERE pf.user_uuid = @user_uuid
			  AND pf.created_at BETWEEN @start_date AND @end_date
			  AND pf.deleted_at IS NULL AND pfi.deleted_at IS NULL
		)
		SELECT
			RANK() OVER (ORDER BY SUM(pfi.number_farde) DESC)::int            AS rank,
			p.uuid                                                              AS pos_uuid,
			p.name                                                              AS pos_name,
			COALESCE(p.shop, '')                                                AS shop,
			COALESCE(NULLIF(p.postype, ''), 'Non défini')                      AS postype,
			COALESCE(co.name, '')                                               AS commune_name,
			COUNT(DISTINCT pf.uuid)                                             AS total_visits,
			ROUND(SUM(pfi.number_farde)::numeric, 2)                           AS total_farde,
			ROUND(SUM(pfi.sold)::numeric, 2)                                    AS total_sold,
			ROUND(AVG(pf.price)::numeric, 2)                                    AS avg_price,
			ROUND((SUM(pfi.number_farde) * 100.0 / NULLIF((SELECT g_farde FROM global_farde), 0))::numeric, 2) AS farde_share
		FROM pos_form_items pfi
		INNER JOIN pos_forms pf ON pfi.pos_form_uuid = pf.uuid
		INNER JOIN pos p        ON pf.pos_uuid = p.uuid
		LEFT  JOIN communes co  ON pf.commune_uuid = co.uuid
		WHERE pf.user_uuid = @user_uuid
		  AND pf.created_at BETWEEN @start_date AND @end_date
		  AND pf.deleted_at IS NULL AND pfi.deleted_at IS NULL
		GROUP BY p.uuid, p.name, p.shop, p.postype, co.name
		ORDER BY total_farde DESC
		LIMIT ` + limit + `
	`

	var rows []TopPosRow
	if err := db.Raw(sqlQuery, map[string]interface{}{
		"user_uuid":  userUUID,
		"start_date": startDate,
		"end_date":   endDate,
	}).Scan(&rows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   rows,
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// 8. HEATMAP DAY OF WEEK — activité par jour de la semaine
// ─────────────────────────────────────────────────────────────────────────────

func GetSEIHeatmap(c *fiber.Ctx) error {
	db := database.DB

	userUUID := c.Params("user_uuid")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if userUUID == "" || startDate == "" || endDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "user_uuid, start_date et end_date sont obligatoires",
		})
	}

	type HeatmapRow struct {
		DayOfWeek   int     `json:"day_of_week"`
		DayName     string  `json:"day_name"`
		BrandName   string  `json:"brand_name"`
		TotalFarde  float64 `json:"total_farde"`
		TotalSold   float64 `json:"total_sold"`
		TotalVisits int64   `json:"total_visits"`
		AvgFarde    float64 `json:"avg_farde"`
	}

	sqlQuery := `
		SELECT
			EXTRACT(ISODOW FROM pf.created_at)::int                  AS day_of_week,
			TO_CHAR(pf.created_at, 'Day')                            AS day_name,
			b.name                                                     AS brand_name,
			ROUND(SUM(pfi.number_farde)::numeric, 2)                  AS total_farde,
			ROUND(SUM(pfi.sold)::numeric, 2)                           AS total_sold,
			COUNT(DISTINCT pf.uuid)                                    AS total_visits,
			ROUND(AVG(pfi.number_farde)::numeric, 2)                   AS avg_farde
		FROM pos_form_items pfi
		INNER JOIN pos_forms pf ON pfi.pos_form_uuid = pf.uuid
		INNER JOIN brands b     ON pfi.brand_uuid = b.uuid
		WHERE pf.user_uuid = @user_uuid
		  AND pf.created_at BETWEEN @start_date AND @end_date
		  AND pf.deleted_at IS NULL AND pfi.deleted_at IS NULL
		GROUP BY day_of_week, day_name, b.name
		ORDER BY day_of_week, total_farde DESC
	`

	var rows []HeatmapRow
	if err := db.Raw(sqlQuery, map[string]interface{}{
		"user_uuid":  userUUID,
		"start_date": startDate,
		"end_date":   endDate,
	}).Scan(&rows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   rows,
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// 9. PRICE PIE CHART — distribution des prix déclarés par l'agent
// ─────────────────────────────────────────────────────────────────────────────

func GetSEIPricePieChart(c *fiber.Ctx) error {
	db := database.DB

	userUUID := c.Params("user_uuid")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if userUUID == "" || startDate == "" || endDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "user_uuid, start_date et end_date sont obligatoires",
		})
	}

	type PriceSliceRow struct {
		Price    float64 `json:"price"`
		Count    int64   `json:"count"`
		SharePct float64 `json:"share_pct"`
	}

	sqlQuery := `
		WITH base AS (
			SELECT
				pf.price,
				COUNT(*) AS cnt
			FROM pos_forms pf
			WHERE pf.user_uuid = @user_uuid
			  AND pf.created_at BETWEEN @start_date AND @end_date
			  AND pf.deleted_at IS NULL
			  AND pf.price > 0
			GROUP BY pf.price
		),
		total AS (SELECT COALESCE(SUM(cnt), 1) AS total_cnt FROM base)
		SELECT
			b.price,
			b.cnt                                               AS count,
			ROUND(b.cnt * 100.0 / t.total_cnt, 2)              AS share_pct
		FROM base b, total t
		ORDER BY b.cnt DESC
	`

	var rows []PriceSliceRow
	if err := db.Raw(sqlQuery, map[string]interface{}{
		"user_uuid":  userUUID,
		"start_date": startDate,
		"end_date":   endDate,
	}).Scan(&rows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   rows,
	})
}
