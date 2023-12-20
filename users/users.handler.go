package users

import (
	"fmt"
	"gofiber/config"
	"gofiber/helper"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	if config.DB == nil {
		return helper.SendErrorResponse(c, http.StatusInternalServerError, "Database connection not initialized")
	}

	kodeDaerah := c.Params("kode_daerah")
	fmt.Println("Kode Daerah:", kodeDaerah) // Log nilai kode_daerah

	var user []Users
	if err := config.DB.Where("kode_daerah = ?", kodeDaerah).Find(&user).Error; err != nil {
		return helper.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return helper.SendSuccessResponse(c, http.StatusOK, "Success", kodeDaerah)
}
