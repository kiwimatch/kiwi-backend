package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1liale/maze-backend/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSystemCheck_Success(t *testing.T) {
	record := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(record) // create gin context for record

	engine.GET("/api-health", handlers.SystemCheck)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/api-health", nil)

	engine.ServeHTTP(record, ctx.Request)
	assert.Equal(t, http.StatusOK, record.Code)
}

func TestSystemCheck_Failure(t *testing.T) {
	record := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(record) // create gin context for record

	engine.GET("/api-health", handlers.SystemCheck)
	ctx.Request = httptest.NewRequest(http.MethodPost, "/api-health", nil)

	engine.ServeHTTP(record, ctx.Request)
	assert.Equal(t, http.StatusNotFound, record.Code)
}
