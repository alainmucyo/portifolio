package article

import (
	"github.com/alainmucyo/my_brand/tests/data"
	"net/http"
	"testing"
)

func ListArticles(t *testing.T) {
	res, err := http.Get(data.MainURL + "/api/article")
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

func Article(t *testing.T) {
	t.Run("List Articles", ListArticles)
}
