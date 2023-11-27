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

func TestQuizHandlerIntegration(t *testing.T) {

	var app = fiber.New()
	testingSetup.quizHandler.SetRoutes(app)

	t.Run("Get All Quizzes", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/quizzes?limit=10", nil)
		assert.NoError(t, err)
		res, err := app.Test(req, 5000)
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		jsonDataFromHttp, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		var response = struct {
			Success bool          `json:"success"`
			Data    []models.Quiz `json:"data"`
		}{}
		err = json.Unmarshal(jsonDataFromHttp, &response)

		assert.NoError(t, err)
		assert.Equal(t, true, response.Success)
		assert.NotEmpty(t, response.Data)
		assert.Greater(t, len(response.Data), 0)

	})

	t.Run("Get Quiz By Id", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/quizzes/1", nil)
		assert.NoError(t, err)
		res, err := app.Test(req, 5000)
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		jsonDataFromHttp, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		var response = struct {
			Success bool        `json:"success"`
			Data    models.Quiz `json:"data"`
		}{}

		err = json.Unmarshal(jsonDataFromHttp, &response)

		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.NotEmpty(t, response.Data)

	})

}
