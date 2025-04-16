package cyclo

import (
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Paginate
func GetPaginatedCyclo(c *fiber.Ctx) error {
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

	var dataList []models.Cyclo
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Cyclo{}).
		Joins("JOIN users ON cyclos.user_uuid=users.uuid").
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN users ON cyclos.user_uuid=users.uuid").
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("cyclos.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("Asm").
		Preload("Sup").
		Preload("Dr"). 
		// Preload("User").
		Preload("Pos").
		Preload("PosForms").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch cyclos",
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
		"message":    "Cyclo retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate Province by ID
func GetPaginatedCycloProvinceByID(c *fiber.Ctx) error {
	db := database.DB

	ProvinceUUID := c.Params("province_uuid")

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

	var dataList []models.Cyclo
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Cyclo{}).
		Joins("JOIN users ON cyclos.user_uuid=users.uuid").
		Where("cyclos.province_uuid = ?", ProvinceUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN users ON cyclos.user_uuid=users.uuid").
		Where("cyclos.province_uuid = ?", ProvinceUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("cyclos.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("Asm").
		Preload("Sup").
		Preload("Dr"). 
		// Preload("User").
		Preload("Pos").
		Preload("PosForms").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch cyclos",
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
		"message":    "Cyclo retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate  Area by ID
func GetPaginatedCycloByAreaUUID(c *fiber.Ctx) error {
	db := database.DB

	AreaUUID := c.Params("area_uuid")

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

	var dataList []models.Cyclo
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Cyclo{}).
		Joins("JOIN users ON cyclos.user_uuid=users.uuid").
		Where("cyclos.area_uuid = ?", AreaUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN users ON cyclos.user_uuid=users.uuid").
		Where("cyclos.area_uuid = ?", AreaUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("cyclos.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("Asm").
		Preload("Sup").
		Preload("Dr"). 
		// Preload("User").
		Preload("Pos").
		Preload("PosForms").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch cyclos",
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
		"message":    "Cyclo retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate SubArea by ID
func GetPaginatedSubAreaByID(c *fiber.Ctx) error {
	db := database.DB

	subAreaUUID := c.Params("subarea_uuid")

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

	var dataList []models.Cyclo
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Cyclo{}).
		Joins("JOIN users ON cyclos.user_uuid=users.uuid").
		Where("cyclos.subarea_uuid = ?", subAreaUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN users ON cyclos.user_uuid=users.uuid").
		Where("cyclos.subarea_uuid = ?", subAreaUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("cyclos.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("Asm").
		Preload("Sup").
		Preload("Dr"). 
		// Preload("User").
		Preload("Pos").
		Preload("PosForms").
		Preload("PosForms").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch cyclos",
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
		"message":    "Cyclo retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate by Commune ID
func GetPaginatedCycloCommune(c *fiber.Ctx) error {
	db := database.DB

	CommuneUUID := c.Params("commune_uuid")

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

	var dataList []models.Cyclo
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Cyclo{}).
		Joins("JOIN users ON cyclos.user_uuid=users.uuid").
		Where("cyclos.commune_uuid = ?", CommuneUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN users ON cyclos.user_uuid=users.uuid").
		Where("cyclos.commune_uuid = ?", CommuneUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("cyclos.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("Asm").
		Preload("Sup").
		Preload("Dr"). 
		// Preload("User").
		Preload("Pos").
		Preload("PosForms").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch cyclos",
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
		"message":    "Cyclo retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllCyclo(c *fiber.Ctx) error {
	db := database.DB

	var data []models.Cyclo
	db.
		Preload("Province").
		Preload("Area").
		Preload("Asm").
		Preload("Sup").
		Preload("Dr").
		Order("Cyclo.updated_at DESC").
		Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Cyclos",
		"data":    data,
	})
}

// Get one data
func GetOneCyclo(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	var cyclo models.Cyclo
	db.Where("uuid = ?", uuid).First(&cyclo)
	if cyclo.Signature == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No cyclo signature found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "cyclo found",
			"data":    cyclo,
		},
	)
}

// Create data
func CreateCyclo(c *fiber.Ctx) error {
	p := &models.Cyclo{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	p.UUID = uuid.New().String()
	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Cyclo created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateCyclo(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	type UpdateData struct {
		UUID string `json:"uuid"`

		CountryUUID  string `json:"country_uuid" gorm:"type:varchar(255);not null"`
		ProvinceUUID string `json:"province_uuid" gorm:"type:varchar(255);not null"`
		SubAreaUUID  string `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
		CommuneUUID  string `json:"commune_uuid" gorm:"type:varchar(255);not null"`
		AsmUUID      string `json:"asm_uuid" gorm:"type:varchar(255);not null"`
		SupUUID      string `json:"sup_uuid" gorm:"type:varchar(255);not null"`
		DrUUID       string `json:"dr_uuid" gorm:"type:varchar(255);not null"`
		Signature    string `json:"signature"`
		UserUUID     string `json:"user_uuid"`
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

	cyclo := new(models.Cyclo)

	db.Where("uuid = ?", uuid).First(&cyclo)
	cyclo.CountryUUID = updateData.CountryUUID
	cyclo.ProvinceUUID = updateData.ProvinceUUID
	cyclo.SubAreaUUID = updateData.SubAreaUUID
	cyclo.AsmUUID = updateData.AsmUUID
	cyclo.SupUUID = updateData.SupUUID
	cyclo.DrUUID = updateData.DrUUID
	cyclo.Signature = updateData.Signature
	cyclo.UserUUID = updateData.UserUUID

	db.Save(&cyclo)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "cyclo updated success",
			"data":    cyclo,
		},
	)

}

// Delete data
func DeleteCyclo(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	db := database.DB

	var cyclo models.Cyclo
	db.Where("uuid = ?", uuid).First(&cyclo)
	if cyclo.Signature == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Cyclo name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&cyclo)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "cyclo deleted success",
			"data":    nil,
		},
	)
}
