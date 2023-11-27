package handlers

import (
	"strconv"

	"github.com/frkntplglu/emir-backend/internal/models"
	"github.com/frkntplglu/emir-backend/internal/services"
	"github.com/frkntplglu/emir-backend/pkg/utils"
	. "github.com/go-swagno/swagno"
	"github.com/gofiber/fiber/v2"
)

type QuestionHandler struct {
	questionService services.QuestionService
}

func NewQuestionHandler(questionService services.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		questionService: questionService,
	}
}

func (h *QuestionHandler) handleGetAllQuestions(ctx *fiber.Ctx) error {
	pagination := utils.GeneratePaginationFromCtx(ctx)
	quizId := ctx.QueryInt("quiz_id")

	var questions []models.Question
	var err error

	if quizId == 0 {
		questions, err = h.questionService.GetAllQuestions(models.Question{}, &pagination)
	} else {
		questions, err = h.questionService.GetQuestionsByQuizId(quizId, &pagination)
	}

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
		Data:    questions,
	})
}

func (h *QuestionHandler) handleGetQuestionById(ctx *fiber.Ctx) error {
	questionId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir id giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	question, err := h.questionService.GetQuestionById(questionId)

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
		Data:    question,
	})
}

func (h *QuestionHandler) handleAnswerQuestion(ctx *fiber.Ctx) error {
	questionId, err := strconv.Atoi(ctx.Params("id"))

	var reqBody models.Answer
	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir cevap giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	question, err := h.questionService.GetQuestionById(questionId)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Bir hata oluştu",
				Reason:  err.Error(),
			},
		})
	}

	var result bool = false
	if question.Answer == reqBody.Answer {
		result = true
	}
	return ctx.Status(fiber.StatusOK).JSON(models.SuccessResponse{
		Success: true,
		Data:    result,
	})
}

func (h *QuestionHandler) SetRoutes(a *fiber.App) {
	questionGroup := a.Group("/questions")
	questionGroup.Get("/", h.handleGetAllQuestions)

	questionGroup.Get("/:id<int>", h.handleGetQuestionById)
	questionGroup.Post("/:id<int>", h.handleAnswerQuestion)
}

var QuestionSwaggerEndpoints = []Endpoint{
	EndPoint(GET, "/questions", "Quizler", Params(IntQuery("page", true, "Pagination için sayfa numarası"), IntQuery("limit", true, "Bir sayfada dönecek toplam veri sayısı")), nil, []models.Word{}, models.FailureResponse{}, "Tüm quizleri listeler", nil),
	EndPoint(GET, "/question/{quizId}", "Quizler", Params(IntParam("wordId", true, "Bilgileri istenen kelimenin idsi")), nil, models.Word{}, models.FailureResponse{}, "Id'ye göre tek bir quiz bilgisini döner", nil),
}
