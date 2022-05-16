package controller

import (
	"bytes"
	"encoding/json"
	"kentang/config"
	"kentang/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func CreatePhone(name, price, design, display, perfomance, cameras, audio, battery, features string) *models.Phone {
	DB := config.Connect()
	phone := models.Phone{
		Name:       name,
		Price:      price,
		Design:     design,
		Display:    display,
		Perfomance: perfomance,
		Cameras:    cameras,
		Audio:      audio,
		Battery:    battery,
		Features:   features,
	}

	if err := DB.Create(&phone).Error; err != nil {
		return nil
	}

	return &phone
}

func TruncatePhonesTable() {
	DB := config.Connect()
	DB.Exec("TRUNCATE TABLE phones")
}

func TestCreatePhoneController(t *testing.T) {
	config.InitialMigration()

	token := GetToken()

	testCases := []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "create phone normal",
			path:       "/jwt/redirected/phones",
			expectCode: http.StatusOK,
		},
	}

	reqBody := map[string]interface{}{
		"name":       "Iphone X",
		"price":      "4000000",
		"design":     "89",
		"display":    "73",
		"perfomance": "79",
		"cameras":    "80",
		"audio":      "70",
		"battery":    "88",
		"features":   "69",
	}
	reqBodyJSON, _ := json.Marshal(reqBody)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/phones", bytes.NewReader(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreatePhoneController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response struct {
				Phone models.Phone `json:"data"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, response.Phone.Name, "Iphone X")
			assert.Equal(t, response.Phone.Price, "4000000")
			assert.Equal(t, response.Phone.Design, "89")
			assert.Equal(t, response.Phone.Display, "73")
			assert.Equal(t, response.Phone.Perfomance, "79")
			assert.Equal(t, response.Phone.Cameras, "80")
			assert.Equal(t, response.Phone.Audio, "70")
			assert.Equal(t, response.Phone.Battery, "88")
			assert.Equal(t, response.Phone.Features, "69")
		}
	}

	TruncateUsersTable()
	TruncatePhonesTable()
}
