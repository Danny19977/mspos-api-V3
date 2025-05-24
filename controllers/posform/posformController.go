package posform

import (
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
)

// Paginate ALL data
func GetPaginatedPosForm(c *fiber.Ctx) error {
	db := database.DB

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	// Provide default values if start_date or end_date are empty
	if start_date == "" {
		start_date = "1970-01-01T00:00:00Z" // Default start date
	}
	if end_date == "" {
		end_date = "2100-01-01T00:00:00Z" // Default end date
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	limit, err := strconv.Atoi(c.Query("limit", "15"))
	if err != nil || limit <= 0 {
		limit = 15
	}
	offset := (page - 1) * limit

	// Deferent filter
	search := c.Query("search", "")

	var dataList []models.PosForm
	var totalRecords int64

	db.Model(&models.PosForm{}).
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.name ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where(" pos.name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("pos_forms.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("ASM").
		Preload("Sup").
		Preload("Dr").
		Preload("Cyclo").
		Preload("Users").
		Preload("Pos").
		Preload("PosFormItems").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch provinces",
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
		"message":    "POSFORM retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data province by UUID
func GetPaginatedPosFormProvine(c *fiber.Ctx) error {
	db := database.DB

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	// Provide default values if start_date or end_date are empty
	if start_date == "" {
		start_date = "1970-01-01T00:00:00Z" // Default start date
	}
	if end_date == "" {
		end_date = "2100-01-01T00:00:00Z" // Default end date
	}

	provinceUUID := c.Params("province_uuid")

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	limit, err := strconv.Atoi(c.Query("limit", "15"))
	if err != nil || limit <= 0 {
		limit = 15
	}
	offset := (page - 1) * limit

	// Deferent filter
	search := c.Query("search", "")

	var dataList []models.PosForm
	var totalRecords int64

	db.Model(&models.PosForm{}).
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.province_uuid = ?", provinceUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.name ILIKE ?", "%"+search+"%")

	err = db.
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.province_uuid = ?", provinceUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("pos_forms.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("ASM").
		Preload("Sup").
		Preload("Dr").
		Preload("Cyclo").
		Preload("Pos").
		Preload("PosFormItems").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch posforms",
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
		"message":    "posforms retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data area by UUID
func GetPaginatedPosFormArea(c *fiber.Ctx) error {
	db := database.DB

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	// Provide default values if start_date or end_date are empty
	if start_date == "" {
		start_date = "1970-01-01T00:00:00Z" // Default start date
	}
	if end_date == "" {
		end_date = "2100-01-01T00:00:00Z" // Default end date
	}

	areaUUID := c.Params("area_uuid")

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	limit, err := strconv.Atoi(c.Query("limit", "15"))
	if err != nil || limit <= 0 {
		limit = 15
	}
	offset := (page - 1) * limit

	// Deferent filter
	search := c.Query("search", "")

	var dataList []models.PosForm
	var totalRecords int64

	db.Model(&models.PosForm{}).
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.area_uuid = ?", areaUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where(" pos.name ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.area_uuid = ?", areaUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where(" pos.name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("pos_forms.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("ASM").
		Preload("Sup").
		Preload("Dr").
		Preload("Cyclo").
		Preload("Pos").
		Preload("PosFormItems").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch provinces",
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
		"message":    "posform retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data subarea by UUID
func GetPaginatedPosFormSubArea(c *fiber.Ctx) error {
	db := database.DB

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	// Provide default values if start_date or end_date are empty
	if start_date == "" {
		start_date = "1970-01-01T00:00:00Z" // Default start date
	}
	if end_date == "" {
		end_date = "2100-01-01T00:00:00Z" // Default end date
	}

	DrUUID := c.Params("dr_uuid")

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	limit, err := strconv.Atoi(c.Query("limit", "15"))
	if err != nil || limit <= 0 {
		limit = 15
	}
	offset := (page - 1) * limit

	// Deferent filter
	search := c.Query("search", "")

	var dataList []models.PosForm
	var totalRecords int64

	db.Model(&models.PosForm{}).
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.dr_uuid = ?", DrUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where(" pos.name ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.dr_uuid = ?", DrUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where(" pos.name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("pos_forms.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("ASM").
		Preload("Sup").
		Preload("Dr").
		Preload("Cyclo").
		Preload("Pos").
		Preload("PosFormItems").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch provinces",
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
		"message":    "posform retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data commune by UUID
func GetPaginatedPosFormCommune(c *fiber.Ctx) error {
	db := database.DB

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	// Provide default values if start_date or end_date are empty
	if start_date == "" {
		start_date = "1970-01-01T00:00:00Z" // Default start date
	}
	if end_date == "" {
		end_date = "2100-01-01T00:00:00Z" // Default end date
	}

	CycloUUID := c.Params("cyclo_uuid")

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	limit, err := strconv.Atoi(c.Query("limit", "15"))
	if err != nil || limit <= 0 {
		limit = 15
	}
	offset := (page - 1) * limit

	// Deferent filter
	search := c.Query("search", "")

	var dataList []models.PosForm
	var totalRecords int64

	db.Model(&models.PosForm{}).
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.cyclo_uuid = ?", CycloUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where(" pos.name ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.cyclo_uuid = ?", CycloUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where(" pos.name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("pos_forms.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("ASM.User").
		Preload("Sup.User").
		Preload("Dr.User").
		Preload("Cyclo.User").
		Preload("Pos").
		Preload("PosFormItems").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch provinces",
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
		"message":    "posform retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data pos by UUID
func GetPaginatedPosFormByPOS(c *fiber.Ctx) error {
	db := database.DB

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	// Provide default values if start_date or end_date are empty
	if start_date == "" {
		start_date = "1970-01-01T00:00:00Z" // Default start date
	}
	if end_date == "" {
		end_date = "2100-01-01T00:00:00Z" // Default end date
	}

	posUUID := c.Params("pos_uuid")

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	limit, err := strconv.Atoi(c.Query("limit", "15"))
	if err != nil || limit <= 0 {
		limit = 15
	}
	offset := (page - 1) * limit

	// Deferent filter
	search := c.Query("search", "")

	var dataList []models.PosForm
	var totalRecords int64

	db.Model(&models.PosForm{}).
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.pos_uuid = ?", posUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.name ILIKE ?", "%"+search+"%")

	err = db.
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.pos_uuid = ?", posUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos.name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("pos_forms.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("ASM").
		Preload("Sup").
		Preload("Dr").
		Preload("Cyclo").
		Preload("Pos").
		Preload("PosFormItems").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch posforms",
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
		"message":    "posforms retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllPosforms(c *fiber.Ctx) error {
	db := database.DB
	var data []models.PosForm
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All PosForms",
		"data":    data,
	})
}

// Get one data
func GetPosform(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB
	var posform models.PosForm
	db.Where("uuid = ?", uuid).First(&posform)
	if posform.ID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No posform name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "posform found",
			"data":    posform,
		},
	)
}

// Create data
func CreatePosform(c *fiber.Ctx) error {
	p := &models.PosForm{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	// p.UUID = uuid.New().String()
	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "posForm created success",
			"data":    p,
		},
	)
}

// Update data
func UpdatePosform(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	type UpdateData struct {
		Price   int    `gorm:"default:0" json:"price"`
		Comment string `json:"comment"`

		Latitude  float64 `json:"latitude"`  // Latitude of the user
		Longitude float64 `json:"longitude"` // Longitude of the user
		Signature string  `json:"signature"`

		PosUUID      string `json:"pos_uuid" gorm:"type:varchar(255);not null"`
		ProvinceUUID string `json:"province_uuid" gorm:"type:varchar(255);not null"`
		AreaUUID     string `json:"area_uuid" gorm:"type:varchar(255);not null"`
		SubAreaUUID  string `json:"sub_area_uuid" gorm:"type:varchar(255);not null"`

		

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

	posform := new(models.PosForm)

	db.Where("uuid = ?", uuid).First(&posform)

	posform.Price = updateData.Price
	posform.Comment = updateData.Comment
	posform.Latitude = updateData.Latitude
	posform.Longitude = updateData.Longitude
	posform.Signature = updateData.Signature
	posform.PosUUID = updateData.PosUUID
	posform.ProvinceUUID = updateData.ProvinceUUID
	posform.AreaUUID = updateData.AreaUUID
	posform.SubAreaUUID = updateData.SubAreaUUID  

	db.Save(&posform)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "posform updated success",
			"data":    posform,
		},
	)

}

// Delete data
func DeletePosform(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	db := database.DB

	var posform models.PosForm
	db.Where("uuid = ?", uuid).First(&posform)
	if posform.ID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No posform name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&posform)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "posform deleted success",
			"data":    nil,
		},
	)
}
