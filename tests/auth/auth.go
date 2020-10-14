package auth

import (
	"bytes"
	"encoding/json"
	"github.com/alainmucyo/my_brand/tests/data"
	"net/http"
	"testing"
)

func MockUser(t *testing.T) {
	res, err := http.Get(data.MainURL + "/api/auth/mock")
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected %d, received %d", http.StatusCreated, res.StatusCode)
	}
}

func CurrentLoggedInUser(t *testing.T) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", data.MainURL+"/api/auth/details", nil)
	req.Header.Set("Authorization", "Bearer "+(data.Token).(string))
	//res, err := http.Get(data.MainURL+"/api/auth/details")
	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

func Login(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		form, _ := json.Marshal(validData)
		res, err := http.Post(data.MainURL+"/api/auth/login", "application/json", bytes.NewBuffer(form))
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		var result map[string]map[string]interface{}
		json.NewDecoder(res.Body).Decode(&result)
		data.Token = result["data"]["token"]
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
		}
	})
	t.Run("Fail", func(t *testing.T) {
		form, _ := json.Marshal(invalidData)
		res, err := http.Post(data.MainURL+"/api/auth/login", "application/json", bytes.NewBuffer(form))
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("Expected %d, received %d", http.StatusUnprocessableEntity, res.StatusCode)
		}
	})

}
func Auth(t *testing.T) {
	t.Run("Mock user", MockUser)
	t.Run("Login", Login)
	t.Run("Current user", CurrentLoggedInUser)
}
