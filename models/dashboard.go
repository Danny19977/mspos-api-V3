package models

// ======================== STRUCTURES POUR LES DASHBOARDS SUMMARY ========================

// OverviewMetrics contient les métriques générales d'aperçu
type OverviewMetrics struct {
	TotalPOS            int64   `json:"total_pos"`
	ActivePOS           int64   `json:"active_pos"`
	TotalVisits         int64   `json:"total_visits"`
	TotalUsers          int64   `json:"total_users"`
	TotalProvinces      int64   `json:"total_provinces"`
	TotalAreas          int64   `json:"total_areas"`
	MarketPenetration   float64 `json:"market_penetration_percent"`
	AverageVisitsPerDay float64 `json:"average_visits_per_day"`
	RevenueGenerated    float64 `json:"revenue_generated"`
	GrowthRate          float64 `json:"growth_rate_percent"`
}

// PerformanceMetrics contient les métriques de performance opérationnelle
type PerformanceMetrics struct {
	VisitObjectiveRate   float64 `json:"visit_objective_rate_percent"`
	CompletionRate       float64 `json:"completion_rate_percent"`
	EfficiencyScore      float64 `json:"efficiency_score"`
	QualityIndex         float64 `json:"quality_index"`
	CustomerSatisfaction float64 `json:"customer_satisfaction_percent"`
	TopBrandPerformance  string  `json:"top_brand_performance"`
	AverageFormScore     float64 `json:"average_form_score"`
	ProductDistribution  float64 `json:"product_distribution_rate"`
}

// GeographicMetrics contient les métriques de distribution géographique
type GeographicMetrics struct {
	TopPerformingProvince string  `json:"top_performing_province"`
	TopPerformingArea     string  `json:"top_performing_area"`
	TopPerformingSubArea  string  `json:"top_performing_subarea"`
	TopPerformingCommune  string  `json:"top_performing_commune"`
	CoveragePercentage    float64 `json:"coverage_percentage"`
	UnderutilizedAreas    int     `json:"underutilized_areas"`
	MarketConcentration   float64 `json:"market_concentration"`
	GeographicEfficiency  float64 `json:"geographic_efficiency"`
}

// TeamPerformanceMetrics contient les métriques de performance d'équipe
type TeamPerformanceMetrics struct {
	TotalTeamMembers      int64   `json:"total_team_members"`
	ActiveTeamMembers     int64   `json:"active_team_members"`
	TopPerformer          string  `json:"top_performer"`
	AverageTeamEfficiency float64 `json:"average_team_efficiency"`
	TeamObjectiveRate     float64 `json:"team_objective_rate_percent"`
	TrainingNeeded        int64   `json:"team_members_needing_training"`
}

// TrendMetrics contient les métriques d'analyse des tendances
type TrendMetrics struct {
	VisitTrend         string  `json:"visit_trend"`
	RevenueTrend       string  `json:"revenue_trend"`
	EfficiencyTrend    string  `json:"efficiency_trend"`
	MonthlyGrowth      float64 `json:"monthly_growth_percent"`
	SeasonalPattern    string  `json:"seasonal_pattern"`
	PredictedNextMonth float64 `json:"predicted_next_month_visits"`
}

// ExecutiveSummaryResponse contient la réponse complète du résumé exécutif
type ExecutiveSummaryResponse struct {
	Overview               OverviewMetrics        `json:"overview"`
	Performance            PerformanceMetrics     `json:"performance"`
	GeographicDistribution GeographicMetrics      `json:"geographic_distribution"`
	TeamPerformance        TeamPerformanceMetrics `json:"team_performance"`
	TrendAnalysis          TrendMetrics           `json:"trend_analysis"`
}

// ======================== STRUCTURES POUR LES RÉSUMÉS RÉGIONAUX ========================

// RegionInfo contient les informations sur une région
type RegionInfo struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	TotalPOS      int64  `json:"total_pos"`
	TotalUsers    int64  `json:"total_users"`
	TotalSubAreas int64  `json:"total_sub_areas"`
	TotalCommunes int64  `json:"total_communes"`
}

// RegionalPerformance contient les métriques de performance régionale
type RegionalPerformance struct {
	VisitsThisPeriod int64   `json:"visits_this_period"`
	ObjectiveRate    float64 `json:"objective_rate_percent"`
	RevenueGenerated float64 `json:"revenue_generated"`
	MarketShare      float64 `json:"market_share_percent"`
	EfficiencyRating string  `json:"efficiency_rating"`
	TrendDirection   string  `json:"trend_direction"`
	GrowthRate       float64 `json:"growth_rate_percent"`
}

// RegionalComparison contient les métriques de comparaison régionale
type RegionalComparison struct {
	RankAmongRegions     int     `json:"rank_among_regions"`
	PerformanceVsAverage float64 `json:"performance_vs_average_percent"`
	BestMetric           string  `json:"best_metric"`
	WeakestMetric        string  `json:"weakest_metric"`
}

// TopPerformer représente un top performer dans une région
type TopPerformer struct {
	Name               string  `json:"name"`
	Role               string  `json:"role"`
	Visits             int     `json:"visits"`
	ObjectiveRate      float64 `json:"objective_rate_percent"`
	SpecialAchievement string  `json:"special_achievement"`
}

