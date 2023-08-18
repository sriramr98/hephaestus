package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitub.com/sriramr98/hephaestus/mocks"
)

func TestHealthController(t *testing.T) {
	t.Run("Returns 200 status code", func(t *testing.T) {
		res := httptest.NewRecorder()
		gc := mocks.MockGin{
			ResponseRecorder: res,
		}
		context := gc.GetTestGinContext()

		GetHealth(context)

		assert.Equal(t, http.StatusOK, res.Code)

	})
}
