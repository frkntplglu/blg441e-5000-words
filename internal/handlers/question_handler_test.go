package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/frkntplglu/emir-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestQuestionHandlerIntegration(t *testing.T) {

	var app = fiber.New()
	testingSetup.questionHandler.SetRoutes(app)

	t.Run("Get All Questions", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/questions?limit=10", nil)
		assert.NoError(t, err)
		res, err := app.Test(req, 5000)
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		jsonDataFromHttp, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		var response = struct {
			Success bool              `json:"success"`
			Data    []models.Question `json:"data"`
		}{}
		err = json.Unmarshal(jsonDataFromHttp, &response)

		assert.NoError(t, err)
		assert.Equal(t, true, response.Success)
		assert.NotEmpty(t, response.Data)
		assert.Greater(t, len(response.Data), 0)

	})

	t.Run("Get Question By Id", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/questions/1", nil)
		assert.NoError(t, err)
		res, err := app.Test(req, 5000)
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		jsonDataFromHttp, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		var response = struct {
			Success bool            `json:"success"`
			Data    models.Question `json:"data"`
		}{}

		err = json.Unmarshal(jsonDataFromHttp, &response)

		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.NotEmpty(t, response.Data)

	})

	t.Run("Get Question By Quiz Id", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/questions?quiz_id=1", nil)
		assert.NoError(t, err)
		res, err := app.Test(req, 5000)
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		jsonDataFromHttp, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		var response = struct {
			Success bool              `json:"success"`
			Data    []models.Question `json:"data"`
		}{}

		err = json.Unmarshal(jsonDataFromHttp, &response)

		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.NotEmpty(t, response.Data)

	})

}