// Opportunity représente une opportunité d'amélioration
type Opportunity struct {
	Area            string  `json:"area"`
	Potential       string  `json:"potential"`
	EstimatedImpact float64 `json:"estimated_impact"`
	Effort          string  `json:"effort_required"`
	Timeline        string  `json:"timeline"`
}

// Recommendation représente une recommandation d'action
type Recommendation struct {
	Priority        string `json:"priority"`
	Action          string `json:"action"`
	ExpectedROI     string `json:"expected_roi"`
	Timeline        string `json:"timeline"`
	ResponsibleTeam string `json:"responsible_team"`
}

// RegionalSummaryResponse contient la réponse complète du résumé régional
type RegionalSummaryResponse struct {
	RegionInfo      RegionInfo          `json:"region_info"`
	Performance     RegionalPerformance `json:"performance"`
	Comparison      RegionalComparison  `json:"comparison"`
	TopPerformers   []TopPerformer      `json:"top_performers"`
	Opportunities   []Opportunity       `json:"opportunities"`
	Recommendations []Recommendation    `json:"recommendations"`
}

// ======================== STRUCTURES POUR LE DASHBOARD RAPIDE ========================

// QuickMetrics contient les métriques rapides
type QuickMetrics struct {
	VisitsToday         int64   `json:"visits_today"`
	VisitsThisWeek      int64   `json:"visits_this_week"`
	VisitsThisMonth     int64   `json:"visits_this_month"`
	ObjectiveRateWeek   float64 `json:"objective_rate_week_percent"`
	ActiveUsersToday    int64   `json:"active_users_today"`
	RevenueToday        float64 `json:"revenue_today"`
	CompletionRateToday float64 `json:"completion_rate_today_percent"`
	TopBrandToday       string  `json:"top_brand_today"`
}

// TodayStatistics contient les statistiques du jour
type TodayStatistics struct {
	HourlyVisits     []int   `json:"hourly_visits"`
	PeakHour         int     `json:"peak_hour"`
	ActiveProvinces  int64   `json:"active_provinces"`
	NewPOSVisited    int64   `json:"new_pos_visited"`
	AverageVisitTime float64 `json:"average_visit_duration_minutes"`
}

// UrgentAction représente une action urgente à entreprendre
type UrgentAction struct {
	Priority    string `json:"priority"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Owner       string `json:"owner"`
	Impact      string `json:"impact"`
}

// QuickDashboardResponse contient la réponse complète du dashboard rapide
type QuickDashboardResponse struct {
	LastUpdated   string          `json:"last_updated"`
	KeyMetrics    QuickMetrics    `json:"key_metrics"`
	TodayStats    TodayStatistics `json:"today_stats"`
	UrgentActions []UrgentAction  `json:"urgent_actions"`
}

// ======================== STRUCTURES POUR L'ANALYSE COMPARATIVE ========================

// PeriodAnalysis contient l'analyse d'une période
type PeriodAnalysis struct {
	Name              string  `json:"name"`
	TotalVisits       int64   `json:"total_visits"`
	TotalRevenue      float64 `json:"total_revenue"`
	AveragePrice      float64 `json:"average_price"`
	CompletionRate    float64 `json:"completion_rate_percent"`
	UniqueUsersActive int64   `json:"unique_users_active"`
	UniquePOSVisited  int64   `json:"unique_pos_visited"`
	TopBrand          string  `json:"top_brand"`
	TopProvince       string  `json:"top_province"`
	TopArea           string  `json:"top_area"`
	TopSubArea        string  `json:"top_subarea"`
	TopCommune        string  `json:"top_commune"`
	EfficiencyScore   float64 `json:"efficiency_score"`
}

// ComparisonMetrics contient les métriques de comparaison
type ComparisonMetrics struct {
	VisitGrowth          float64 `json:"visit_growth_percent"`
	RevenueGrowth        float64 `json:"revenue_growth_percent"`
	EfficiencyChange     float64 `json:"efficiency_change_percent"`
	UserEngagementChange float64 `json:"user_engagement_change_percent"`
	MarketExpansion      float64 `json:"market_expansion_percent"`
	OverallPerformance   string  `json:"overall_performance"`
}

// TrendAnalysis contient l'analyse des tendances
type TrendAnalysis struct {
	Metric       string  `json:"metric"`
	Direction    string  `json:"direction"`
	Magnitude    float64 `json:"magnitude"`
	Significance string  `json:"significance"`
	Prediction   string  `json:"prediction"`
}

// CompetitiveInsight représente un insight compétitif
type CompetitiveInsight struct {
	Category    string `json:"category"`
	Finding     string `json:"finding"`
	Implication string `json:"implication"`
	ActionPlan  string `json:"action_plan"`
	Priority    string `json:"priority"`
}

// CompetitiveAnalysisResponse contient la réponse complète de l'analyse comparative
type CompetitiveAnalysisResponse struct {
	CurrentPeriod  PeriodAnalysis       `json:"current_period"`
	PreviousPeriod PeriodAnalysis       `json:"previous_period"`
	Comparison     ComparisonMetrics    `json:"comparison"`
	Trends         []TrendAnalysis      `json:"trends"`
	Insights       []CompetitiveInsight `json:"insights"`
}
