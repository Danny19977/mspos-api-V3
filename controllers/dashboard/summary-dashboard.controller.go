package dashboard

import (
	"time"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ======================== RÉSUMÉ EXÉCUTIF GLOBAL ========================

// ExecutiveSummary fournit un aperçu global de la performance pour les décideurs
func ExecutiveSummary(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var response models.ExecutiveSummaryResponse

	// 1. Métriques générales
	response.Overview = getOverviewMetrics(db, country_uuid, start_date, end_date)

	// 2. Performance opérationnelle
	response.Performance = getPerformanceMetrics(db, country_uuid, start_date, end_date)

	// 3. Distribution géographique
	response.GeographicDistribution = getGeographicMetrics(db, country_uuid, start_date, end_date)

	// 4. Performance des équipes
	response.TeamPerformance = getTeamPerformanceMetrics(db, country_uuid, start_date, end_date)

	// 5. Analyse des tendances
	response.TrendAnalysis = getTrendMetrics(db, country_uuid, start_date, end_date)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Résumé exécutif généré avec succès",
		"data":    response,
	})
}

// ======================== FONCTIONS DE CALCUL DES MÉTRIQUES ========================

func getOverviewMetrics(db any, country_uuid, start_date, end_date string) models.OverviewMetrics {
	var metrics models.OverviewMetrics

	// Total POS
	db.(*gorm.DB).Model(&models.Pos{}).Where("country_uuid = ? AND deleted_at IS NULL", country_uuid).Count(&metrics.TotalPOS)

	// POS actifs (ayant eu au moins une visite)
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at BETWEEN ? AND ? AND deleted_at IS NULL", country_uuid, start_date, end_date).
		Distinct("pos_uuid").Count(&metrics.ActivePOS)

	// Total visites
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at BETWEEN ? AND ? AND deleted_at IS NULL", country_uuid, start_date, end_date).
		Count(&metrics.TotalVisits)

	// Total utilisateurs
	db.(*gorm.DB).Model(&models.User{}).Where("country_uuid = ? AND deleted_at IS NULL", country_uuid).Count(&metrics.TotalUsers)

	// Total provinces
	db.(*gorm.DB).Model(&models.Province{}).Where("country_uuid = ? AND deleted_at IS NULL", country_uuid).Count(&metrics.TotalProvinces)

	// Total aires
	db.(*gorm.DB).Model(&models.Area{}).Where("country_uuid = ? AND deleted_at IS NULL", country_uuid).Count(&metrics.TotalAreas)

	// Calcul des métriques dérivées
	if metrics.TotalPOS > 0 {
		metrics.MarketPenetration = float64(metrics.ActivePOS) / float64(metrics.TotalPOS) * 100
	}

	// Moyenne visites par jour
	if start_date != "" && end_date != "" {
		startTime, _ := time.Parse("2006-01-02", start_date)
		endTime, _ := time.Parse("2006-01-02", end_date)
		days := endTime.Sub(startTime).Hours() / 24
		if days > 0 {
			metrics.AverageVisitsPerDay = float64(metrics.TotalVisits) / days
		}
	}

	return metrics
}

func getPerformanceMetrics(db any, country_uuid, start_date, end_date string) models.PerformanceMetrics {
	var metrics models.PerformanceMetrics

	// Taux d'objectif de visite
	var totalObjective, totalVisits int64
	db.(*gorm.DB).Raw(`
		SELECT 
			COALESCE(SUM(
				CASE
					WHEN users.title = 'ASM' THEN 10 * ((?::date - ?::date) + 1)
					WHEN users.title = 'Supervisor' THEN 20 * ((?::date - ?::date) + 1)
					WHEN users.title = 'DR' THEN 40 * ((?::date - ?::date) + 1)
					WHEN users.title = 'Cyclo' THEN 40 * ((?::date - ?::date) + 1)
					ELSE 1
				END
			), 0) as total_objective,
			COUNT(pos_forms.uuid) as total_visits
		FROM pos_forms
		JOIN users ON users.uuid = pos_forms.user_uuid
		WHERE pos_forms.country_uuid = ? 
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
	`, end_date, start_date, end_date, start_date, end_date, start_date, end_date, start_date, country_uuid, start_date, end_date).
		Row().Scan(&totalObjective, &totalVisits)

	if totalObjective > 0 {
		metrics.VisitObjectiveRate = float64(totalVisits) / float64(totalObjective) * 100
	}

	// Taux de completion (formulaires complets)
	var completeForms int64
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at BETWEEN ? AND ? AND pos_uuid IS NOT NULL AND pos_uuid != '' AND deleted_at IS NULL",
			country_uuid, start_date, end_date).Count(&completeForms)

	if totalVisits > 0 {
		metrics.CompletionRate = float64(completeForms) / float64(totalVisits) * 100
	}

	// Score d'efficacité (combinaison de plusieurs facteurs)
	metrics.EfficiencyScore = (metrics.VisitObjectiveRate + metrics.CompletionRate) / 2

	// Marque la plus performante
	var topBrand string
	db.(*gorm.DB).Raw(`
		SELECT brands.name
		FROM pos_form_items
		JOIN brands ON pos_form_items.brand_uuid = brands.uuid
		JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		WHERE pos_forms.country_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY brands.name
		ORDER BY SUM(pos_form_items.number_farde) DESC
		LIMIT 1
	`, country_uuid, start_date, end_date).Row().Scan(&topBrand)
	metrics.TopBrandPerformance = topBrand

	// Score moyen des formulaires (basé sur le prix)
	var avgPrice float64
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at BETWEEN ? AND ? AND deleted_at IS NULL", country_uuid, start_date, end_date).
		Select("COALESCE(AVG(price), 0)").Scan(&avgPrice)
	metrics.AverageFormScore = avgPrice

	return metrics
}

