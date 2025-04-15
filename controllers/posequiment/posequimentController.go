package posequiment

import (
	"fmt"
	"strconv"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Paginate
func GetPaginatedPosEquipment(c *fiber.Ctx) error {
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

	var dataList []models.PosEquipment
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.PosEquipment{}).
		Joins("JOIN pos ON pos.id = pos_equipment.pos_id").
		Where("pos.name ILIKE ? OR pos_equipment.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Joins("JOIN pos ON pos.id = pos_equipment.pos_id").
		Where("pos.name ILIKE ? OR pos_equipment.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("pos_equipment.updated_at DESC").
		Preload("Pos").
		Preload("Parasol").
		Preload("Stand").
		Preload("Kiosk").
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

	fmt.Printf("Total Records: %d,Total Page: %d, Total Pages: %d\n", totalRecords, page, totalPages)

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

// Paginate
func GetPaginatedPosEquipmentByPos(c *fiber.Ctx) error {
	db := database.DB

	PosUUID := c.Params("pos_id")

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

	var dataList []models.PosEquipment
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.PosEquipment{}).
		Joins("JOIN pos ON pos.id = pos_equipment.pos_id").
		Where("pos.name ILIKE ? OR pos_equipment.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Where("pos.id = ?", PosUUID).
		Count(&totalRecords)

	// Fetch paginated data
	err = db.
		Joins("JOIN pos ON pos.id = pos_equipment.pos_id").
		Where("pos.name ILIKE ? OR pos_equipment.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Where("pos.id = ?", PosUUID).
		Offset(offset).
		Limit(limit).
		Order("pos_equipment.updated_at DESC").
		Preload("Pos").
		Preload("Parasol").
		Preload("Stand").
		Preload("Kiosk").
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

	fmt.Printf("Total Records: %d,Total Page: %d, Total Pages: %d\n", totalRecords, page, totalPages)

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

// Get all data by search
func GetAllPosEquipmentSearch(c *fiber.Ctx) error {
	db := database.DB
	search := c.Query("search", "")

	var data []models.PosEquipment
	if search != "" {
		db.Where("name ILIKE ? OR shop ILIKE ?", "%"+search+"%", "%"+search+"%").Find(&data)
	} else {
		db.Find(&data)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Pos",
		"data":    data,
	})
}

// Get All data
func GetAllPosEquipments(c *fiber.Ctx) error {
	db := database.DB
	var data []models.PosEquipment
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Pos",
		"data":    data,
	})
}

// query data
func GetPosEquipmentByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var poss []models.PosEquipment
	db.Where("pos_id = ?", id).Find(&poss)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "poss by id found",
		"data":    poss,
	})
}

// Get one data
func GetPosEquipment(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB
	var posequipment models.PosEquipment
	db.Where("uuid = ?", uuid).First(&posequipment)
	if posequipment.PosUUID == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No PosEquipment  name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Equipment  found",
			"data":    posequipment,
		},
	)
}

// Create data
func CreatePosEquipment(c *fiber.Ctx) error {
	p := &models.PosEquipment{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	p.UUID = uuid.New().String()
	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "poPosEquipments created success",
			"data":    p,
		},
	)
}

// Update data
func UpdatePosEquipment(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	type UpdateData struct {
		UUID          string `json:"uuid"`
		PosUUID       string   `json:"pos_uuid" gorm:"type:varchar(255);not null"`
		Parasol       string `json:"parasol"`                        //Equrtor, Compatition, Parasol
		ParasolStatus string `gorm:"not null" json:"parasol_status"` // Status d'equipements  Casser, Vieux, Bien

		Stand       string `json:"stand"`                        //Equrtor, Compatition, Parasol
		StandStatus bool   `gorm:"not null" json:"stand_status"` // Status d'equipements  Casser, Vieux, Bien

		Kiosk       string `json:"kiosk"`                        //Equrtor, Compatition, Parasol
		KioskStatus bool   `gorm:"not null" json:"kiosk_Status"` // Status d'equipements  Casser, Vieux, Bien

		Signature string `json:"signature"`
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

	posequiment := new(models.PosEquipment)

	db.Where("uuid = ?", uuid).First(&posequiment)
	posequiment.PosUUID = updateData.PosUUID
	posequiment.Parasol = updateData.Parasol
	posequiment.ParasolStatus = updateData.ParasolStatus
	posequiment.Stand = updateData.Stand
	posequiment.StandStatus = updateData.StandStatus
	posequiment.Kiosk = updateData.Kiosk
	posequiment.KioskStatus = updateData.KioskStatus
	posequiment.Signature = updateData.Signature

	db.Save(&posequiment)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "POSEQIPMENT updated success",
			"data":    posequiment,
		},
	)

}

// Delete data
func DeletePosEquipment(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	db := database.DB

	var PosEquipment models.PosEquipment
	db.Where("uuid = ?", uuid).First(&PosEquipment)
	if PosEquipment.PosUUID == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No POSEQIPMENT name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&PosEquipment)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "POSEQIPMENT deleted success",
			"data":    nil,
		},
	)
}
