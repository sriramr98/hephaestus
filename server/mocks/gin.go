package mocks

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type MockGin struct {
	ResponseRecorder httptest.ResponseRecorder
	Request          http.Request
}

func (mc MockGin) GetTestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(&mc.ResponseRecorder)
	// ctx.Request = &mc.Request

	return ctx
}
