package asm

import ( 
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
)

// Paginate
func GetPaginatedASM(c *fiber.Ctx) error {
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

	var dataList []models.User
	var totalRecords int64

	// Count total records matching the search query
	db.
		Where("users.role = ?", "ASM").
		Where("fullname ILIKE ? ", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Where("users.role = ?", "ASM"). 
		Where("fullname ILIKE ? ", "%"+search+"%").
		Select(` 
			users.*, 
			(
				SELECT COUNT(DISTINCT u2.sup_uuid)
				FROM users u2
				WHERE u2.role = 'Supervisor' AND u2.province_uuid = users.province_uuid
			) AS total_sup,
			(
				SELECT COUNT(DISTINCT u2.dr_uuid)
				FROM users u2
				WHERE u2.role = 'DR' AND u2.province_uuid = users.province_uuid
			) AS total_dr,
			(
				SELECT COUNT(DISTINCT u2.cyclo_uuid)
				FROM users u2
				WHERE u2.role = 'Cyclo' AND u2.province_uuid = users.province_uuid
			) AS total_cyclo, 
			(
				SELECT COUNT(DISTINCT p.uuid)
				FROM pos p 
				WHERE users.province_uuid = p.province_uuid
			) AS total_pos, 
			(
				SELECT
				COUNT(DISTINCT ps.uuid)
				FROM
				pos_forms ps 
				WHERE
				users.province_uuid = ps.province_uuid
			) AS total_posforms
		`).
		Offset(offset).
		Limit(limit).
		Order("users.updated_at DESC").
		Preload("Country").
		Preload("Province").
		// Preload("Pos").
		// Preload("PosForms").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch asms",
			"error":   err.Error(),
		})
	}

	// Calculate total pages
	totalPages := int((totalRecords + int64(limit) - 1) / int64(limit))

	// Prepare pagination metadata
	pagination := map[string]any{
		"total_records": totalRecords,
		"total_pages":   totalPages,
		"current_page":  page,
		"page_size":     limit,
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "Asm retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data Province
func GetPaginatedASMByProvince(c *fiber.Ctx) error {
	db := database.DB

	province_uuid := c.Params("province_uuid")

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

	var dataList []models.User
	var totalRecords int64

	// Count total records matching the search query
	db.
		Where("users.role = ?", "ASM").
		Where("province_uuid = ?", province_uuid).
		Where("fullname ILIKE ? ", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Where("users.role = ?", "ASM").
		Where("province_uuid = ?", province_uuid).
		Where("fullname ILIKE ? ", "%"+search+"%").
		Select(` 
			users.*, 
			(
				SELECT COUNT(DISTINCT u2.sup_uuid)
				FROM users u2
				WHERE u2.role = 'Supervisor' AND u2.province_uuid = users.province_uuid
			) AS total_sup,
			(
				SELECT COUNT(DISTINCT u2.dr_uuid)
				FROM users u2
				WHERE u2.role = 'DR' AND u2.province_uuid = users.province_uuid
			) AS total_dr,
			(
				SELECT COUNT(DISTINCT u2.cyclo_uuid)
				FROM users u2
				WHERE u2.role = 'Cyclo' AND u2.province_uuid = users.province_uuid
			) AS total_cyclo, 
			(
				SELECT COUNT(DISTINCT p.uuid)
				FROM pos p 
				WHERE users.province_uuid = p.province_uuid
			) AS total_pos, 
			(
				SELECT
				COUNT(DISTINCT ps.uuid)
				FROM
				pos_forms ps 
				WHERE
				users.province_uuid = ps.province_uuid
			) AS total_posforms
		`).
		Offset(offset).
		Limit(limit).
		Order("users.updated_at DESC").
		Preload("Country").
		Preload("Province").
		// Preload("Pos").
		// Preload("PosForms").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch asms",
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
		"message":    "Provinces retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}