func getGeographicMetrics(db any, country_uuid, start_date, end_date string) models.GeographicMetrics {
	var metrics models.GeographicMetrics

	// Province la plus performante
	var topProvince string
	db.(*gorm.DB).Raw(`
		SELECT provinces.name
		FROM pos_forms
		JOIN provinces ON pos_forms.province_uuid = provinces.uuid
		WHERE pos_forms.country_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY provinces.name
		ORDER BY COUNT(pos_forms.uuid) DESC
		LIMIT 1
	`, country_uuid, start_date, end_date).Row().Scan(&topProvince)
	metrics.TopPerformingProvince = topProvince

	// Aire la plus performante
	var topArea string
	db.(*gorm.DB).Raw(`
		SELECT areas.name
		FROM pos_forms
		JOIN areas ON pos_forms.area_uuid = areas.uuid
		WHERE pos_forms.country_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY areas.name
		ORDER BY COUNT(pos_forms.uuid) DESC
		LIMIT 1
	`, country_uuid, start_date, end_date).Row().Scan(&topArea)
	metrics.TopPerformingArea = topArea

	// Sous-aire la plus performante
	var topSubArea string
	db.(*gorm.DB).Raw(`
		SELECT sub_areas.name
		FROM pos_forms
		JOIN sub_areas ON pos_forms.sub_area_uuid = sub_areas.uuid
		WHERE pos_forms.country_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY sub_areas.name
		ORDER BY COUNT(pos_forms.uuid) DESC
		LIMIT 1
	`, country_uuid, start_date, end_date).Row().Scan(&topSubArea)
	metrics.TopPerformingSubArea = topSubArea

	// Commune la plus performante
	var topCommune string
	db.(*gorm.DB).Raw(`
		SELECT communes.name
		FROM pos_forms
		JOIN communes ON pos_forms.commune_uuid = communes.uuid
		WHERE pos_forms.country_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY communes.name
		ORDER BY COUNT(pos_forms.uuid) DESC
		LIMIT 1
	`, country_uuid, start_date, end_date).Row().Scan(&topCommune)
	metrics.TopPerformingCommune = topCommune

	// Calcul du pourcentage de couverture
	var totalAreas, activeAreas int64
	db.(*gorm.DB).Model(&models.Area{}).Where("country_uuid = ? AND deleted_at IS NULL", country_uuid).Count(&totalAreas)
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at BETWEEN ? AND ? AND deleted_at IS NULL", country_uuid, start_date, end_date).
		Distinct("area_uuid").Count(&activeAreas)

	if totalAreas > 0 {
		metrics.CoveragePercentage = float64(activeAreas) / float64(totalAreas) * 100
	}

	return metrics
}

func getTeamPerformanceMetrics(db any, country_uuid, start_date, end_date string) models.TeamPerformanceMetrics {
	var metrics models.TeamPerformanceMetrics

	// Total membres d'équipe
	db.(*gorm.DB).Model(&models.User{}).Where("country_uuid = ? AND deleted_at IS NULL", country_uuid).Count(&metrics.TotalTeamMembers)

	// Membres actifs (ayant fait au moins une visite)
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at BETWEEN ? AND ? AND deleted_at IS NULL", country_uuid, start_date, end_date).
		Distinct("user_uuid").Count(&metrics.ActiveTeamMembers)

	// Top performer
	var topPerformer string
	db.(*gorm.DB).Raw(`
		SELECT users.fullname
		FROM pos_forms
		JOIN users ON pos_forms.user_uuid = users.uuid
		WHERE pos_forms.country_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY users.fullname
		ORDER BY COUNT(pos_forms.uuid) DESC
		LIMIT 1
	`, country_uuid, start_date, end_date).Row().Scan(&topPerformer)
	metrics.TopPerformer = topPerformer

	// Efficacité moyenne de l'équipe
	if metrics.TotalTeamMembers > 0 {
		metrics.AverageTeamEfficiency = float64(metrics.ActiveTeamMembers) / float64(metrics.TotalTeamMembers) * 100
	}

	return metrics
}

