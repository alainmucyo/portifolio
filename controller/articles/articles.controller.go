package articles

import (
	"encoding/json"
	"fmt"
	"github.com/alainmucyo/my_brand/model"
	"github.com/alainmucyo/my_brand/resources"
	"github.com/alainmucyo/my_brand/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var validate *validator.Validate

func Index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	articles, err := model.Article{}.Get()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	response := resources.JsonResponse("Articles", articles)
	json.NewEncoder(writer).Encode(response)
}
func Store(writer http.ResponseWriter, request *http.Request) {
	//writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Content-Type", "multipart/form-data")
	fmt.Println("Title: ", request.FormValue("title"))
	_, _, err := request.FormFile("image")
	if err != nil {
		writer.WriteHeader(422)
		json.NewEncoder(writer).Encode("Image is required")
		return
	}
	uploadedFile, err := utils.FileUpload(request, "image")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	validate = validator.New()
	var article model.Article
	//json.NewDecoder(request.Body).Decode(&article)
	article.Title = request.FormValue("title")
	article.Content = request.FormValue("content")
	article.Image = uploadedFile
	validationError := validate.Struct(article)
	if validationError != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		for _, err := range validationError.(validator.ValidationErrors) {
			json.NewEncoder(writer).Encode(err.Error())
		}

		return
	}
	createdArticle, err := article.Create()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		json.NewEncoder(writer).Encode("Server error")
		return
	}
	response := resources.JsonResponse("Article Created successfully!", createdArticle)
	json.NewEncoder(writer).Encode(response)
}
func Show(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	articleId, _ := strconv.Atoi(id)
	article, err := model.Article{}.FindById(articleId)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode("Article not found!")
		return
	}
	response := resources.JsonResponse("Article found", article)
	json.NewEncoder(writer).Encode(response)
}
func Update(writer http.ResponseWriter, request *http.Request) {
	validate = validator.New()
	writer.Header().Set("Content-Type", "application/json")
	id := mux.Vars(request)["id"]
	articleId, _ := strconv.Atoi(id)
	var article model.Article
	json.NewDecoder(request.Body).Decode(&article)
	validationError := validate.Struct(article)
	if validationError != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		for _, err := range validationError.(validator.ValidationErrors) {
			json.NewEncoder(writer).Encode(err.Error())
		}

		return
	}
	updatedArticle, err := article.Update(uint64(articleId))
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode("Article not found")
		return
	}
	updatedArticle.ID = uint64(articleId)
	response := resources.JsonResponse("Article Updated successfully!", updatedArticle)
	json.NewEncoder(writer).Encode(response)
}
func Delete(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	articleId, _ := strconv.Atoi(id)
	err := model.Article{}.Delete(uint64(articleId))
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode("Article not found!")
		return
	}
	response := resources.JsonResponse("Article Deleted", "")
	json.NewEncoder(writer).Encode(response)
}
