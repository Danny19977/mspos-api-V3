package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/danny19977/mspos-api-v3/database"
	"github.com/danny19977/mspos-api-v3/models"
	"github.com/danny19977/mspos-api-v3/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func Register(c *fiber.Ctx) error {

	nu := new(models.User)

	if err := c.BodyParser(&nu); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if nu.Password != nu.PasswordConfirm {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	u := &models.User{
		Fullname:     nu.Fullname,
		Email:        nu.Email,
		Title:        nu.Title,
		Phone:        nu.Phone,
		Role:         nu.Role,
		Permission:   nu.Permission,
		Image:        nu.Image,
		Status:       nu.Status,
		Signature:    nu.Signature,
		CountryUUID:  nu.CountryUUID,
		ProvinceUUID: nu.ProvinceUUID,
		AreaUUID:     nu.AreaUUID,
		SubAreaUUID:  nu.SubAreaUUID,
		CommuneUUID:  nu.CommuneUUID,
	}

	u.SetPassword(nu.Password)

	if err := utils.ValidateStruct(*u); err != nil {
		c.Status(400)
		return c.JSON(err)
	}

	u.UUID = uuid.New().String()

	database.DB.Create(u)

	// if err := database.DB.Create(u).Error; err != nil {
	// 	c.Status(500)
	// 	sm := strings.Split(err.Error(), ":")
	// 	m := strings.TrimSpace(sm[1])

	// 	return c.JSON(fiber.Map{
	// 		"message": m,
	// 	})
	// }

	return c.JSON(fiber.Map{
		"message": "user account created",
		"data":    u,
	})
}

func Login(c *fiber.Ctx) error {

	lu := new(models.Login)

	if err := c.BodyParser(&lu); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := utils.ValidateStruct(*lu); err != nil {
		c.Status(400)
		return c.JSON(err)
	}

	u := &models.User{}

	database.DB.Where("email = ? OR phone = ?", lu.Identifier, lu.Identifier).First(&u)

	if u.UUID == "00000000-0000-0000-0000-000000000000" {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "invalid email or phone 😰",
		})
	}

	if err := u.ComparePassword(lu.Password); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "mot de passe incorrect! 😰",
		})
	}

	if !u.Status {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "vous n'êtes pas autorisé de se connecter 😰",
		})
	}

	token, err := utils.GenerateJwt(u.UUID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    token,
	})

}

func AuthUser(c *fiber.Ctx) error {

	token := c.Query("token")

	fmt.Println("token", token)

	// cookie := c.Cookies("token")
	UserUUID, _ := utils.VerifyJwt(token)

	fmt.Println("UserUUID", UserUUID)

	u := models.User{}

	database.DB.
		Joins("LEFT JOIN countries ON countries.uuid = users.country_uuid").
		Joins("LEFT JOIN provinces ON provinces.uuid = users.province_uuid").
		Joins("LEFT JOIN areas ON areas.uuid = users.area_uuid").
		Joins("LEFT JOIN sub_areas ON sub_areas.uuid = users.sub_area_uuid").
		Joins("LEFT JOIN communes ON communes.uuid = users.commune_uuid").
		Joins("LEFT JOIN asms ON asms.user_uuid = users.uuid").
		Joins("LEFT JOIN sups ON sups.user_uuid = users.uuid").
		Joins("LEFT JOIN drs ON drs.user_uuid = users.uuid").
		Joins("LEFT JOIN cyclos ON cyclos.user_uuid = users.uuid").
		Where("users.uuid = ?", UserUUID).
		Select(`
			users.*, 
			countries.name as country_name, 
			provinces.name as province_name, 
			areas.name as area_name, 
			sub_areas.name as subarea_name, 
			communes.name as commune_name,
			asms.uuid as asm_uuid,
			sups.uuid as sup_uuid,
			drs.uuid as dr_uuid,
			cyclos.uuid as cyclo_uuid
		`).
		Preload("Country").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
		Preload("Commune").
		Preload("Asm").
		Preload("Sup").
		Preload("Dr").
		Preload("Cyclo").
		First(&u)

	r := &models.UserResponse{
		ID:           u.ID,
		UUID:         u.UUID,
		Fullname:     u.Fullname,
		Email:        u.Email,
		Title:        u.Title,
		Phone:        u.Phone,
		Role:         u.Role,
		Permission:   u.Permission,
		Status:       u.Status,
		CountryUUID:  u.CountryUUID,
		Country:      u.Country,
		ProvinceUUID: u.ProvinceUUID,
		Province:     u.Province,
		AreaUUID:     u.AreaUUID,
		Area:         u.Area,
		SubAreaUUID:  u.SubAreaUUID,
		SubArea:      u.SubArea,
		CommuneUUID:  u.CommuneUUID,
		Commune:      u.Commune,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		Asm:          u.Asm,
		Sup:          u.Sup,
		Dr:           u.Dr,
		Cyclo:        u.Cyclo,
	}

	// json, err := json.MarshalIndent(r, "", "  ")
	// if err != nil {
	// 	fmt.Println("error", err)
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"status":  "error",
	// 		"message": "error",
	// 		"errors":  err.Error(),
	// 	})
	// }

	// fmt.Println("user", string(json))

	return c.JSON(r)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // 1 day ,
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
		"Logout":  "success",
	})

}

// User bioprofile
func UpdateInfo(c *fiber.Ctx) error {
	type UpdateDataInput struct {
		Fullname  string `json:"fullname"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Signature string `json:"signature"`
	}
	var updateData UpdateDataInput

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"errors":  err.Error(),
		})
	}

	cookie := c.Cookies("token")

	Id, _ := utils.VerifyJwt(cookie)

	UserUUID, _ := strconv.Atoi(Id)

	user := new(models.User)

	db := database.DB

	db.First(&user, UserUUID)
	user.Fullname = updateData.Fullname
	user.Email = updateData.Email
	user.Phone = updateData.Phone
	user.Signature = updateData.Signature

	db.Save(&user)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully updated",
		"data":    user,
	})

}

func ChangePassword(c *fiber.Ctx) error {
	type UpdateDataInput struct {
		OldPassword     string `json:"old_password"`
		Password        string `json:"password"`
		PasswordConfirm string `json:"password_confirm"`
	}
	var updateData UpdateDataInput

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"errors":  err.Error(),
		})
	}

	cookie := c.Cookies("token")

	UserUUID, _ := utils.VerifyJwt(cookie)

	user := new(models.User)

	database.DB.Where("id = ?", UserUUID).First(&user)

	if err := user.ComparePassword(updateData.OldPassword); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "votre mot de passe n'est pas correct! 😰",
		})
	}

	if updateData.Password != updateData.PasswordConfirm {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	p, err := utils.HashPassword(updateData.Password)
	if err != nil {
		return err
	}

	db := database.DB

	db.First(&user, user.ID)
	user.Password = p

	db.Save(&user)

	// successful update remove cookies
	rmCookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //1 day ,
		HTTPOnly: true,
	}
	c.Cookie(&rmCookie)

	return c.JSON(user)

}
