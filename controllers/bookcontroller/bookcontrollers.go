package bookcontroller

import (
	"go-restapi/models"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(context *fiber.Ctx) error {
	var book []models.Book
	models.DB.Find(&book)

	return context.Status(fiber.StatusOK).JSON(book)
}

func Show(context *fiber.Ctx) error {
	id := context.Params("id")
	var book models.Book
	// validasi data tersedia atau tidak
	if err := models.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return context.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data Tidak Ditemukan",
			})
		}

		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data Tidak Ditemukan",
		})
	}

	return context.Status(fiber.StatusOK).JSON(book)
}

func Create(context *fiber.Ctx) error {
	var book models.Book
	if err := context.BodyParser(&book); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&book).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data Berhasil Dibuat",
	})
}

func Update(context *fiber.Ctx) error {
	id := context.Params("id")
	var book models.Book
	if err := context.BodyParser(&book); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tidak Dapat Mengupdate Data",
		})
	}

	return context.JSON(fiber.Map{
		"message": "Berhasil Update Data",
	})
}

func Delete(context *fiber.Ctx) error {
	id := context.Params("id")
	var book models.Book
	if models.DB.Delete(&book, id).RowsAffected == 0 {
		return context.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak Dapat Menghapus Data",
		})
	}
	models.DB.Exec("ALTER TABLE books DROP id")
	models.DB.Exec("ALTER TABLE books ADD id INT NOT NULL AUTO_INCREMENT PRIMARY KEY FIRST")
	models.DB.Exec("ALTER TABLE books AUTO_INCREMENT = 1")

	return context.JSON(fiber.Map{
		"message": "Data Berhasil Dihapus",
	})
}
