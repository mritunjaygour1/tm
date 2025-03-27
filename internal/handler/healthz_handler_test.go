package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHealthz(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/healthz", NewHealthzHandler().GetHealthz)

	// Test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz", nil)
	router.ServeHTTP(w, req)

	// Assert
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, `{"message":"OK"}`, w.Body.String())
}