func getTrendMetrics(db any, country_uuid, start_date, end_date string) models.TrendMetrics {
	var metrics models.TrendMetrics

	// Calcul simple des tendances (dernière semaine vs semaine précédente)
	var currentWeekVisits, previousWeekVisits int64

	// Visites de la semaine actuelle
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at >= ? AND deleted_at IS NULL", country_uuid,
			time.Now().AddDate(0, 0, -7).Format("2006-01-02")).Count(&currentWeekVisits)

	// Visites de la semaine précédente
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at BETWEEN ? AND ? AND deleted_at IS NULL", country_uuid,
			time.Now().AddDate(0, 0, -14).Format("2006-01-02"),
			time.Now().AddDate(0, 0, -7).Format("2006-01-02")).Count(&previousWeekVisits)

	// Détermination de la tendance
	if currentWeekVisits > previousWeekVisits {
		metrics.VisitTrend = "croissante"
		if previousWeekVisits > 0 {
			metrics.MonthlyGrowth = float64(currentWeekVisits-previousWeekVisits) / float64(previousWeekVisits) * 100
		}
	} else if currentWeekVisits < previousWeekVisits {
		metrics.VisitTrend = "décroissante"
		if previousWeekVisits > 0 {
			metrics.MonthlyGrowth = -float64(previousWeekVisits-currentWeekVisits) / float64(previousWeekVisits) * 100
		}
	} else {
		metrics.VisitTrend = "stable"
		metrics.MonthlyGrowth = 0
	}

	// Prédiction simple pour le mois prochain (basée sur la tendance actuelle)
	if metrics.MonthlyGrowth > 0 {
		metrics.PredictedNextMonth = float64(currentWeekVisits) * 4 * (1 + metrics.MonthlyGrowth/100)
	} else {
		metrics.PredictedNextMonth = float64(currentWeekVisits) * 4
	}

	return metrics
}

// ======================== RÉSUMÉS SPÉCIALISÉS ========================

// RegionalSummary fournit un résumé focalisé sur une région spécifique
func RegionalSummary(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	province_uuid := c.Query("province_uuid")
	area_uuid := c.Query("area_uuid")
	sub_area_uuid := c.Query("sub_area_uuid")
	commune_uuid := c.Query("commune_uuid")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var response models.RegionalSummaryResponse

	// Informations sur la région
	response.RegionInfo = getRegionInfo(db, country_uuid, province_uuid, area_uuid, sub_area_uuid, commune_uuid)

	// Performance régionale
	response.Performance = getRegionalPerformance(db, country_uuid, province_uuid, area_uuid, sub_area_uuid, commune_uuid, start_date, end_date)

	// Comparaison avec d'autres régions
	response.Comparison = getRegionalComparison(db, country_uuid, province_uuid, area_uuid, sub_area_uuid, commune_uuid, start_date, end_date)

	// Top performers
	response.TopPerformers = getTopPerformers(db, country_uuid, province_uuid, area_uuid, sub_area_uuid, commune_uuid, start_date, end_date)

	// Opportunités
	response.Opportunities = identifyOpportunities(response.Performance)

	// Recommandations
	response.Recommendations = generateRecommendations(response)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Résumé régional généré avec succès",
		"data":    response,
	})
}

