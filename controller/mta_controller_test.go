package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoutes(t *testing.T) {
	// Set the value of THRESHOLD environment variable
	_ = os.Setenv("THRESHOLD", "2")

	router := gin.New()
	RegisterRoutes(router)

	req, err := http.NewRequest("GET", "/getHostnames", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []string
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))

	expectedResult := []string{"mta-prod-1", "mta-prod-2", "mta-prod-3"}
	assert.ElementsMatch(t, expectedResult, response)
}
