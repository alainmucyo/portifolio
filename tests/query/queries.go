package query

import (
	"bytes"
	"encoding/json"
	"github.com/alainmucyo/my_brand/tests/data"
	"net/http"
	"testing"
)

func ListQueries(t *testing.T) {
	t.Run("Not listed", func(t *testing.T) {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", data.MainURL+"/api/query", nil)
		res, err := client.Do(req)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusUnauthorized {
			t.Errorf("Expected %d, received %d", http.StatusUnauthorized, res.StatusCode)
		}
	})
	t.Run("Listed", func(t *testing.T) {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", data.MainURL+"/api/query", nil)
		req.Header.Set("Authorization", "Bearer "+(data.Token).(string))
		res, err := client.Do(req)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
		}
	})

}
func CreateQuery(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		form, _ := json.Marshal(validData)
		res, err := http.Post(data.MainURL+"/api/query", "application/json", bytes.NewBuffer(form))
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
		}
	})
	t.Run("Not Create", func(t *testing.T) {
		form, _ := json.Marshal(invalidData)
		res, err := http.Post(data.MainURL+"/api/query", "application/json", bytes.NewBuffer(form))
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != 422 {
			t.Errorf("Expected %d, received %d", 422, res.StatusCode)
		}
	})

}
func Query(t *testing.T) {
	t.Run("Create Query", CreateQuery)
	t.Run("List queries", ListQueries)
}
