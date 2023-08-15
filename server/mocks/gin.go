package mocks

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type MockGin struct {
	Method           string
	ResponseRecorder httptest.ResponseRecorder
}

func (mc MockGin) GetTestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(&mc.ResponseRecorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}
