package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/frkntplglu/emir-backend/internal/config"
	"github.com/frkntplglu/emir-backend/internal/models"
	"github.com/frkntplglu/emir-backend/internal/services"
	"github.com/frkntplglu/emir-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type SavedWordHandler struct {
	savedWordService services.SavedWordService
	wordService      services.WordService
}

func NewSavedWordHandler(savedWordService services.SavedWordService, wordService services.WordService) *SavedWordHandler {
	return &SavedWordHandler{
		savedWordService: savedWordService,
		wordService:      wordService,
	}
}

func (h *SavedWordHandler) handleGetAllSavedWords(ctx *fiber.Ctx) error {
	pagination := utils.GeneratePaginationFromCtx(ctx)

	words, err := h.savedWordService.GetAllWords(models.SavedWord{}, &pagination)

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
		Data:    words,
	})
}

func (h *SavedWordHandler) handleCreateSavedWord(ctx *fiber.Ctx) error {
	var reqBody struct {
		Id     int `json:"int"`
		WordId int `json:"wordId"`
	}

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir body giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	cookie := ctx.Cookies("accessToken")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ACCESS_TOKEN), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Yetkisiz işlem hatası",
				Reason:  err.Error(),
			},
		})
	}

	_, err = h.wordService.GetWordById(reqBody.WordId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Girmiş olduğunuz kelime sistemimizde bulunmamaktadır.",
				Reason:  err.Error(),
			},
		})
	}

	userId := int(claims["user_id"].(float64))

	savedWord := models.SavedWord{
		Id:     reqBody.Id,
		WordId: reqBody.WordId,
		UserId: userId,
	}

	err = h.savedWordService.CreateSavedWord(&savedWord)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Bir hata oluştu",
				Reason:  err.Error(),
			},
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(models.SuccessResponse{
		Success: true,
		Data:    "Kelime başarıyla öğrendikleriniz arasına eklendi.",
	})
}

/* func (h *WordHandler) handleGetWordById(ctx *fiber.Ctx) error {
	wordId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir id giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	word, err := h.wordService.GetWordById(wordId)

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

func (h *WordHandler) handleCreateWord(ctx *fiber.Ctx) error {
	var reqBody models.Word

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir body giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	word := models.Word{
		Vocabulary:  reqBody.Vocabulary,
		Definition:  reqBody.Definition,
		Sentence:    reqBody.Sentence,
		Translation: reqBody.Translation,
		Level:       reqBody.Level,
	}

	err := h.wordService.CreateWord(&word)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Bir hata oluştu",
				Reason:  err.Error(),
			},
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(models.SuccessResponse{
		Success: true,
		Data:    word,
	})
}

func (h *WordHandler) handleUpdateWordById(ctx *fiber.Ctx) error {
	wordId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir id giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	var reqBody models.Word

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir body giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	blog, err := h.wordService.GetWordById(wordId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Girilen id'ye ait kullanıcı bulunamadı",
				Reason:  err.Error(),
			},
		})
	}

	err = h.wordService.UpdateWordById(&blog, models.Word{
		Vocabulary:  reqBody.Vocabulary,
		Definition:  reqBody.Definition,
		Sentence:    reqBody.Sentence,
		Translation: reqBody.Translation,
		Level:       reqBody.Level,
	})

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
		Data:    blog,
	})
}

func (h *WordHandler) handleDeleteWordById(ctx *fiber.Ctx) error {
	wordId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir id giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	_, err = h.wordService.GetWordById(wordId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Girilen id'ye ait kullanıcı bulunamadı",
				Reason:  err.Error(),
			},
		})
	}

	err = h.wordService.DeleteWordById(wordId)

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
		Data:    "Kayıt başarıyla silindi.",
	})
} */

func (h *SavedWordHandler) SetRoutes(a *fiber.App) {
	wordGroup := a.Group("/saved-words")
	wordGroup.Get("/", h.handleGetAllSavedWords)
	wordGroup.Post("/", h.handleCreateSavedWord)
	/*
		 	wordGroup.Get("/:id<int>", h.handleGetWordById)
			wordGroup.Post("/", h.handleCreateWord)
			wordGroup.Put("/:id", h.handleUpdateWordById)
			wordGroup.Delete("/:id", h.handleDeleteWordById)
	*/
}

/* var WordSwaggerEndpoints = []Endpoint{
	EndPoint(GET, "/words", "Kelimeler", Params(IntQuery("page", true, "Pagination için sayfa numarası"), IntQuery("limit", true, "Bir sayfada dönecek toplam veri sayısı")), nil, []models.Word{}, models.FailureResponse{}, "Tüm kelimeleri listeler", nil),
	EndPoint(GET, "/words/{wordId}", "Kelimeler", Params(IntParam("wordId", true, "Bilgileri istenen kelimenin idsi")), nil, models.Word{}, models.FailureResponse{}, "Id'ye göre tek bir kelimenin bilgilerini döner", nil),
	EndPoint(POST, "/words", "Kelimeler", Params(), models.Word{}, models.Word{}, models.FailureResponse{}, "Yeni bir kelime oluşturur", nil),
	EndPoint(PUT, "/words/{wordId}", "Kelimeler", Params(IntParam("wordId", true, "Güncellenmek istenen kelimenin idsi")), models.Word{}, models.Word{}, models.FailureResponse{}, "Mevcut bir kelimenin bilgilerini günceller", nil),
	EndPoint(DELETE, "/words/{wordId}", "Kelimeler", Params(IntParam("wordId", true, "Silinmek istenen kelimenin idsi")), nil, []models.Word{}, models.FailureResponse{}, "Bir kelimeyi siler", nil),
}
*/