func getRegionInfo(db any, country_uuid, province_uuid, area_uuid, sub_area_uuid, commune_uuid string) models.RegionInfo {
	var info models.RegionInfo

	if commune_uuid != "" {
		// Informations pour une commune spécifique
		db.(*gorm.DB).Model(&models.Commune{}).Where("uuid = ?", commune_uuid).Select("name").Scan(&info.Name)
		info.Type = "Commune"

		db.(*gorm.DB).Model(&models.Pos{}).Where("commune_uuid = ? AND deleted_at IS NULL", commune_uuid).Count(&info.TotalPOS)
		db.(*gorm.DB).Model(&models.User{}).Where("commune_uuid = ? AND deleted_at IS NULL", commune_uuid).Count(&info.TotalUsers)
		// Pour une commune, les sous-aires et communes sont 0 car c'est le niveau le plus bas
		info.TotalSubAreas = 0
		info.TotalCommunes = 1
	} else if sub_area_uuid != "" {
		// Informations pour une sous-aire spécifique
		db.(*gorm.DB).Model(&models.SubArea{}).Where("uuid = ?", sub_area_uuid).Select("name").Scan(&info.Name)
		info.Type = "SubArea"

		db.(*gorm.DB).Model(&models.Pos{}).Where("sub_area_uuid = ? AND deleted_at IS NULL", sub_area_uuid).Count(&info.TotalPOS)
		db.(*gorm.DB).Model(&models.User{}).Where("sub_area_uuid = ? AND deleted_at IS NULL", sub_area_uuid).Count(&info.TotalUsers)
		info.TotalSubAreas = 1
		db.(*gorm.DB).Model(&models.Commune{}).Where("sub_area_uuid = ? AND deleted_at IS NULL", sub_area_uuid).Count(&info.TotalCommunes)
	} else if area_uuid != "" {
		// Informations pour une aire spécifique
		db.(*gorm.DB).Model(&models.Area{}).Where("uuid = ?", area_uuid).Select("name").Scan(&info.Name)
		info.Type = "Area"

		db.(*gorm.DB).Model(&models.Pos{}).Where("area_uuid = ? AND deleted_at IS NULL", area_uuid).Count(&info.TotalPOS)
		db.(*gorm.DB).Model(&models.User{}).Where("area_uuid = ? AND deleted_at IS NULL", area_uuid).Count(&info.TotalUsers)
		db.(*gorm.DB).Model(&models.SubArea{}).Where("area_uuid = ? AND deleted_at IS NULL", area_uuid).Count(&info.TotalSubAreas)
		db.(*gorm.DB).Model(&models.Commune{}).Where("area_uuid = ? AND deleted_at IS NULL", area_uuid).Count(&info.TotalCommunes)
	} else if province_uuid != "" {
		// Informations pour une province spécifique
		db.(*gorm.DB).Model(&models.Province{}).Where("uuid = ?", province_uuid).Select("name").Scan(&info.Name)
		info.Type = "Province"

		db.(*gorm.DB).Model(&models.Pos{}).Where("province_uuid = ? AND deleted_at IS NULL", province_uuid).Count(&info.TotalPOS)
		db.(*gorm.DB).Model(&models.User{}).Where("province_uuid = ? AND deleted_at IS NULL", province_uuid).Count(&info.TotalUsers)
		db.(*gorm.DB).Model(&models.Area{}).Where("province_uuid = ? AND deleted_at IS NULL", province_uuid).Count(&info.TotalSubAreas)
		db.(*gorm.DB).Model(&models.Commune{}).Where("province_uuid = ? AND deleted_at IS NULL", province_uuid).Count(&info.TotalCommunes)
	}

	return info
}

func getRegionalPerformance(db any, country_uuid, province_uuid, area_uuid, sub_area_uuid, commune_uuid, start_date, end_date string) models.RegionalPerformance {
	var performance models.RegionalPerformance

	whereClause := "country_uuid = ? AND deleted_at IS NULL"
	args := []interface{}{country_uuid}

	if commune_uuid != "" {
		whereClause += " AND commune_uuid = ?"
		args = append(args, commune_uuid)
	} else if sub_area_uuid != "" {
		whereClause += " AND sub_area_uuid = ?"
		args = append(args, sub_area_uuid)
	} else if area_uuid != "" {
		whereClause += " AND area_uuid = ?"
		args = append(args, area_uuid)
	} else if province_uuid != "" {
		whereClause += " AND province_uuid = ?"
		args = append(args, province_uuid)
	}
	if start_date != "" && end_date != "" {
		whereClause += " AND created_at BETWEEN ? AND ?"
		args = append(args, start_date, end_date)
	}

	// Visites de la période
	db.(*gorm.DB).Model(&models.PosForm{}).Where(whereClause, args...).Count(&performance.VisitsThisPeriod)

	// Détermination du rating d'efficacité
	if performance.ObjectiveRate >= 90 {
		performance.EfficiencyRating = "Excellent"
	} else if performance.ObjectiveRate >= 75 {
		performance.EfficiencyRating = "Bon"
	} else if performance.ObjectiveRate >= 60 {
		performance.EfficiencyRating = "Moyen"
	} else {
		performance.EfficiencyRating = "À améliorer"
	}

	return performance
}

func getRegionalComparison(db any, country_uuid, province_uuid, area_uuid, sub_area_uuid, commune_uuid, start_date, end_date string) models.RegionalComparison {
	var comparison models.RegionalComparison

	// Paramètres disponibles pour calculs futurs
	_ = db
	_ = country_uuid
	_ = province_uuid
	_ = area_uuid
	_ = sub_area_uuid
	_ = commune_uuid
	_ = start_date
	_ = end_date

	// Logique de comparaison simplifiée
	comparison.RankAmongRegions = 1        // À calculer selon les besoins
	comparison.PerformanceVsAverage = 15.5 // Exemple
	comparison.BestMetric = "Taux de visite"
	comparison.WeakestMetric = "Completion rate"

	return comparison
}

