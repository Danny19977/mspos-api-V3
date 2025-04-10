package user

import (
	"strconv"
	"strings"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/danny19977/mspos-api-v3/utils"
	"github.com/gofiber/fiber/v2"
)

// Paginate
func GetPaginatedUsers(c *fiber.Ctx) error {
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

	var users []models.User
	var totalRecords int64

	// Count total records matching the search query
	db.Model(&models.User{}).
		Where("fullname ILIKE ? OR title ILIKE ?", "%"+search+"%", "%"+search+"%").
		Count(&totalRecords)

	err = db.
		Where("fullname ILIKE ? OR title ILIKE ?", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("users.updated_at DESC").
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("Cyclo").
		Preload("Dr").
		Preload("Sup").
		Preload("Asm").
		Preload("Managers").
		Find(&users).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch Users",
			"error":   err.Error(),
		})
	}

	// Calculate total pages
	totalPages := int((totalRecords + int64(limit) - 1) / int64(limit))

	//  Prepare pagination metadata
	pagination := map[string]interface{}{
		"total_records": totalRecords,
		"total_pages":   totalPages,
		"current_page":  page,
		"page_size":     limit,
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "Users retrieved successfully",
		"data":       users,
		"pagination": pagination,
	})
}

// query all data
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User
	db.Find(&users)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All users",
		"data":    users,
	})
}

// query data
func GetUserByID(c *fiber.Ctx) error {
	ProvinceUUID := c.Params("id")
	db := database.DB
	var users []models.User
	db.Where("province_uuid = ?", ProvinceUUID).Where("role = ?", "DR").Where("status = ?", true).Find(&users)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "users by id found",
		"data":    users,
	})
}

// Get one data
func GetUser(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB
	var user models.User
	db.Where("uuid = ?", uuid).First(&user)
	if user.Fullname == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No User name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "User found",
			"data":    user,
		},
	)
}

// Create data
func CreateUser(c *fiber.Ctx) error {
	p := &models.User{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	if p.Fullname == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Form not complete",
				"data":    nil,
			},
		)
	}

	if p.Password != p.PasswordConfirm {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	user := &models.User{
		// UUID:       p.UUID,
		Fullname:   p.Fullname,
		Email:      p.Email,
		Title:      p.Title,
		Phone:      p.Phone,
		Role:       p.Role,
		Permission: p.Permission,
		Image:      p.Image,
		Status:     p.Status,
		// IsManager:  p.IsManager,
		Signature: p.Signature,
	}

	user.SetPassword(p.Password)

	if err := utils.ValidateStruct(*user); err != nil {
		c.Status(400)
		return c.JSON(err)
	}

	user.UUID = utils.GenerateUUID()

	if err := database.DB.Create(user).Error; err != nil {
		c.Status(500)
		sm := strings.Split(err.Error(), ":")
		m := strings.TrimSpace(sm[1])

		return c.JSON(fiber.Map{
			"message": m,
		})
	}

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "User Created success",
			"data":    user,
		},
	)
}

// Update data
func UpdateUser(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB

	type UpdateDataInput struct {
		UUID            string `json:"uuid"`
		Fullname        string `json:"fullname"`
		Email           string `json:"email"`
		Title           string `json:"title"`
		Phone           string `json:"phone"`
		Password        string `json:"password" validate:"required"`
		PasswordConfirm string `json:"password_confirm"`
		Role            string `json:"role"`
		Permission      string `json:"permission"`
		Image           string `json:"image"`
		Status          bool   `json:"status"`
		CountryUUID     string `json:"country_uuid" gorm:"type:varchar(255);not null"`
		ProvinceUUID    string `json:"province_uuid" gorm:"type:varchar(255);not null"`
		AreaUUID        string `json:"area_uuid" gorm:"type:varchar(255);not null"`
		SubAreaUUID     string `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
		CommuneUUID     string `json:"commune_uuid" gorm:"type:varchar(255);not null"`
		// IsManager  bool   `json:"is_manager"`
		Signature string `json:"signature"`
	}
	var updateData UpdateDataInput

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Review your input",
				"data":    nil,
			},
		)
	}

	user := new(models.User)

	db.Where("uuid = ?", uuid).First(&user)
	user.Fullname = updateData.Fullname
	user.Email = updateData.Email
	user.Title = updateData.Title
	user.Phone = updateData.Phone
	user.Role = updateData.Role
	user.Permission = updateData.Permission
	user.Image = updateData.Image
	user.Status = updateData.Status
	user.CountryUUID = updateData.CountryUUID
	user.ProvinceUUID = updateData.ProvinceUUID
	user.AreaUUID = updateData.AreaUUID
	user.SubAreaUUID = updateData.SubAreaUUID
	user.CommuneUUID = updateData.CommuneUUID
	user.Signature = updateData.Signature

	db.Save(&user)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "User updated success",
			"data":    user,
		},
	)
}

// Delete data
func DeleteUser(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	db := database.DB

	var User models.User
	db.Where("uuid = ?", uuid).First(&User)
	if User.Fullname == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No User name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&User)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "User deleted success",
			"data":    nil,
		},
	)
}
