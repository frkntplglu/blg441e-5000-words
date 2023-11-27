package handlers

import (
	"strconv"

	"github.com/frkntplglu/emir-backend/internal/models"
	"github.com/frkntplglu/emir-backend/internal/services"
	"github.com/frkntplglu/emir-backend/pkg/utils"
	. "github.com/go-swagno/swagno"
	"github.com/gofiber/fiber/v2"
)

type QuizHandler struct {
	quizService services.QuizService
}

func NewQuizHandler(quizService services.QuizService) *QuizHandler {
	return &QuizHandler{
		quizService: quizService,
	}
}

func (h *QuizHandler) handleGetAllQuizzes(ctx *fiber.Ctx) error {
	pagination := utils.GeneratePaginationFromCtx(ctx)

	quizzes, err := h.quizService.GetAllQuizzes(models.Quiz{}, &pagination)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Bir hata oluştu",
				Reason:  err.Error(),
			},
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(models.SuccessResponse{
		Success: true,
		Data:    quizzes,
	})
}

func (h *QuizHandler) handleGetQuizById(ctx *fiber.Ctx) error {
	quizId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir id giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	word, err := h.quizService.GetQuizById(quizId)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Bir hata oluştu",
				Reason:  err.Error(),
			},
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(models.SuccessResponse{
		Success: true,
		Data:    word,
	})
}

func (h *QuizHandler) SetRoutes(a *fiber.App) {
	quizGroup := a.Group("/quizzes")
	quizGroup.Get("/", h.handleGetAllQuizzes)

	quizGroup.Get("/:id<int>", h.handleGetQuizById)
}

var QuizSwaggerEndpoints = []Endpoint{
	EndPoint(GET, "/quizzes", "Quizler", Params(IntQuery("page", true, "Pagination için sayfa numarası"), IntQuery("limit", true, "Bir sayfada dönecek toplam veri sayısı")), nil, []models.Word{}, models.FailureResponse{}, "Tüm quizleri listeler", nil),
	EndPoint(GET, "/quizzes/{quizId}", "Quizler", Params(IntParam("wordId", true, "Bilgileri istenen kelimenin idsi")), nil, models.Word{}, models.FailureResponse{}, "Id'ye göre tek bir quiz bilgisini döner", nil),
}