func getTopPerformers(db any, country_uuid, province_uuid, area_uuid, sub_area_uuid, commune_uuid, start_date, end_date string) []models.TopPerformer {
	var performers []models.TopPerformer

	whereClause := "pos_forms.country_uuid = ? AND pos_forms.deleted_at IS NULL"
	args := []interface{}{country_uuid}

	if commune_uuid != "" {
		whereClause += " AND pos_forms.commune_uuid = ?"
		args = append(args, commune_uuid)
	} else if sub_area_uuid != "" {
		whereClause += " AND pos_forms.sub_area_uuid = ?"
		args = append(args, sub_area_uuid)
	} else if area_uuid != "" {
		whereClause += " AND pos_forms.area_uuid = ?"
		args = append(args, area_uuid)
	} else if province_uuid != "" {
		whereClause += " AND pos_forms.province_uuid = ?"
		args = append(args, province_uuid)
	}
	if start_date != "" && end_date != "" {
		whereClause += " AND pos_forms.created_at BETWEEN ? AND ?"
		args = append(args, start_date, end_date)
	}

	rows, err := db.(*gorm.DB).Raw(`
		SELECT 
			users.fullname,
			users.title,
			COUNT(pos_forms.uuid) as visits,
			ROUND((COUNT(pos_forms.uuid) / (
				CASE
					WHEN users.title = 'ASM' THEN 10
					WHEN users.title = 'Supervisor' THEN 20
					WHEN users.title = 'DR' THEN 40
					WHEN users.title = 'Cyclo' THEN 40
					ELSE 1
				END
			) * 100.0), 2) as objective_rate
		FROM pos_forms
		JOIN users ON pos_forms.user_uuid = users.uuid
		WHERE `+whereClause+`
		GROUP BY users.fullname, users.title
		ORDER BY visits DESC
		LIMIT 5
	`, args...).Rows()

	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var performer models.TopPerformer
			rows.Scan(&performer.Name, &performer.Role, &performer.Visits, &performer.ObjectiveRate)

			// Attribution des récompenses
			if performer.ObjectiveRate >= 100 {
				performer.SpecialAchievement = "Objectif dépassé"
			} else if performer.ObjectiveRate >= 90 {
				performer.SpecialAchievement = "Performance excellente"
			} else {
				performer.SpecialAchievement = "Bon contributeur"
			}

			performers = append(performers, performer)
		}
	}

	return performers
}

func identifyOpportunities(performance models.RegionalPerformance) []models.Opportunity {
	var opportunities []models.Opportunity

	if performance.ObjectiveRate < 80 {
		opportunities = append(opportunities, models.Opportunity{
			Area:            "Performance des visites",
			Potential:       "Amélioration du taux d'objectif",
			EstimatedImpact: 25.0,
			Effort:          "Moyen",
			Timeline:        "3 mois",
		})
	}

	// Opportunité basée sur le nombre de visites
	if performance.VisitsThisPeriod < 100 {
		opportunities = append(opportunities, models.Opportunity{
			Area:            "Fréquence des visites",
			Potential:       "Augmentation du nombre de visites",
			EstimatedImpact: 20.0,
			Effort:          "Moyen",
			Timeline:        "2 mois",
		})
	}

	return opportunities
}

func generateRecommendations(summary interface{}) []models.Recommendation {
	var recommendations []models.Recommendation

	// Cette fonction sera implémentée avec des recommandations basées sur les métriques reçues
	// Pour l'instant, on retourne un exemple

	recommendations = append(recommendations, models.Recommendation{
		Priority:        "Haute",
		Action:          "Améliorer la formation des équipes",
		ExpectedROI:     "150%",
		Timeline:        "3 mois",
		ResponsibleTeam: "Management régional",
	})

	return recommendations
}

// ======================== TABLEAUX DE BORD RAPIDES ========================

// QuickDashboard fournit un aperçu rapide pour les décisions urgentes
func QuickDashboard(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")

	var response models.QuickDashboardResponse
	response.LastUpdated = time.Now().Format("2006-01-02 15:04:05")

	// Métriques rapides
	response.KeyMetrics = getQuickMetrics(db, country_uuid)

	// Statistiques du jour
	response.TodayStats = getTodayStatistics(db, country_uuid)

	// Actions urgentes
	response.UrgentActions = getUrgentActions(response.KeyMetrics, response.TodayStats)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Dashboard rapide généré avec succès",
		"data":    response,
	})
}

func getQuickMetrics(db any, country_uuid string) models.QuickMetrics {
	var metrics models.QuickMetrics

	today := time.Now().Format("2006-01-02")
	weekStart := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	monthStart := time.Now().AddDate(0, -1, 0).Format("2006-01-02")

	// Visites aujourd'hui
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND DATE(created_at) = ? AND deleted_at IS NULL", country_uuid, today).
		Count(&metrics.VisitsToday)

	// Visites cette semaine
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at >= ? AND deleted_at IS NULL", country_uuid, weekStart).
		Count(&metrics.VisitsThisWeek)

	// Visites ce mois
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at >= ? AND deleted_at IS NULL", country_uuid, monthStart).
		Count(&metrics.VisitsThisMonth)

	// Utilisateurs actifs aujourd'hui
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND DATE(created_at) = ? AND deleted_at IS NULL", country_uuid, today).
		Distinct("user_uuid").Count(&metrics.ActiveUsersToday)

	// Taux de completion aujourd'hui
	var completeForms int64
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND DATE(created_at) = ? AND pos_uuid IS NOT NULL AND pos_uuid != '' AND deleted_at IS NULL",
			country_uuid, today).Count(&completeForms)

	if metrics.VisitsToday > 0 {
		metrics.CompletionRateToday = float64(completeForms) / float64(metrics.VisitsToday) * 100
	}

	// Top brand aujourd'hui
	var topBrand string
	db.(*gorm.DB).Raw(`
		SELECT brands.name
		FROM pos_form_items
		JOIN brands ON pos_form_items.brand_uuid = brands.uuid
		JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		WHERE pos_forms.country_uuid = ?
		AND DATE(pos_forms.created_at) = ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY brands.name
		ORDER BY SUM(pos_form_items.number_farde) DESC
		LIMIT 1
	`, country_uuid, today).Row().Scan(&topBrand)
	metrics.TopBrandToday = topBrand

	return metrics
}

