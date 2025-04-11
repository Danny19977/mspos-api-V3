package ruteplanitem

import (
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Paginate
func GetPaginatedRutePlanItem(c *fiber.Ctx) error {
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

	var dataList []models.RutePlanItem
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.RutePlanItem{}).
		Where("name ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Where("name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("RoutePlan").
		Preload("Pos").
		Preload("Status").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch ruteplanitem",
			"error":   err.Error(),
		})
	}

	// Calculate total pages
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
		"message":    "ruteplanitem retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllRutePlanItem(c *fiber.Ctx) error {
	db := database.DB

	var data []models.RutePlanItem
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All RutePlanItems",
		"data":    data,
	})
}

// Create data
func CreateRutePlanItem(c *fiber.Ctx) error {
	p := &models.RutePlanItem{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	p.UUID = uuid.New().String()
	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "RutePlanItem created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateRutePlanItem(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	type UpdateData struct {
		UUID        string `json:"uuid"`
		PosUUID     uint   `json:"pos_uuid" gorm:"type:varchar(255);not null"`
		RoutePlanID uint   `json:"routeplan_uuid"`
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

	rutePlanItem := new(models.RutePlanItem)

	db.Where("uuid = ?", uuid).First(&rutePlanItem)

	db.Save(&rutePlanItem)
	rutePlanItem.PosUUID = updateData.PosUUID
	rutePlanItem.RoutePlanID = updateData.RoutePlanID

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stock updated success",
			"data":    rutePlanItem,
		},
	)

}

// Delete data
func DeleteRutePlanItem(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	db := database.DB

	var rutePlanItems models.RutePlanItem
	db.Where("uuid = ?", uuid).First(&rutePlanItems)
	if rutePlanItems.PosUUID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No stock name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&rutePlanItems)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stock deleted success",
			"data":    nil,
		},
	)
}
