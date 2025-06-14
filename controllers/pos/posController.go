package pos

import (
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Paginate
func GetPaginatedPos(c *fiber.Ctx) error {
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

	var dataList []models.Pos
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Pos{}).
		Where("name ILIKE ? OR shop ILIKE ? OR postype ILIKE ? OR gerant ILIKE ? OR quartier ILIKE ? OR reference ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Where("name ILIKE ? OR shop ILIKE ? OR postype ILIKE ? OR gerant ILIKE ? OR quartier ILIKE ? OR reference ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("User").
		Preload("PosForms").
		Preload("PosEquipments").
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
		"message":    "POS retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate by province id
func GetPaginatedPosByProvinceUUID(c *fiber.Ctx) error {
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

	var dataList []models.Pos
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Pos{}).
		Where("pos.province_uuid = ?", ProvinceUUID).
		Where("name ILIKE ? OR shop ILIKE ? OR postype ILIKE ? OR gerant ILIKE ? OR quartier ILIKE ? OR reference ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Where("pos.province_uuid = ?", ProvinceUUID).
		Where("name ILIKE ? OR shop ILIKE ? OR postype ILIKE ? OR gerant ILIKE ? OR quartier ILIKE ? OR reference ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("User").
		Preload("PosForms").
		Preload("PosEquipments").
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
		"message":    "POS retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate by area id
func GetPaginatedPosByAreaUUID(c *fiber.Ctx) error {
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

	var dataList []models.Pos
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Pos{}).
		Where("pos.area_uuid = ?", AreaUUID).
		Where("name ILIKE ? OR shop ILIKE ? OR postype ILIKE ? OR gerant ILIKE ? OR quartier ILIKE ? OR reference ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Where("pos.area_uuid = ?", AreaUUID).
		Where("name ILIKE ? OR shop ILIKE ? OR postype ILIKE ? OR gerant ILIKE ? OR quartier ILIKE ? OR reference ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("User").
		Preload("PosEquipments").
		Preload("PosForms").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch Area",
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
		"message":    "POS retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate by SubArea id
func GetPaginatedPosBySubAreaUUID(c *fiber.Ctx) error {
	db := database.DB

	SubAreaUUID := c.Params("sub_area_uuid")

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

	var dataList []models.Pos
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Pos{}).
		Where("pos.sub_area_uuid = ?", SubAreaUUID).
		Where("name ILIKE ? OR shop ILIKE ? OR postype ILIKE ? OR gerant ILIKE ? OR quartier ILIKE ? OR reference ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Where("pos.sub_area_uuid = ?", SubAreaUUID).
		Where("name ILIKE ? OR shop ILIKE ? OR postype ILIKE ? OR gerant ILIKE ? OR quartier ILIKE ? OR reference ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("User").
		Preload("PosEquipments").
		Preload("PosForms").
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
		"message":    "POS retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate by Commune id / UserUUID
func GetPaginatedPosByCommuneUUID(c *fiber.Ctx) error {
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

	var dataList []models.Pos
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.Pos{}).
		Where("pos.user_uuid = ?", UserUUID).
		Where("name ILIKE ? OR shop ILIKE ? OR postype ILIKE ? OR gerant ILIKE ? OR quartier ILIKE ? OR reference ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Where("pos.user_uuid = ?", UserUUID).
		Where("name ILIKE ? OR shop ILIKE ? OR postype ILIKE ? OR gerant ILIKE ? OR quartier ILIKE ? OR reference ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("User").
		Preload("PosForms").
		Preload("PosEquipments").
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
		"message":    "POS retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllPoss(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Pos
	db.Where("status = ?", true).Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Pos",
		"data":    data,
	})
}

// Get All data by manager
func GetAllPosByManager(c *fiber.Ctx) error {
	db := database.DB

	countryUUID := c.Params("country_uuid")

	var data []models.Pos
	db.Where("country_uuid = ?", countryUUID).
		Where("status = ?", true).Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Pos",
		"data":    data,
	})
}

// Get All data by ASM
func GetAllPosByASM(c *fiber.Ctx) error {
	db := database.DB

	ProvinceUUID := c.Params("province_uuid")

	var data []models.Pos
	db.Where("province_uuid = ?", ProvinceUUID).
		Where("status = ?", true).Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Pos",
		"data":    data,
	})
}

// Get All data by Supervisor
func GetAllPosBySup(c *fiber.Ctx) error {
	db := database.DB

	AreaUUID := c.Params("area_uuid")

	var data []models.Pos
	db.Where("area_uuid = ?", AreaUUID).
		Where("status = ?", true).Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Pos",
		"data":    data,
	})
}

// Get All data by DR
func GetAllPosByDR(c *fiber.Ctx) error {
	db := database.DB

	SubAreaUUID := c.Params("sub_area_uuid")

	var data []models.Pos
	db.Where("sub_area_uuid = ?", SubAreaUUID).
		Where("status = ?", true).Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Pos",
		"data":    data,
	})
}