func getTodayStatistics(db any, country_uuid string) models.TodayStatistics {
	var stats models.TodayStatistics

	today := time.Now().Format("2006-01-02")

	// Visites par heure (24 heures)
	stats.HourlyVisits = make([]int, 24)

	rows, err := db.(*gorm.DB).Raw(`
		SELECT 
			EXTRACT(HOUR FROM created_at) as hour,
			COUNT(*) as visit_count
		FROM pos_forms
		WHERE country_uuid = ?
		AND DATE(created_at) = ?
		AND deleted_at IS NULL
		GROUP BY EXTRACT(HOUR FROM created_at)
		ORDER BY hour
	`, country_uuid, today).Rows()

	if err == nil {
		defer rows.Close()
		maxVisits := 0
		for rows.Next() {
			var hour, count int
			rows.Scan(&hour, &count)
			if hour >= 0 && hour < 24 {
				stats.HourlyVisits[hour] = count
				if count > maxVisits {
					maxVisits = count
					stats.PeakHour = hour
				}
			}
		}
	}

	// Provinces actives aujourd'hui
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND DATE(created_at) = ? AND deleted_at IS NULL", country_uuid, today).
		Distinct("province_uuid").Count(&stats.ActiveProvinces)

	// Nouveaux POS visités aujourd'hui
	db.(*gorm.DB).Raw(`
		SELECT COUNT(DISTINCT pos_forms.pos_uuid)
		FROM pos_forms
		WHERE pos_forms.country_uuid = ?
		AND DATE(pos_forms.created_at) = ?
		AND pos_forms.deleted_at IS NULL
		AND pos_forms.pos_uuid NOT IN (
			SELECT DISTINCT pos_uuid 
			FROM pos_forms 
			WHERE country_uuid = ? 
			AND DATE(created_at) < ?
			AND deleted_at IS NULL
			AND pos_uuid IS NOT NULL
		)
	`, country_uuid, today, country_uuid, today).Row().Scan(&stats.NewPOSVisited)

	return stats
}

func getUrgentActions(metrics models.QuickMetrics, stats models.TodayStatistics) []models.UrgentAction {
	var actions []models.UrgentAction

	// Action si les visites du jour sont trop faibles
	if metrics.VisitsToday < 50 { // Seuil exemple
		actions = append(actions, models.UrgentAction{
			Priority:    "Critique",
			Description: "Visites journalières en dessous du seuil critique",
			Deadline:    "Fin de journée",
			Owner:       "Managers régionaux",
			Impact:      "Objectifs mensuels compromis",
		})
	}

	// Action si le taux de completion est faible
	if metrics.CompletionRateToday < 60 {
		actions = append(actions, models.UrgentAction{
			Priority:    "Haute",
			Description: "Taux de completion des formulaires insuffisant",
			Deadline:    "24 heures",
			Owner:       "Équipes terrain",
			Impact:      "Qualité des données compromise",
		})
	}

	// Action si peu d'utilisateurs actifs
	if metrics.ActiveUsersToday < 10 {
		actions = append(actions, models.UrgentAction{
			Priority:    "Moyenne",
			Description: "Mobilisation insuffisante des équipes",
			Deadline:    "48 heures",
			Owner:       "RH et Management",
			Impact:      "Couverture territoriale réduite",
		})
	}

	return actions
}

// ======================== ANALYSE COMPARATIVE ========================

// CompetitiveAnalysis fournit une analyse comparative avec les périodes précédentes
func CompetitiveAnalysis(c *fiber.Ctx) error {
	db := database.DB

	country_uuid := c.Query("country_uuid")
	current_start := c.Query("current_start")
	current_end := c.Query("current_end")
	previous_start := c.Query("previous_start")
	previous_end := c.Query("previous_end")

	var response models.CompetitiveAnalysisResponse

	// Analyse de la période actuelle
	response.CurrentPeriod = getPeriodAnalysis(db, country_uuid, current_start, current_end, "Période actuelle")

	// Analyse de la période précédente
	response.PreviousPeriod = getPeriodAnalysis(db, country_uuid, previous_start, previous_end, "Période précédente")

	// Comparaison
	response.Comparison = comparePerformance(response.CurrentPeriod, response.PreviousPeriod)

	// Tendances
	response.Trends = analyzeTrends(response.CurrentPeriod, response.PreviousPeriod)

	// Insights compétitifs
	response.Insights = generateCompetitiveInsights(response.Comparison)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Analyse comparative générée avec succès",
		"data":    response,
	})
}

