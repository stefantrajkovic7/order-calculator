package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"backend/api/handlers"
	"backend/internal/config"

	"github.com/gin-gonic/gin"
)

func TestCalculateEndpoint(t *testing.T) {
	// Initialize test config
	cfg := &config.Config{PackSizes: []int{250, 500, 1000, 2000, 5000}}

	// Set up router and handler
	r := gin.Default()
	r.POST("/calculate", handlers.CalculateHandler(cfg))

	// Create test request payload
	payload := map[string]int{"order": 501}
	body, _ := json.Marshal(payload)

	// Create test HTTP request
	req := httptest.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Validate response
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status %d but got %d", http.StatusOK, w.Code)
	}

	var response map[string]map[int]int
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	expected := map[int]int{500: 1, 250: 1}
	if !reflect.DeepEqual(response["packs"], expected) {
		t.Errorf("Expected %v but got %v", expected, response["packs"])
	}
}

func TestUpdatePackSizesEndpoint(t *testing.T) {
	// Initialize test config
	cfg := &config.Config{PackSizes: []int{250, 500, 1000, 2000, 5000}}

	// Set up router and handler
	r := gin.Default()
	r.POST("/packs", handlers.UpdatePackSizesHandler(cfg))

	// Create test request payload
	payload := map[string][]int{"pack_sizes": {300, 600, 1200}}
	body, _ := json.Marshal(payload)

	// Create test HTTP request
	req := httptest.NewRequest("POST", "/packs", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Validate response
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status %d but got %d", http.StatusOK, w.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	expected := "Pack sizes updated successfully"
	if response["message"] != expected {
		t.Errorf("Expected %v but got %v", expected, response["message"])
	}

	// Verify config update
	expectedPackSizes := []int{300, 600, 1200}
	if !reflect.DeepEqual(cfg.PackSizes, expectedPackSizes) {
		t.Errorf("Expected pack sizes %v but got %v", expectedPackSizes, cfg.PackSizes)
	}
}
