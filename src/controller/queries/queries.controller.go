package queries

import (
	"encoding/json"
	"fmt"
	"github.com/alainmucyo/my_brand/src/model"
	"github.com/alainmucyo/my_brand/src/resources"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate *validator.Validate

func GetQuery(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	queries, err := model.Query{}.FindAll()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	response := resources.JsonResponse("Queries", queries)
	json.NewEncoder(writer).Encode(response)
}
func CreateQuery(writer http.ResponseWriter, request *http.Request) {
	validate = validator.New()
	writer.Header().Set("Content-Type", "application/json")
	var query model.Query
	json.NewDecoder(request.Body).Decode(&query)
	query.Prepare()
	validationError := validate.Struct(query)
	if validationError != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		for _, err := range validationError.(validator.ValidationErrors) {
			json.NewEncoder(writer).Encode(err.Error())
		}

		return
	}
	createdQuery, err := query.Save()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode("Server error")
		return
	}
	json.NewEncoder(writer).Encode(createdQuery)
}