func getPeriodAnalysis(db any, country_uuid, start_date, end_date, name string) models.PeriodAnalysis {
	var analysis models.PeriodAnalysis
	analysis.Name = name

	// Total visites
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at BETWEEN ? AND ? AND deleted_at IS NULL", country_uuid, start_date, end_date).
		Count(&analysis.TotalVisits)

	// Taux de completion
	var completeForms int64
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at BETWEEN ? AND ? AND pos_uuid IS NOT NULL AND pos_uuid != '' AND deleted_at IS NULL",
			country_uuid, start_date, end_date).Count(&completeForms)

	if analysis.TotalVisits > 0 {
		analysis.CompletionRate = float64(completeForms) / float64(analysis.TotalVisits) * 100
	}

	// Utilisateurs actifs uniques
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at BETWEEN ? AND ? AND deleted_at IS NULL", country_uuid, start_date, end_date).
		Distinct("user_uuid").Count(&analysis.UniqueUsersActive)

	// POS uniques visités
	db.(*gorm.DB).Model(&models.PosForm{}).
		Where("country_uuid = ? AND created_at BETWEEN ? AND ? AND pos_uuid IS NOT NULL AND deleted_at IS NULL",
			country_uuid, start_date, end_date).
		Distinct("pos_uuid").Count(&analysis.UniquePOSVisited)

	// Top brand
	var topBrand string
	db.(*gorm.DB).Raw(`
		SELECT brands.name
		FROM pos_form_items
		JOIN brands ON pos_form_items.brand_uuid = brands.uuid
		JOIN pos_forms ON pos_form_items.pos_form_uuid = pos_forms.uuid
		WHERE pos_forms.country_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY brands.name
		ORDER BY SUM(pos_form_items.number_farde) DESC
		LIMIT 1
	`, country_uuid, start_date, end_date).Row().Scan(&topBrand)
	analysis.TopBrand = topBrand

	// Top province
	var topProvince string
	db.(*gorm.DB).Raw(`
		SELECT provinces.name
		FROM pos_forms
		JOIN provinces ON pos_forms.province_uuid = provinces.uuid
		WHERE pos_forms.country_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY provinces.name
		ORDER BY COUNT(pos_forms.uuid) DESC
		LIMIT 1
	`, country_uuid, start_date, end_date).Row().Scan(&topProvince)
	analysis.TopProvince = topProvince

	// Top aire
	var topArea string
	db.(*gorm.DB).Raw(`
		SELECT areas.name
		FROM pos_forms
		JOIN areas ON pos_forms.area_uuid = areas.uuid
		WHERE pos_forms.country_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY areas.name
		ORDER BY COUNT(pos_forms.uuid) DESC
		LIMIT 1
	`, country_uuid, start_date, end_date).Row().Scan(&topArea)
	analysis.TopArea = topArea

	// Top sous-aire
	var topSubArea string
	db.(*gorm.DB).Raw(`
		SELECT sub_areas.name
		FROM pos_forms
		JOIN sub_areas ON pos_forms.sub_area_uuid = sub_areas.uuid
		WHERE pos_forms.country_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY sub_areas.name
		ORDER BY COUNT(pos_forms.uuid) DESC
		LIMIT 1
	`, country_uuid, start_date, end_date).Row().Scan(&topSubArea)
	analysis.TopSubArea = topSubArea

	// Top commune
	var topCommune string
	db.(*gorm.DB).Raw(`
		SELECT communes.name
		FROM pos_forms
		JOIN communes ON pos_forms.commune_uuid = communes.uuid
		WHERE pos_forms.country_uuid = ?
		AND pos_forms.created_at BETWEEN ? AND ?
		AND pos_forms.deleted_at IS NULL
		GROUP BY communes.name
		ORDER BY COUNT(pos_forms.uuid) DESC
		LIMIT 1
	`, country_uuid, start_date, end_date).Row().Scan(&topCommune)
	analysis.TopCommune = topCommune

	// Score d'efficacité
	analysis.EfficiencyScore = (analysis.CompletionRate + float64(analysis.UniqueUsersActive)*10) / 2

	return analysis
}

