package posform

import (
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Paginate ALL data
func GetPaginatedPosForm(c *fiber.Ctx) error {
	db := database.DB

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

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
		Joins("JOIN countries ON pos_forms.country_uuid=countries.uuid").
		Joins("JOIN provinces ON pos_forms.province_uuid=provinces.uuid").
		Joins("JOIN sups ON pos_forms.sup_uuid=sups.uuid").
		Joins("JOIN areas ON pos_forms.area_uuid=areas.uuid").
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("provinces.name ILIKE ? OR pos.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN countries ON pos_forms.country_uuid=countries.uuid").
		Joins("JOIN provinces ON pos_forms.province_uuid=provinces.uuid").
		Joins("JOIN sups ON pos_forms.sup_uuid=sups.uuid").
		Joins("JOIN areas ON pos_forms.area_uuid=areas.uuid").
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("provinces.name ILIKE ? OR pos.name ILIKE ?", "%"+search+"%", "%"+search+"%").
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
	pagination := map[string]any{
		"total_pages": totalPages,
		"page":        page,
		"page_size":   limit,
		"length":      dataList,
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "All PosForms Successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data UserUUID
func GetPosformByUserUUID(c *fiber.Ctx) error {
	db := database.DB

	UserUUID := c.Params("user_uuid")

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

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
		Joins("JOIN provinces ON pos_forms.province_uuid=provinces.uuid").
		Joins("JOIN sups ON pos_forms.sup_uuid=sups.uuid").
		Joins("JOIN areas ON pos_forms.area_uuid=areas.uuid").
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.user_uuid = ?", UserUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("provinces.name ILIKE ? OR pos.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN provinces ON pos_forms.province_uuid=provinces.uuid").
		Joins("JOIN sups ON pos_forms.sup_uuid=sups.uuid").
		Joins("JOIN areas ON pos_forms.area_uuid=areas.uuid").
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.user_uuid = ?", UserUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("provinces.name ILIKE ? OR pos.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("pos_forms.updated_at DESC").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("ASM").
		Preload("Sup").
		Preload("Dr").
		Preload("Cyclo").
		Preload("Pos").
		Preload("PosFormItems.Brand").
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
	pagination := map[string]any{
		"total_pages": totalPages,
		"page":        page,
		"page_size":   limit,
		"length":      dataList,
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "All PosForms Successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data province
func GetPosformByProvinceUUID(c *fiber.Ctx) error {
	db := database.DB

	ProvinceUUID := c.Params("province_uuid")

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

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
		Joins("JOIN provinces ON pos_forms.province_uuid=provinces.uuid").
		Joins("JOIN sups ON pos_forms.sup_uuid=sups.uuid").
		Joins("JOIN areas ON pos_forms.area_uuid=areas.uuid").
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("pos_forms.province_uuid = ?", ProvinceUUID).
		Where("provinces.name ILIKE ? OR pos.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN provinces ON pos_forms.province_uuid=provinces.uuid").
		Joins("JOIN sups ON pos_forms.sup_uuid=sups.uuid").
		Joins("JOIN areas ON pos_forms.area_uuid=areas.uuid").
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.province_uuid = ?", ProvinceUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("provinces.name ILIKE ? OR pos.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("pos_forms.updated_at DESC").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("ASM").
		Preload("Sup").
		Preload("Dr").
		Preload("Cyclo").
		Preload("Pos").
		Preload("PosFormItems.Brand").
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
	pagination := map[string]any{
		"total_pages": totalPages,
		"page":        page,
		"page_size":   limit,
		"length":      dataList,
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "All PosForms Successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data area
func GetPosformByAreaUUID(c *fiber.Ctx) error {
	db := database.DB

	AreaUUID := c.Params("area_uuid")

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

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
		Joins("JOIN provinces ON pos_forms.province_uuid=provinces.uuid").
		Joins("JOIN sups ON pos_forms.sup_uuid=sups.uuid").
		Joins("JOIN users ON pos_forms.user_uuid=users.uuid").
		Joins("JOIN areas ON pos_forms.area_uuid=areas.uuid").
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.area_uuid = ?", AreaUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("users.fullname ILIKE ? OR countries.name ILIKE ? OR provinces.name ILIKE ? OR sups.name ILIKE ? OR users.name ILIKE ? OR pos.name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN provinces ON pos_forms.province_uuid=provinces.uuid").
		Joins("JOIN sups ON pos_forms.sup_uuid=sups.uuid").
		Joins("JOIN users ON pos_forms.user_uuid=users.uuid").
		Joins("JOIN areas ON pos_forms.area_uuid=areas.uuid").
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.area_uuid = ?", AreaUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("users.fullname ILIKE ? OR countries.name ILIKE ? OR provinces.name ILIKE ? OR sups.name ILIKE ? OR users.name ILIKE ? OR pos.name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
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
		Preload("PosFormItems.Brand").
		Preload("User").
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
		"total_pages": totalPages,
		"page":        page,
		"page_size":   limit,
		"length":      dataList,
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "All PosForms Successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data Subarea ID
func GetPaginatedPosFormBySubAreaUUID(c *fiber.Ctx) error {
	db := database.DB

	subAreaUUID := c.Params("subarea_uuid")

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

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
		Joins("JOIN provinces ON pos_forms.province_uuid=provinces.uuid").
		Joins("JOIN sups ON pos_forms.sup_uuid=sups.uuid").
		Joins("JOIN users ON pos_forms.user_uuid=users.uuid").
		Joins("JOIN areas ON pos_forms.area_uuid=areas.uuid").
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.subarea_uuid = ?", subAreaUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("users.fullname ILIKE ? OR countries.name ILIKE ? OR provinces.name ILIKE ? OR sups.name ILIKE ? OR users.name ILIKE ? OR pos.name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN provinces ON pos_forms.province_uuid=provinces.uuid").
		Joins("JOIN sups ON pos_forms.sup_uuid=sups.uuid").
		Joins("JOIN users ON pos_forms.user_uuid=users.uuid").
		Joins("JOIN areas ON pos_forms.area_uuid=areas.uuid").
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.subarea_uuid = ?", subAreaUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("users.fullname ILIKE ? OR countries.name ILIKE ? OR provinces.name ILIKE ? OR sups.name ILIKE ? OR users.name ILIKE ? OR pos.name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
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
		Preload("PosFormItems.Brand").
		Preload("User").
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
		"total_pages": totalPages,
		"page":        page,
		"page_size":   limit,
		"length":      dataList,
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "All PosForms Successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data Commune ID
func GetPaginatedPosFormByCommuneUUID(c *fiber.Ctx) error {
	db := database.DB

	CommuneUUID := c.Params("commune_uuid")

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

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
		Joins("JOIN provinces ON pos_forms.province_uuid=provinces.uuid").
		Joins("JOIN sups ON pos_forms.sup_uuid=sups.uuid").
		Joins("JOIN users ON pos_forms.user_uuid=users.uuid").
		Joins("JOIN areas ON pos_forms.area_uuid=areas.uuid").
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.commune_uuid = ?", CommuneUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("users.fullname ILIKE ? OR countries.name ILIKE ? OR provinces.name ILIKE ? OR sups.name ILIKE ? OR users.name ILIKE ? OR pos.name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Joins("JOIN provinces ON pos_forms.province_uuid=provinces.uuid").
		Joins("JOIN sups ON pos_forms.sup_uuid=sups.uuid").
		Joins("JOIN users ON pos_forms.user_uuid=users.uuid").
		Joins("JOIN areas ON pos_forms.area_uuid=areas.uuid").
		Joins("JOIN pos ON pos_forms.pos_uuid=pos.uuid").
		Where("pos_forms.commune_uuid = ?", CommuneUUID).
		Where("pos_forms.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("users.fullname ILIKE ? OR countries.name ILIKE ? OR provinces.name ILIKE ? OR sups.name ILIKE ? OR users.name ILIKE ? OR pos.name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
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
		Preload("PosFormItems.Brand").
		Preload("User").
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
		"total_pages": totalPages,
		"page":        page,
		"page_size":   limit,
		"length":      dataList,
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "All PosForms Successfully",
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

	p.UUID = uuid.New().String()
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
		UUID string `json:"uustringid"`

		Price   float64 `gorm:"default: 0" json:"price"`
		Comment string  `json:"comment"`

		Latitude  string `json:"latitude"`  // Latitude of the user
		Longitude string `json:"longitude"` // Longitude of the user
		Signature string `json:"signature"`

		PosUUID      string `json:"pos_uuid" gorm:"type:varchar(255);not null"`
		ProvinceUUID string `json:"province_uuid" gorm:"type:varchar(255);not null"`
		AreaUUID     string `json:"area_uuid" gorm:"type:varchar(255);not null"`
		SubAreaUUID  string `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
		AsmUUID      string `json:"asm_uuid" gorm:"type:varchar(255);not null"`
		SupUUID      string `json:"sup_uuid" gorm:"type:varchar(255);not null"`
		DrUUID       string `json:"dr_uuid" gorm:"type:varchar(255);not null"`
		CycloUUID    string `json:"cyclo_uuid" gorm:"type:varchar(255);not null"`

		// Eq        int64  `json:"eq"`
		// Eq1       int64  `json:"eq1"`
		// Sold      int64  `json:"sold"`
		// Dhl       int64  `json:"dhl"`
		// Dhl1      int64  `json:"dhl1"`
		// Ar        int64  `json:"ar"`
		// Ar1       int64  `json:"ar1"`
		// Sbl       int64  `json:"sbl"`
		// Sbl1      int64  `json:"sbl1"`
		// Pmf       int64  `json:"pmf"`
		// Pmf1      int64  `json:"pmf1"`
		// Pmm       int64  `json:"pmm"`
		// Pmm1      int64  `json:"pmm1"`
		// Ticket    int64  `json:"ticket"`
		// Ticket1   int64  `json:"ticket1"`
		// Mtc       int64  `json:"mtc"`
		// Mtc1      int64  `json:"mtc1"`
		// Ws        int64  `json:"ws"`
		// Ws1       int64  `json:"ws1"`
		// Mast      int64  `json:"mast"`
		// Mast1     int64  `json:"mast1"`
		// Oris      int64  `json:"oris"`
		// Oris1     int64  `json:"oris1"`
		// Elite     int64  `json:"elite"`
		// Elite1    int64  `json:"elite1"`
		// Yes       int64  `json:"yes"`
		// Yes1      int64  `json:"yes1"`
		// Time      int64  `json:"time"`
		// Time1     int64  `json:"time1"`

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
	posform.AsmUUID = updateData.AsmUUID
	posform.SupUUID = updateData.SupUUID
	posform.DrUUID = updateData.DrUUID
	posform.CycloUUID = updateData.CycloUUID

	// posform.Eq = updateData.Eq
	// posform.Eq1 = updateData.Eq1
	// posform.Sold = updateData.Sold
	// posform.Dhl = updateData.Dhl
	// posform.Dhl1 = updateData.Dhl1
	// posform.Ar = updateData.Ar
	// posform.Ar1 = updateData.Ar1
	// posform.Sbl = updateData.Sbl
	// posform.Sbl1 = updateData.Sbl1
	// posform.Pmf = updateData.Pmf
	// posform.Pmf1 = updateData.Pmf1
	// posform.Pmm = updateData.Pmm
	// posform.Pmm1 = updateData.Pmm1
	// posform.Ticket = updateData.Ticket
	// posform.Ticket1 = updateData.Ticket1
	// posform.Mtc = updateData.Mtc
	// posform.Mtc1 = updateData.Mtc1
	// posform.Ws = updateData.Ws
	// posform.Ws1 = updateData.Ws1
	// posform.Mast = updateData.Mast
	// posform.Mast1 = updateData.Mast1
	// posform.Oris = updateData.Oris
	// posform.Oris1 = updateData.Oris1
	// posform.Elite = updateData.Elite
	// posform.Elite1 = updateData.Elite1
	// posform.Yes = updateData.Yes
	// posform.Yes1 = updateData.Yes1
	// posform.Time = updateData.Time
	// posform.Time1 = updateData.Time1

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
