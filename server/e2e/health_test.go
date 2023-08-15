package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitub.com/sriramr98/hephaestus/router"
)

func TestHealthController(t *testing.T) {
	t.Run("Retrns 200 status code", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/health", nil)

		router := router.GetGinRouter()

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

	})
}
