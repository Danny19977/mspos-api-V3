package dr

import (
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	var dataList []models.Dr
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Dr{}).
		Joins("JOIN users ON drs.user_uuid=users.uuid").
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN users ON drs.user_uuid=users.uuid").
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("drs.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Asm.User").
		Preload("Sup.User").
		// Preload("User").
		Preload("Cyclos").
		Preload("Pos").
		Preload("PosForms").
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

// Paginate by  Province ID
func GetPaginatedDrByProvince(c *fiber.Ctx) error {
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

	var dataList []models.Dr
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Dr{}).
		Joins("JOIN users ON drs.user_uuid=users.uuid").
		Where("drs.province_uuid = ?", province_uuid).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN users ON drs.user_uuid=users.uuid").
		Where("drs.province_uuid = ?", province_uuid).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("drs.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Asm.User").
		Preload("Sup.User").
		// Preload("User").
		Preload("Cyclos").
		Preload("Pos").
		Preload("PosForms").
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
func GetPaginatedArea(c *fiber.Ctx) error {
	db := database.DB

	area_uuid := c.Params("area_uuid")

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

	var dataList []models.Dr
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Dr{}).
		Joins("JOIN users ON drs.user_uuid=users.uuid").
		Where("drs.area_uuid = ?", area_uuid).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN users ON drs.user_uuid=users.uuid").
		Where("drs.area_uuid = ?", area_uuid).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("drs.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Asm.User").
		Preload("Sup.User").
		// Preload("User").
		Preload("Cyclos").
		Preload("Pos").
		Preload("PosForms").
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
func GetPaginatedSubArea(c *fiber.Ctx) error {
	db := database.DB

	subarea_uuid := c.Params("subarea_uuid")

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

	var dataList []models.Dr
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Dr{}).
		Joins("JOIN users ON drs.user_uuid=users.uuid").
		Where("drs.subarea_uuid = ?", subarea_uuid).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN users ON drs.user_uuid=users.uuid").
		Where("drs.subarea_uuid = ?", subarea_uuid).
		Where("countries.name ILIKE ? OR provinces.name ILIKE ? OR areas.name ILIKE ? OR users.fullname ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("drs.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Asm.User").
		Preload("Sup.User").
		// Preload("User").
		Preload("Cyclos").
		Preload("Pos").
		Preload("PosForms").
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

// Get All data
func GetAllDr(c *fiber.Ctx) error {
	db := database.DB

	var data []models.Dr
	db.
		Preload("User").
		Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Drs",
		"data":    data,
	})
}

// Get one data
func GetOneDr(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	var Dr models.Dr
	db.Where("uuid = ?", uuid).
		Preload("User").
		First(&Dr)
	if Dr.Signature == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Dr signature found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Dr found",
			"data":    Dr,
		},
	)
}

// Create data
func CreateDr(c *fiber.Ctx) error {
	p := &models.Dr{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	p.UUID = uuid.New().String()
	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Dr created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateDr(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	type UpdateData struct {
		UUID string `json:"uuid"`

		CountryUUID  string `json:"country_uuid" gorm:"type:varchar(255);not null"`
		ProvinceUUID string `json:"province_uuid" gorm:"type:varchar(255);not null"`
		SubAreaUUID  string `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
		// AsmUUID      uint   `json:"asm_uuid" gorm:"type:varchar(255);not null"`
		// SupUUID      uint   `json:"sup_uuid" gorm:"type:varchar(255);not null"`
		Signature string `json:"signature"`
		UserUUID  string `json:"user_uuid"`
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

	Dr := new(models.Dr)

	db.Where("uuid = ?", uuid).First(&Dr)
	Dr.CountryUUID = updateData.CountryUUID
	Dr.ProvinceUUID = updateData.ProvinceUUID
	Dr.SubAreaUUID = updateData.SubAreaUUID
	// Dr.AsmUUID = updateData.AsmUUID
	// Dr.SupUUID = updateData.SupUUID
	Dr.Signature = updateData.Signature
	Dr.UserUUID = updateData.UserUUID

	db.Save(&Dr)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Dr updated success",
			"data":    Dr,
		},
	)

}

// Delete data
func DeleteDr(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	db := database.DB

	var Dr models.Dr
	db.Where("uuid = ?", uuid).First(&Dr)
	if Dr.Signature == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Dr name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&Dr)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Dr deleted success",
			"data":    nil,
		},
	)
}
