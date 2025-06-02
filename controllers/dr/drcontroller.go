package dr

import (
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
)

// Paginate
func GetPaginatedDr(c *fiber.Ctx) error {
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
		Where("users.role = ?", "DR").
		Where(`
		title ILIKE ? OR EXISTS 
		(SELECT 1 FROM provinces WHERE users.province_uuid = provinces.uuid AND provinces.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM areas WHERE users.area_uuid = areas.uuid AND areas.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM sub_areas WHERE users.sub_area_uuid = sub_areas.uuid AND sub_areas.name ILIKE ?)
		`, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Where("users.role = ?", "DR").
		Where(`
		title ILIKE ? OR EXISTS 
		(SELECT 1 FROM provinces WHERE users.province_uuid = provinces.uuid AND provinces.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM areas WHERE users.area_uuid = areas.uuid AND areas.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM sub_areas WHERE users.sub_area_uuid = sub_areas.uuid AND sub_areas.name ILIKE ?)
		`, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Select(`
			users.*,   
			(
				SELECT COUNT(DISTINCT cyclo_uuid)
				FROM users
				WHERE role = 'ASM' AND province_uuid = users.province_uuid
				AND area_uuid = users.area_uuid 
				AND sub_area_uuid = users.sub_area_uuid
			) AS total_cyclo,
			  (
				SELECT COUNT(DISTINCT p.uuid)
				FROM pos p 
				WHERE users.province_uuid = p.province_uuid
				AND users.area_uuid = p.area_uuid
				AND users.sub_area_uuid = p.sub_area_uuid
			) AS total_pos, 
			(
				SELECT
				COUNT(DISTINCT ps.uuid)
				FROM
				pos_forms ps 
				WHERE
				users.province_uuid = ps.province_uuid
				AND users.area_uuid = ps.area_uuid
				AND users.sub_area_uuid = ps.sub_area_uuid
			) AS total_posforms
		`).
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		// Preload("Pos").
		// Preload("PosForms").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch DRs",
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
		"message":    "DRS retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate by  Province ID
func GetPaginatedDrByProvince(c *fiber.Ctx) error {
	db := database.DB

	UserUUID := c.Params("user_uuid")

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
		Where("users.role = ?", "DR").
		Where("users.asm_uuid = ?", UserUUID).
		Where(`
		title ILIKE ? OR EXISTS 
		(SELECT 1 FROM provinces WHERE users.province_uuid = provinces.uuid AND provinces.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM areas WHERE users.area_uuid = areas.uuid AND areas.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM sub_areas WHERE users.sub_area_uuid = sub_areas.uuid AND sub_areas.name ILIKE ?)
		`, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Where("users.role = ?", "DR").
		Where("users.asm_uuid = ?", UserUUID).
		Where(`
		title ILIKE ? OR EXISTS 
		(SELECT 1 FROM provinces WHERE users.province_uuid = provinces.uuid AND provinces.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM areas WHERE users.area_uuid = areas.uuid AND areas.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM sub_areas WHERE users.sub_area_uuid = sub_areas.uuid AND sub_areas.name ILIKE ?)
		`, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Select(`
			users.*,   
			(
				SELECT COUNT(DISTINCT cyclo_uuid)
				FROM users
				WHERE role = 'DR' AND province_uuid = users.province_uuid
				AND area_uuid = users.area_uuid  
			) AS total_cyclo,
			   (
				SELECT COUNT(DISTINCT p.uuid)
				FROM pos p 
				WHERE users.role = 'DR' AND users.province_uuid = p.province_uuid
				AND users.area_uuid = p.area_uuid
				AND users.sub_area_uuid = p.sub_area_uuid
			) AS total_pos, 
			(
				SELECT
				COUNT(DISTINCT ps.uuid)
				FROM
				pos_forms ps 
				WHERE users.role = 'DR' AND 
				users.province_uuid = ps.province_uuid
				AND users.area_uuid = ps.area_uuid
				AND users.sub_area_uuid = ps.sub_area_uuid
			) AS total_posforms
		`).
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		// Preload("Pos").
		// Preload("PosForms").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch sups",
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
		"message":    "DRS retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate by  Area ID
func GetPaginatedDrByArea(c *fiber.Ctx) error {
	db := database.DB

	UserUUID := c.Params("user_uuid")

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
		Where("users.role = ?", "DR").
		Where("users.sup_uuid = ?", UserUUID). 
		Where(`
		title ILIKE ? OR EXISTS 
		(SELECT 1 FROM provinces WHERE users.province_uuid = provinces.uuid AND provinces.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM areas WHERE users.area_uuid = areas.uuid AND areas.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM sub_areas WHERE users.sub_area_uuid = sub_areas.uuid AND sub_areas.name ILIKE ?)
		`, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Where("users.role = ?", "DR").
		Where("users.sup_uuid = ?", UserUUID). 
		Where(`
		title ILIKE ? OR EXISTS 
		(SELECT 1 FROM provinces WHERE users.province_uuid = provinces.uuid AND provinces.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM areas WHERE users.area_uuid = areas.uuid AND areas.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM sub_areas WHERE users.sub_area_uuid = sub_areas.uuid AND sub_areas.name ILIKE ?)
		`, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Select(`
			users.*,   
			(
				SELECT COUNT(DISTINCT cyclo_uuid)
				FROM users
				WHERE role = 'ASM' AND province_uuid = users.province_uuid
				AND area_uuid = users.area_uuid
			) AS total_cyclo,
			   (
				SELECT COUNT(DISTINCT p.uuid)
				FROM pos p 
				WHERE users.province_uuid = p.province_uuid
				AND users.area_uuid = p.area_uuid
				AND users.sub_area_uuid = p.sub_area_uuid
			) AS total_pos, 
			(
				SELECT
				COUNT(DISTINCT ps.uuid)
				FROM
				pos_forms ps 
				WHERE
				users.province_uuid = ps.province_uuid
				AND users.area_uuid = ps.area_uuid
				AND users.sub_area_uuid = ps.sub_area_uuid
			) AS total_posforms
		`).
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		// Preload("Pos").
		// Preload("PosForms").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch sups",
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
		"message":    "DRS retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate by  SubArea ID
func GetPaginatedDrBySubArea(c *fiber.Ctx) error {
	db := database.DB

	UserUUID := c.Params("user_uuid")

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
	db.Where("users.role = ?", "DR").
		Where("users.dr_uuid = ?", UserUUID). 
		Where(`
		title ILIKE ? OR EXISTS 
		(SELECT 1 FROM provinces WHERE users.province_uuid = provinces.uuid AND provinces.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM areas WHERE users.area_uuid = areas.uuid AND areas.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM sub_areas WHERE users.sub_area_uuid = sub_areas.uuid AND sub_areas.name ILIKE ?)
		`, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Where("users.role = ?", "DR").
		Where("users.dr_uuid = ?", UserUUID). 
		Where(`
		title ILIKE ? OR EXISTS 
		(SELECT 1 FROM provinces WHERE users.province_uuid = provinces.uuid AND provinces.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM areas WHERE users.area_uuid = areas.uuid AND areas.name ILIKE ?) OR EXISTS
		(SELECT 1 FROM sub_areas WHERE users.sub_area_uuid = sub_areas.uuid AND sub_areas.name ILIKE ?)
		`, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Select(`
			users.*,   
			(
				SELECT COUNT(DISTINCT cyclo_uuid)
				FROM users
				WHERE role = 'ASM' AND province_uuid = users.province_uuid
				AND area_uuid = users.area_uuid
			) AS total_cyclo,
			   (
				SELECT COUNT(DISTINCT p.uuid)
				FROM pos p 
				WHERE users.province_uuid = p.province_uuid
				AND users.area_uuid = p.area_uuid
				AND users.sub_area_uuid = p.sub_area_uuid
			) AS total_pos, 
			(
				SELECT
				COUNT(DISTINCT ps.uuid)
				FROM
				pos_forms ps 
				WHERE
				users.province_uuid = ps.province_uuid
				AND users.area_uuid = ps.area_uuid
				AND users.sub_area_uuid = ps.sub_area_uuid
			) AS total_posforms
		`).
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		// Preload("Pos").
		// Preload("PosForms").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch sups",
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
		"message":    "DRS retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}
