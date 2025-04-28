package routeplan

import (
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
	// "github.com/google/uuid"
)

// Paginate
func GetPaginatedRouteplan(c *fiber.Ctx) error {

	// Initialize database connection
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

	var dataList []models.RoutePlan
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.RoutePlan{}).
		Joins("JOIN users ON users.uuid = route_plans.user_uuid").
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Joins("JOIN users ON users.uuid = route_plans.user_uuid").
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("User").
		Preload("RoutePlanItems").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch provinces",
			"error":   err.Error(),
		})
	}

	/// Calculate total pages
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
		"message":    "Routeplan retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// GetPaginatedRouthplaByProvinceID
func GetPaginatedRouthplaByProvinceUUID(c *fiber.Ctx) error {

	provinceUUID := c.Params("province_uuid")

	// Initialize database connection
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

	var dataList []models.RoutePlan
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.RoutePlan{}).
		Joins("JOIN users ON users.uuid = route_plans.user_uuid").
		Where("route_plans.province_uuid = ?", provinceUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Joins("JOIN users ON users.uuid = route_plans.user_uuid").
		Where("route_plans.province_uuid = ?", provinceUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("User").
		Preload("RoutePlanItems").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch provinces",
			"error":   err.Error(),
		})
	}

	/// Calculate total pages
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
		"message":    "Routeplan retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// GetPaginatedRouthplaByareaUUID
func GetPaginatedRouthplaByareaUUID(c *fiber.Ctx) error {

	areaUUID := c.Params("area_uuid")

	// Initialize database connection
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

	var dataList []models.RoutePlan
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.RoutePlan{}).
		Joins("JOIN users ON users.uuid = route_plans.user_uuid").
		Where("route_plans.area_uuid = ?", areaUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Joins("JOIN users ON users.uuid = route_plans.user_uuid").
		Where("route_plans.area_uuid = ?", areaUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("User").
		Preload("RoutePlanItems").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch provinces",
			"error":   err.Error(),
		})
	}

	/// Calculate total pages
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
		"message":    "Routeplan retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// GetPaginatedRouthplaBysubareaUUID
func GetPaginatedRouthplaBySubareaUUID(c *fiber.Ctx) error {

	subareaUUID := c.Params("subarea_uuid")

	// Initialize database connection
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

	var dataList []models.RoutePlan
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.RoutePlan{}).
		Joins("JOIN users ON users.uuid = route_plans.user_uuid").
		Where("route_plans.sub_area_uuid = ?", subareaUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Joins("JOIN users ON users.uuid = route_plans.user_uuid").
		Where("route_plans.sub_area_uuid = ?", subareaUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("User").
		Preload("RoutePlanItems").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch provinces",
			"error":   err.Error(),
		})
	}

	/// Calculate total pages
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
		"message":    "Routeplan retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// GetPaginatedRouteplaBycommuneUUID
func GetPaginatedRouteplaBycommuneUUID(c *fiber.Ctx) error {

	communeUUID := c.Params("commune_uuid")

	// Initialize database connection
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

	var dataList []models.RoutePlan
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.RoutePlan{}).
		Joins("JOIN users ON users.uuid = route_plans.user_uuid").
		Where("route_plans.commune_uuid = ?", communeUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Joins("JOIN users ON users.uuid = route_plans.user_uuid").
		Where("route_plans.commune_uuid = ?", communeUUID).
		Where("users.fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("User").
		Preload("RoutePlanItems").
		Find(&dataList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch provinces",
			"error":   err.Error(),
		})
	}

	/// Calculate total pages
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
		"message":    "Routeplan retrieved successfully",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllRouteplan(c *fiber.Ctx) error {
	db := database.DB

	var data []models.RoutePlan
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Routeplan",
		"data":    data,
	})
}

// Get All data by id
func GetAllRouteplanBySearch(c *fiber.Ctx) error {
	db := database.DB

	search := c.Query("search", "")

	var data []models.RoutePlan
	db.Where("name ILIKE ?", "%"+search+"%").
		Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Routeplan",
		"data":    data,
	})
}

// Get one data
func GetRouteplan(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	var Routeplan models.RoutePlan
	db.Where("uuid = ?", uuid).First(&Routeplan)
	if Routeplan.UUID == "0" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Routeplan name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "RoutePlan found",
			"data":    Routeplan,
		},
	)
}

// Create data
func CreateRouteplan(c *fiber.Ctx) error {
	p := &models.RoutePlan{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	// p.UUID = uuid.New().String()
	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "routeplan created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateRouteplan(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	type UpdateData struct {
		UUID string `json:"uuid"`

		UserUUID     string `json:"user_uuid"`
		ProvinceUUID string `json:"province_uuid" gorm:"type:varchar(255);not null"`
		AreaUUID     string `json:"area_uuid" gorm:"type:varchar(255);not null"`
		SubAreaUUID  string `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
		CommuneUUID  string `json:"commune_uuid" gorm:"type:varchar(255);not null"`
		TotalPOS     int    `json:"total_pos"`
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

	RoutePlan := new(models.RoutePlan)

	db.Where("uuid = ?", uuid).First(&RoutePlan)
	RoutePlan.UserUUID = updateData.UserUUID
	RoutePlan.ProvinceUUID = updateData.ProvinceUUID
	RoutePlan.AreaUUID = updateData.AreaUUID
	RoutePlan.SubAreaUUID = updateData.SubAreaUUID
	RoutePlan.CommuneUUID = updateData.CommuneUUID
	// RoutePlan.TotalPOS = updateData.TotalPOS
	RoutePlan.Signature = updateData.Signature

	db.Save(&RoutePlan)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "RoutePlan updated success",
			"data":    RoutePlan,
		},
	)

}

// Delete data
func DeleteRouteplan(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	db := database.DB

	var routeplan models.RoutePlan
	db.Where("uuid = ?", uuid).First(&routeplan)
	if routeplan.UUID == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No routeplan name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&routeplan)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "RoutePlan deleted success",
			"data":    nil,
		},
	)
}
