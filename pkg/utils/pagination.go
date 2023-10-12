package utils

import (
	"github.com/frkntplglu/emir-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GeneratePaginationFromCtx(ctx *fiber.Ctx) models.Pagination {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	return models.Pagination{
		Page:  page,
		Limit: limit,
	}
}
