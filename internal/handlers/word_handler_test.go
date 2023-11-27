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

func TestWordHandlerIntegration(t *testing.T) {

	var app = fiber.New()
	testingSetup.wordHandler.SetRoutes(app)

	t.Run("Get All Words with limit = 5", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/words?limit=5", nil)
		assert.NoError(t, err)
		res, err := app.Test(req, 5000)
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		jsonDataFromHttp, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		var response = struct {
			Success bool          `json:"success"`
			Data    []models.Word `json:"data"`
		}{}
		err = json.Unmarshal(jsonDataFromHttp, &response)

		assert.NoError(t, err)
		assert.Equal(t, true, response.Success)
		assert.NotEmpty(t, response.Data)
		assert.Greater(t, len(response.Data), 0)

	})

	t.Run("Get Word By Id", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/words/1", nil)
		assert.NoError(t, err)
		res, err := app.Test(req, 5000)
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		jsonDataFromHttp, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		var response = struct {
			Success bool        `json:"success"`
			Data    models.Word `json:"data"`
		}{}

		err = json.Unmarshal(jsonDataFromHttp, &response)

		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.NotEmpty(t, response.Data)

	})

}
