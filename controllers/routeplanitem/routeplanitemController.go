package RoutePlanItem

import (
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/danny19977/mspos-api-v3/utils"
	"github.com/gofiber/fiber/v2"
)

// Paginate
func GetPaginatedRoutePlanItem(c *fiber.Ctx) error {
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

	var dataList []models.RoutePlanItem
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.RoutePlanItem{}).
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
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch RoutePlanItem",
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
		"message":    "RoutePlanItem retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllRoutePlanItem(c *fiber.Ctx) error {
	db := database.DB

	routePlanUUID := c.Params("route_plan_uuid")

	var dataList []models.RoutePlanItem
	db.
		Where("route_plan_uuid = ?", routePlanUUID).
		Preload("RoutePlan").
		Preload("Pos").
		Find(&dataList)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All RoutePlanItems",
		"data":    dataList,
	})
}

// Get One data
func GetOneByRouteItermUUID(c *fiber.Ctx) error {
	UUID := c.Params("uuid")
	db := database.DB

	var routePlanItem models.RoutePlanItem

	// Fetch the RoutePlanItem by route UUID
	err := db.Where("uuid = ?", UUID).
		Preload("RoutePlan").
		Preload("Pos").
		First(&routePlanItem).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "RoutePlanItem not found",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "RoutePlanItem retrieved successfully",
		"data":    routePlanItem,
	})
}

// Create data
func CreateRoutePlanItem(c *fiber.Ctx) error {
	p := &models.RoutePlanItem{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	p.UUID = utils.GenerateUUID()
	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "RoutePlanItem created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateRoutePlanItem(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	type UpdateData struct {
		Status bool `json:"status"`
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

	RoutePlanItem := new(models.RoutePlanItem)

	db.Where("uuid = ?", uuid).First(&RoutePlanItem)

	RoutePlanItem.Status = updateData.Status

	db.Save(&RoutePlanItem)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stock updated success",
			"data":    RoutePlanItem,
		},
	)
}

// Delete data
func DeleteRoutePlanItem(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	db := database.DB

	var RoutePlanItems models.RoutePlanItem
	db.Where("uuid = ?", uuid).First(&RoutePlanItems)
	if RoutePlanItems.PosUUID == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No stock name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&RoutePlanItems)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stock deleted success",
			"data":    nil,
		},
	)
}