// Get All data by CYclo
func GetAllPosByCyclo(c *fiber.Ctx) error {
	db := database.DB

	UserUUID := c.Params("user_uuid")

	var data []models.Pos
	db.Where("user_uuid = ?", UserUUID).
		Where("status = ?", true).Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Pos",
		"data":    data,
	})
}

// Get one data
func GetPos(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB
	var pos models.Pos
	db.Where("uuid = ?", uuid).First(&pos)
	if pos.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No pos name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "pos found",
			"data":    pos,
		},
	)
}

// Create data
func CreatePos(c *fiber.Ctx) error {
	p := &models.Pos{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	p.UUID = uuid.New().String()
	p.Sync = true
	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "pos created success",
			"data":    p,
		},
	)
}

// Update data
func UpdatePos(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	type UpdateData struct {
		Name      string `gorm:"not null" json:"name"` // Celui qui vend
		Shop      string `json:"shop"`                 // Nom du shop
		Postype   string `json:"postype"`              // Type de POS
		Gerant    string `json:"gerant"`               // name of the onwer of the pos
		Avenue    string `json:"avenue"`
		Quartier  string `json:"quartier"`
		Reference string `json:"reference"`
		Telephone string `json:"telephone"`
		Image     string `json:"image"`

		CountryUUID  string `json:"country_uuid"`
		ProvinceUUID string `json:"province_uuid"`
		AreaUUID     string `json:"area_uuid"`
		SubAreaUUID  string `json:"sub_area_uuid"`

		ManagerUUID string `json:"manager_uuid"`
		Manager     string `json:"manager" gorm:"default:''"`
		SupportUUID string `json:"support_uuid"`
		Support     string `json:"support" gorm:"default:''"`
		AsmUUID     string `json:"asm_uuid"`
		Asm         string `json:"asm" gorm:"default:''"`
		SupUUID     string `json:"sup_uuid"`
		Sup         string `json:"sup" gorm:"default:''"`
		DrUUID      string `json:"dr_uuid"`
		Dr          string `json:"dr" gorm:"default:''"`
		CycloUUID   string `json:"cyclo_uuid"`
		Cyclo       string `json:"cyclo" gorm:"default:''"`

		UserUUID string `json:"user_uuid"`

		Status    bool   `json:"status"`
		Signature string `json:"signature"`
	}

	var updateData UpdateData

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Review your input",
				"data":    nil,
			},
		)
	}

	pos := new(models.Pos)

	db.Where("uuid = ?", uuid).First(&pos)
	pos.Name = updateData.Name
	pos.Shop = updateData.Shop
	pos.Postype = updateData.Postype
	pos.Gerant = updateData.Gerant
	pos.Avenue = updateData.Avenue
	pos.Quartier = updateData.Quartier
	pos.Reference = updateData.Reference
	pos.Telephone = updateData.Telephone
	pos.CountryUUID = updateData.CountryUUID
	pos.ProvinceUUID = updateData.ProvinceUUID
	pos.AreaUUID = updateData.AreaUUID
	pos.SubAreaUUID = updateData.SubAreaUUID
	pos.AsmUUID = updateData.AsmUUID
	pos.Asm = updateData.Asm
	pos.SupUUID = updateData.SupUUID
	pos.Sup = updateData.Sup
	pos.DrUUID = updateData.DrUUID
	pos.Dr = updateData.Dr
	pos.CycloUUID = updateData.CycloUUID
	pos.Cyclo = updateData.Cyclo
	pos.UserUUID = updateData.UserUUID
	pos.Status = updateData.Status
	pos.Signature = updateData.Signature
	pos.Sync = true

	db.Save(&pos)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "POS updated success",
			"data":    pos,
		},
	)
}

// Delete data
func DeletePos(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	db := database.DB

	var pos models.Pos
	db.Where("uuid = ?", uuid).First(&pos)
	if pos.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No POS name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&pos)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "POS deleted success",
			"data":    nil,
		},
	)
}
