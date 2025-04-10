package sup

import (
	"fmt"
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Paginate
func GetPaginatedSups(c *fiber.Ctx) error {
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

	var dataList []models.Sup
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Sup{}).
		Joins("JOIN users ON sups.user_id=users.id").
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN countries ON sups.country_uuid=countries.id").
		Joins("JOIN provinces ON sups.province_uuid=provinces.id").
		Joins("JOIN areas ON sups.area_uuid=areas.id").
		Joins("JOIN asms ON sups.asm_uuid=asms.id").
		Joins("JOIN users ON sups.user_id=users.id").
		Where("countries.name ILIKE ? OR provinces.name ILIKE ? OR areas.name ILIKE ? OR users.fullname ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("sups.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("Asm.User").
		// Preload("User").
		Preload("Drs").
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
		"message":    "Sups retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate by Province ID
func GetPaginatedProvince(c *fiber.Ctx) error {
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
	fmt.Println("Search query:", search)

	var dataList []models.Sup
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Sup{}).
		Joins("JOIN users ON sups.user_id=users.id").
		Where("sups.province_uuid = ?", province_uuid).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN countries ON sups.country_uuid=countries.id").
		Joins("JOIN provinces ON sups.province_uuid=provinces.id").
		Joins("JOIN areas ON sups.area_uuid=areas.id").
		Joins("JOIN asms ON sups.asm_uuid=asms.id").
		Joins("JOIN users ON sups.user_id=users.id").
		Where("sups.province_uuid = ?", province_uuid).
		Where("countries.name ILIKE ? OR provinces.name ILIKE ? OR areas.name ILIKE ? OR users.fullname ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("sups.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("Asm.User").
		// Preload("User").
		Preload("Drs").
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
		"message":    "Sups retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate by Area ID
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
	fmt.Println("Search query:", search)

	var dataList []models.Sup
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Sup{}).
		Joins("JOIN users ON sups.user_id=users.id").
		Where("sups.area_uuid = ?", area_uuid).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN countries ON sups.country_uuid=countries.id").
		Joins("JOIN provinces ON sups.province_uuid=provinces.id").
		Joins("JOIN areas ON sups.area_uuid=areas.id").
		Joins("JOIN asms ON sups.asm_uuid=asms.id").
		Joins("JOIN users ON sups.user_id=users.id").
		Where("sups.area_uuid = ?", area_uuid).
		Where("countries.name ILIKE ? OR provinces.name ILIKE ? OR areas.name ILIKE ? OR users.fullname ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("sups.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("Asm.User").
		// Preload("User").
		Preload("Drs").
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
		"message":    "Sups retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All Sups
func GetAllSups(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Sup
	db.
		Preload("User").
		Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All sups",
		"data":    data,
	})
}

// Total of DR by Sup ID
func GetDrByID(c *fiber.Ctx) error {
	sup_uuid := c.Params("sup_uuid")
	db := database.DB

	var dr []models.Dr
	var count int64
	db.Where("sup_uuid = ?", sup_uuid).Find(&dr).Count(&count)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Count DR by Sup id found",
		"data":    count,
	})
}

// Total of Cyclo by Cyclo ID
func GetCycloByID(c *fiber.Ctx) error {
	sup_uuid := c.Params("sup_uuid")
	db := database.DB

	var cyclo []models.Cyclo
	var count int64
	db.Where("sup_uuid = ?", sup_uuid).Find(&cyclo).Count(&count)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Count Cyclo by Sup id found",
		"data":    count,
	})
}

// Total of POS by Sup ID
func GetPosByID(c *fiber.Ctx) error {
	sup_uuid := c.Params("sup_uuid")
	db := database.DB

	var pos []models.Pos
	var count int64
	db.Where("sup_uuid = ?", sup_uuid).Find(&pos).Count(&count)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Count POS by Sup id found",
		"data":    count,
	})
}

// Get one data
func GetSup(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB
	var sup models.Sup
	db.Where("uuid = ?", uuid).First(&sup)
	if sup.ProvinceUUID == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No sup name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "sup found",
			"data":    sup,
		},
	)
}

// Create data
func CreateSup(c *fiber.Ctx) error {
	p := &models.Sup{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}
	p.UUID = uuid.New().String()
	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Sup created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateSup(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	type UpdateData struct {
		UUID string `json:"uuid"`

		CountryUUID  string `json:"country_uuid" gorm:"type:varchar(255);not null"`
		ProvinceUUID string `json:"province_uuid" gorm:"type:varchar(255);not null"`
		AreaUUID     string `json:"area_uuid" gorm:"type:varchar(255);not null"`
		AsmUUID      string `json:"asm_uuid" gorm:"type:varchar(255);not null"`
		UserUUID     string `json:"user_id"`
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

	sup := new(models.Sup)

	db.Where("uuid = ?", uuid).First(&sup)
	sup.CountryUUID = updateData.CountryUUID
	sup.ProvinceUUID = updateData.ProvinceUUID
	sup.AreaUUID = updateData.AreaUUID
	sup.AsmUUID = updateData.AsmUUID
	sup.UserUUID = updateData.UserUUID
	sup.Signature = updateData.Signature

	db.Save(&sup)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "sup updated success",
			"data":    sup,
		},
	)

}

// Delete data
func DeleteSup(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	db := database.DB

	var sup models.Sup
	db.Where("uuid = ?", uuid).First(&sup)
	if sup.ProvinceUUID == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No sup name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&sup)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "sup deleted success",
			"data":    nil,
		},
	)
}
