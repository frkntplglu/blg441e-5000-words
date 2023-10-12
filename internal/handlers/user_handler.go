package handlers

import (
	"regexp"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/frkntplglu/emir-backend/internal/config"
	"github.com/frkntplglu/emir-backend/internal/middleware"
	"github.com/frkntplglu/emir-backend/internal/models"
	"github.com/frkntplglu/emir-backend/internal/services"
	"github.com/frkntplglu/emir-backend/pkg/utils"
	. "github.com/go-swagno/swagno"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) handleGetAllUsers(ctx *fiber.Ctx) error {
	pagination := utils.GeneratePaginationFromCtx(ctx)

	users, err := h.userService.GetAllUsers(models.User{}, &pagination)

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
		Data:    users,
	})
}

func (h *UserHandler) handleGetUserById(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("userId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir userId giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	user, err := h.userService.GetUserById(userId)

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
		Data:    user,
	})
}

func (h *UserHandler) handleRegister(ctx *fiber.Ctx) error {
	var reqBody models.User

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir body giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	passwordCases := []string{".{8,}", "[a-z]", "[A-Z]", "[0-9]"}
	isPasswordPassed := true
	for _, passwordCase := range passwordCases {
		matched, _ := regexp.MatchString(passwordCase, reqBody.Password)

		if !matched {
			isPasswordPassed = false
			break
		}
	}

	if !isPasswordPassed {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Girdiğiniz şifre gereklilikleri karşılamamaktadır.",
				Reason:  "",
			},
		})
	}

	user := models.User{
		Firstname: reqBody.Firstname,
		Lastname:  reqBody.Lastname,
		Email:     reqBody.Email,
		Password:  utils.Hash(reqBody.Password),
		Address:   reqBody.Address,
		CreatedAt: reqBody.CreatedAt,
		LastLogin: reqBody.LastLogin,
	}

	err := h.userService.CreateUser(&user)

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
		Data:    user,
	})
}

func (h *UserHandler) handleUpdateUserById(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("userId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir userId giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	var reqBody models.User

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir body giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	user, err := h.userService.GetUserById(userId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Girilen id'ye ait kullanıcı bulunamadı",
				Reason:  err.Error(),
			},
		})
	}

	err = h.userService.UpdateUserById(&user, models.User{
		Firstname: reqBody.Firstname,
		Lastname:  reqBody.Lastname,
		Email:     reqBody.Email,
		Address:   reqBody.Address,
		CreatedAt: reqBody.CreatedAt,
		LastLogin: reqBody.LastLogin,
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
		Data:    user,
	})
}

func (h *UserHandler) handleDeleteUserById(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("userId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir userId giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	_, err = h.userService.GetUserById(userId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Girilen id'ye ait kullanıcı bulunamadı",
				Reason:  err.Error(),
			},
		})
	}

	err = h.userService.DeleteUserById(userId)

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
}

func (h *UserHandler) handleLogin(ctx *fiber.Ctx) error {
	var reqBody models.LoginBody

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Lütfen geçerli bir body giriniz.",
				Reason:  err.Error(),
			},
		})
	}

	reqBody.Password = utils.Hash(reqBody.Password)

	user, err := h.userService.GetUserByParams(models.User{
		Email:    reqBody.Email,
		Password: reqBody.Password,
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Kullanıcı adı ya da şifre yanlış.",
				Reason:  "",
			},
		})
	}

	err = h.userService.UpdateUserById(&user, models.User{
		LastLogin: time.Now(),
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Bir hata oluştu.",
				Reason:  err.Error(),
			},
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["full_name"] = user.Firstname + " " + user.Lastname
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["last_login"] = user.LastLogin
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte(config.ACCESS_TOKEN))

	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    t,
		HTTPOnly: true,
	})

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = user.Id
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	rt, err := refreshToken.SignedString([]byte(config.REFRESH_TOKEN))
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    rt,
		HTTPOnly: true,
	})

	return ctx.Status(fiber.StatusOK).JSON(models.SuccessResponse{
		Success: true,
		Data:    "Başarıyla giriş yaptınız..",
	})

}

func (h *UserHandler) handleLogout(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    ctx.Cookies("refreshToken"),
		Expires:  time.Now().Add(-3 * time.Second),
		HTTPOnly: true,
	})
	return ctx.Status(fiber.StatusOK).JSON(models.SuccessResponse{
		Success: true,
		Data:    "Başarıyla çıkış yapıldı..",
	})
}

func (h *UserHandler) handleRefreshToken(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("refreshToken")

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.REFRESH_TOKEN), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Geçersiz refresh token.",
				Reason:  err.Error(),
			},
		})
	}

	if !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Geçersiz refresh token.",
				Reason:  "Token süresi dolmuş veya geçersiz.",
			},
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Bir hata oluştu.",
				Reason:  "Token içeriği çözümlenemedi.",
			},
		})
	}

	userId := claims["sub"]

	user, err := h.userService.GetUserByParams(models.User{
		Id: int(userId.(float64)),
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Bir hata oluştu",
				Reason:  "",
			},
		})
	}

	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessClaims := accessToken.Claims.(jwt.MapClaims)
	accessClaims["full_name"] = user.Firstname + " " + user.Lastname
	accessClaims["email"] = user.Email
	accessClaims["address"] = user.Address
	accessClaims["last_login"] = user.LastLogin
	accessClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	newAccessToken, err := accessToken.SignedString([]byte(config.ACCESS_TOKEN))

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(models.SuccessResponse{
		Success: true,
		Data: struct {
			AccessToken string `json:"accessToken"`
		}{
			AccessToken: newAccessToken,
		},
	})
}

func (h *UserHandler) SetRoutes(a *fiber.App) {
	userGroup := a.Group("/users")
	userGroup.Get("/", middleware.Protected(), h.handleGetAllUsers)
	userGroup.Get("/:userId", h.handleGetUserById)
	userGroup.Put("/:userId", h.handleUpdateUserById)
	userGroup.Delete("/:userId", h.handleDeleteUserById)
	authGroup := a.Group("/auth")
	authGroup.Post("/login", h.handleLogin)
	authGroup.Post("/register", h.handleRegister)
	authGroup.Get("/refresh-token", h.handleRefreshToken)
	authGroup.Get("/logout", h.handleLogout)
}

var UserSwaggerEndpoints = []Endpoint{
	EndPoint(GET, "/users", "Kullanıcılar", Params(IntQuery("page", true, "Pagination için sayfa numarası"), IntQuery("limit", true, "Bir sayfada dönecek toplam veri sayısı")), nil, []models.User{}, models.FailureResponse{}, "Tüm kullanıcıları listeler", nil),
	EndPoint(GET, "/users/{userId}", "Kullanıcılar", Params(IntParam("userId", true, "Bilgileri istenen user'ın idsi")), nil, models.User{}, models.FailureResponse{}, "Id'ye göre tek bir kullanıcı bilgilerini döner", nil),
	EndPoint(POST, "/users", "Kullanıcılar", Params(), models.User{}, models.User{}, models.FailureResponse{}, "Yeni bir kullanıcı oluşturur", nil),
	EndPoint(PUT, "/users/{userId}", "Kullanıcılar", Params(IntParam("userId", true, "Bilgileri güncellenmek istenen user'ın idsi")), models.User{}, models.User{}, models.FailureResponse{}, "Mevcut bir kullanıcının bilgilerini günceller", nil),
	EndPoint(DELETE, "/users/{userId}", "Kullanıcılar", Params(IntParam("userId", true, "Silinmek istenen user'ın idsi")), nil, []models.User{}, models.FailureResponse{}, "Bir kullanıcıyı siler", nil),
}
