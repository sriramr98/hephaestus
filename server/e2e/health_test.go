package e2e

import (
	"github.com/stretchr/testify/assert"
	"gitub.com/sriramr98/hephaestus/mocks"
	"gitub.com/sriramr98/hephaestus/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthRoute(t *testing.T) {
	t.Run("Returns 200 status code", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		request, err := http.NewRequest(http.MethodGet, "/health", nil)
		if err != nil {
			t.Fatal(err)
		}
		mockGin := mocks.MockGin{
			Request:          request,
			ResponseRecorder: responseRecorder,
		}

		mockGin.ServeTestEngine(router.GetRouterEngine())

		assert.Equal(t, http.StatusOK, responseRecorder.Code)
	})
}
