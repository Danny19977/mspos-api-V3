package routeplan

import (
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	// "github.com/lib/pq"
)

// Paginate
func GetPaginatedRouthplan(c *fiber.Ctx) error {
	db := database.DB
	// Parse query parameters for pagination
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit", "15"))
	if err != nil || limit <= 0 {
		limit = 15
	}
	offset := (page - 1) * limit

	// Parse search query
	search := c.Query("search", "")

	var dataList []models.RoutePlan
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.RoutePlan{}).
		Where("route_plans.name ILIKE ? OR users.name ILIKE ? OR provinces.name ILIKE ? OR areas.name ILIKE ? OR subareas.name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Where("name ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Where("route_plans.name ILIKE ? OR users.name ILIKE ? OR provinces.name ILIKE ? OR areas.name ILIKE ? OR subareas.name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Where("name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("route_plans.updated_at DESC").
		Preload("User").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("RutePlanItems").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch provinces",
			"error":   err.Error(),
		})
	}

	/// Calculate total pages
	totalPages := int((totalRecords + int64(limit) - 1) / int64(limit))

	// Prepare pagination metadata
	pagination := map[string]interface{}{
		"total_records": totalRecords,
		"total_pages":   totalPages,
		"current_page":  page,
		"page_size":     limit,
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "Routhplan retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllRouthplan(c *fiber.Ctx) error {
	db := database.DB

	var data []models.RoutePlan
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Routhplan",
		"data":    data,
	})
}

// Get All data by id
func GetAllRouthplanBySearch(c *fiber.Ctx) error {
	db := database.DB

	search := c.Query("search", "")

	var data []models.RoutePlan
	db.Where("name ILIKE ?", "%"+search+"%").
		Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Routhplan",
		"data":    data,
	})
}

// Get one data
func GetRouthplan(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	var Routeplan models.RoutePlan
	db.Where("uuid = ?", uuid).First(&Routeplan)
	if Routeplan.UUID == "0" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Routeplan name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "RoutePlan found",
			"data":    Routeplan,
		},
	)
}

// Create data
func CreateRouthplan(c *fiber.Ctx) error {
	p := &models.RoutePlan{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	p.UUID = uuid.New().String()
	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "routeplan created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateRouthplan(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	type UpdateData struct {
		UUID string `json:"uuid"`

		UserUUID     string   `json:"user_uuid"`
		ProvinceUUID string `json:"province_uuid" gorm:"type:varchar(255);not null"`
		SubAreaUUID  string `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
		TotalPOS     int    `json:"total_pos"`
		Signature    string `json:"signature"`
	}

	var updateData UpdateData
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Review your iunput",
				"data":    nil,
			},
		)
	}

	RoutePlan := new(models.RoutePlan)

	db.Where("uuid = ?", uuid).First(&RoutePlan)
	RoutePlan.UserUUID = updateData.UserUUID
	RoutePlan.ProvinceUUID = updateData.ProvinceUUID
	RoutePlan.SubAreaUUID = updateData.SubAreaUUID
	// RoutePlan.TotalPOS = updateData.TotalPOS
	RoutePlan.Signature = updateData.Signature

	db.Save(&RoutePlan)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "RoutePlan updated success",
			"data":    RoutePlan,
		},
	)

}

// Delete data
func DeleteRouthplan(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	db := database.DB

	var routeplan models.RoutePlan
	db.Where("uuid = ?", uuid).First(&routeplan)
	if routeplan.UUID == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No routeplan name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&routeplan)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "RoutePlan deleted success",
			"data":    nil,
		},
	)
}
