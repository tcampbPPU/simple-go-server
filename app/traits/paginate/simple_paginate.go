package paginate

import (
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tcampbppu/server/app/traits/harmony"
	"gorm.io/gorm"
)

func Paginate(c *fiber.Ctx, db *gorm.DB, model interface{}) error {

	data := reflect.New(reflect.TypeOf(model)).Interface()

	perPage, _ := strconv.Atoi(c.Query("per_page"))
	page, _ := strconv.Atoi(c.Query("page"))

	if perPage == 0 {
		perPage = 10
	}

	if page == 0 {
		page = 1
	}

	offset := (page - 1) * perPage

	db.Limit(perPage).Offset(offset).Find(&data)

	response := harmony.ResponseDTO{
		Type:    "success",
		Success: true,
		Data:    data,
	}

	return harmony.Harmony(c, fiber.StatusOK, response)
}
