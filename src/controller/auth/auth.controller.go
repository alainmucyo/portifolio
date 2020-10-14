package auth

import (
	"encoding/json"
	"fmt"
	"github.com/alainmucyo/my_brand/src/model"
	"github.com/alainmucyo/my_brand/src/resources"
	"github.com/alainmucyo/my_brand/src/utils"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func MockUser(writer http.ResponseWriter, request *http.Request) {
	request.Header.Set("Content-Type", "application/json")
	imageURL := "https://res.cloudinary.com/alainmucyo/image/upload/v1601102855/keboidjyp4mcx38kevhj.png"
	mockedUser := model.User{Name: "Alain MUCYO", Email: "alainmucyo3@gmail.com", Password: "password", Image: imageURL}
	mockedUser.Prepare()

	user, err := mockedUser.Save()
	if err != nil {
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode("Server error")
		return
	}
	writer.WriteHeader(http.StatusCreated)
	user.Password = "hidden"
	response := resources.JsonResponse("User mocked successfully", user)
	json.NewEncoder(writer).Encode(response)

}

func Login(writer http.ResponseWriter, request *http.Request) {
	request.Header.Set("Content-Type", "application/json")
	var user model.User
	json.NewDecoder(request.Body).Decode(&user)
	loggedUser, err := user.CheckUser()
	if err != nil {
		writer.WriteHeader(422)
		json.NewEncoder(writer).Encode(err.Error())
		return
	}
	token, err := utils.CreateToken(loggedUser.ID)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode("Server error")
		return
	}
	userResponse := model.UserResponse{User: loggedUser, Token: token}
	response := resources.JsonResponse("Log in success!", userResponse)
	json.NewEncoder(writer).Encode(response)
}

func UserDetails(writer http.ResponseWriter, request *http.Request) {
	request.Header.Set("Content-Type", "application/json")
	userID, err := utils.ExtractTokenID(request)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(writer).Encode("Unauthorized")
		return
	}
	user, err := model.User{}.FindById(userID)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(writer).Encode("Unauthorized")
		return
	}
	user.Password = "hidden"
	response := resources.JsonResponse("User details!", user)
	json.NewEncoder(writer).Encode(response)
}

func Profile(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "multipart/form-data")
	userID, err := utils.ExtractTokenID(request)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(writer).Encode("Unauthorized")
		return
	}
	user, err := model.User{}.FindById(userID)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(writer).Encode("Unauthorized")
		return
	}
	user.Name = request.FormValue("name")
	user.Email = request.FormValue("email")
	validate := validator.New()
	validationError := validate.Struct(user)
	if validationError != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		for _, err := range validationError.(validator.ValidationErrors) {
			json.NewEncoder(writer).Encode(err.Error())
		}
		return
	}
	if request.FormValue("old_password") != "" {
		oldPassword := request.FormValue("old_password")
		err := model.VerifyPassword(user.Password, oldPassword)
		if err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(writer).Encode("Wrong old password")
			return
		}
		hashedPassword, _ := model.Hash(request.FormValue("password"))
		user.Password = string(hashedPassword)
	}
	file, _, err := request.FormFile("image")
	if file != nil {
		uploadedFile, err := utils.FileUpload(request, "image")
		if err != nil {
			fmt.Println("Error: ", err)
		}
		user.Image = uploadedFile
	}
	updatedUser, err := user.Save()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode("Server error")
		return
	}
	updatedUser.Password = "hidden"
	response := resources.JsonResponse("User details!", updatedUser)
	json.NewEncoder(writer).Encode(response)
}