func comparePerformance(current, previous models.PeriodAnalysis) models.ComparisonMetrics {
	var comparison models.ComparisonMetrics

	// Croissance des visites
	if previous.TotalVisits > 0 {
		comparison.VisitGrowth = (float64(current.TotalVisits-previous.TotalVisits) / float64(previous.TotalVisits)) * 100
	}

	// Changement d'efficacité
	if previous.EfficiencyScore > 0 {
		comparison.EfficiencyChange = ((current.EfficiencyScore - previous.EfficiencyScore) / previous.EfficiencyScore) * 100
	}

	// Changement d'engagement utilisateur
	if previous.UniqueUsersActive > 0 {
		comparison.UserEngagementChange = (float64(current.UniqueUsersActive-previous.UniqueUsersActive) / float64(previous.UniqueUsersActive)) * 100
	}

	// Expansion de marché
	if previous.UniquePOSVisited > 0 {
		comparison.MarketExpansion = (float64(current.UniquePOSVisited-previous.UniquePOSVisited) / float64(previous.UniquePOSVisited)) * 100
	}

	// Performance globale
	avgGrowth := (comparison.VisitGrowth + comparison.EfficiencyChange) / 2
	if avgGrowth > 10 {
		comparison.OverallPerformance = "Excellente"
	} else if avgGrowth > 5 {
		comparison.OverallPerformance = "Bonne"
	} else if avgGrowth > 0 {
		comparison.OverallPerformance = "Modérée"
	} else {
		comparison.OverallPerformance = "À améliorer"
	}

	return comparison
}

func analyzeTrends(current, previous models.PeriodAnalysis) []models.TrendAnalysis {
	var trends []models.TrendAnalysis

	// Tendance des visites
	visitGrowth := float64(current.TotalVisits-previous.TotalVisits) / float64(previous.TotalVisits) * 100
	trends = append(trends, models.TrendAnalysis{
		Metric:       "Visites",
		Direction:    getTrendDirection(visitGrowth),
		Magnitude:    visitGrowth,
		Significance: getTrendSignificance(visitGrowth),
		Prediction:   getTrendPrediction("visites", visitGrowth),
	})

	// Tendance de l'efficacité
	efficiencyGrowth := (current.EfficiencyScore - previous.EfficiencyScore) / previous.EfficiencyScore * 100
	trends = append(trends, models.TrendAnalysis{
		Metric:       "Efficacité",
		Direction:    getTrendDirection(efficiencyGrowth),
		Magnitude:    efficiencyGrowth,
		Significance: getTrendSignificance(efficiencyGrowth),
		Prediction:   getTrendPrediction("efficacité", efficiencyGrowth),
	})

	return trends
}

func getTrendDirection(growth float64) string {
	if growth > 0 {
		return "croissante"
	} else if growth < 0 {
		return "décroissante"
	}
	return "stable"
}

func getTrendSignificance(growth float64) string {
	absGrowth := growth
	if absGrowth < 0 {
		absGrowth = -absGrowth
	}

	if absGrowth > 20 {
		return "très significative"
	} else if absGrowth > 10 {
		return "significative"
	} else if absGrowth > 5 {
		return "modérée"
	}
	return "faible"
}

func getTrendPrediction(metric string, growth float64) string {
	if growth > 15 {
		return "Forte progression attendue pour " + metric
	} else if growth > 5 {
		return "Progression stable attendue pour " + metric
	} else if growth < -10 {
		return "Risque de déclin pour " + metric
	}
	return "Tendance stable attendue pour " + metric
}

func generateCompetitiveInsights(comparison models.ComparisonMetrics) []models.CompetitiveInsight {
	var insights []models.CompetitiveInsight

	// Insight sur la croissance des visites
	if comparison.VisitGrowth > 20 {
		insights = append(insights, models.CompetitiveInsight{
			Category:    "Performance",
			Finding:     "Croissance exceptionnelle des visites (+20%)",
			Implication: "Momentum très positif, équipes motivées",
			ActionPlan:  "Maintenir la dynamique, analyser les facteurs de succès",
			Priority:    "Moyenne",
		})
	} else if comparison.VisitGrowth < -10 {
		insights = append(insights, models.CompetitiveInsight{
			Category:    "Performance",
			Finding:     "Déclin significatif des visites (-10%)",
			Implication: "Problèmes opérationnels ou de motivation",
			ActionPlan:  "Plan d'action correctif immédiat",
			Priority:    "Critique",
		})
	}

	// Insight sur l'efficacité
	if comparison.EfficiencyChange > 15 {
		insights = append(insights, models.CompetitiveInsight{
			Category:    "Efficacité",
			Finding:     "Amélioration significative de l'efficacité",
			Implication: "Optimisation des processus réussie",
			ActionPlan:  "Documenter et répliquer les bonnes pratiques",
			Priority:    "Haute",
		})
	}

	// Insight sur l'expansion
	if comparison.MarketExpansion > 25 {
		insights = append(insights, models.CompetitiveInsight{
			Category:    "Expansion",
			Finding:     "Forte expansion du marché couvert",
			Implication: "Stratégie de pénétration efficace",
			ActionPlan:  "Consolider les gains et planifier la phase suivante",
			Priority:    "Haute",
		})
	}

	return insights
}
