package handler

import (
	"apirestecommerce/entity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

type CategoryHandler struct {
	DB *gorm.DB
}

func (handler CategoryHandler) Index(ctx *fiber.Ctx) error {
	var categories []entity.Category
	handler.DB.Find(&categories)
	return ctx.JSON(entity.ResponseSuccess{
		Message    : "success",
		StatusCode : fiber.StatusOK,
		Data       : categories,
	})
}
func (handler CategoryHandler) Show(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.JSON(entity.ResponseError{
			Message: "Data bulunamadı.Tekil işlemler için id gerekli",
			StatusCode: 422,
		})
	}
	var category entity.Category
	category.ID = uint(id)
	handler.DB.First(&category)

	return ctx.JSON(entity.ResponseSuccess{
		Message    : "success",
		StatusCode : fiber.StatusOK,
		Data       : category,
	})
}

func (handler CategoryHandler) Store(ctx *fiber.Ctx) error {
	category := entity.Category{
		CategoryName:        "Masa",
		ParentID:            1,
		CategoryKeywords:    "Masa,table,furniture,sandalye",
		CategorySlug:        "masa",
		CategoryDescription: "Masa ürünlerimiz gayet iyidir",
		CategoryStatus:      true,
	}
	result := handler.DB.Create(&category)
	if result.RowsAffected > 0 {
		return ctx.JSON(entity.ResponseSuccess{
			Message      : "success",
			StatusCode   : fiber.StatusOK,
			Data         : category,
		})
	} else {
		return ctx.JSON(entity.ResponseError{
			Message    : "error",
			StatusCode : fiber.StatusNoContent,
		})
	}
}

func (handler CategoryHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.JSON(entity.ResponseError{
			Message: "Data Silinemedi.",
			StatusCode: 422,
		})
	}
	var category entity.Category
	category.ID = uint(id)
	handler.DB.Delete(&category)

	return ctx.JSON(entity.ResponseSuccess{
		Message    : "success",
		StatusCode : fiber.StatusOK,
		Data       : category.ID,
	})
}
