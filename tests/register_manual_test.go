package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"slovenia_petconnect/controllers"
	"slovenia_petconnect/models"
	"testing"
)

func TestRegisterWithEmailUser(t *testing.T) {

	gin.SetMode(gin.TestMode)
	SetupTestDB()

	router := gin.Default()
	router.POST("/auth/signup/email", controllers.RegisterWithEmailUser)

	// valid payload
	payload := models.RegisterUserRequest{
		Username: "denis",
		Email:    "denis@example.com",
		Password: "strongPassword",
		Location: "Bosnia",
	}

	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/auth/signup/email", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	req, err := http.NewRequest("POST", "/auth/signup/email", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d. Reponse: %s", rec.Code, rec.Body.String())
	}
}
