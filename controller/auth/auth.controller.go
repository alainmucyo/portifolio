package auth

import (
	"encoding/json"
	"fmt"
	"github.com/alainmucyo/my_brand/model"
	"github.com/alainmucyo/my_brand/resources"
	"github.com/alainmucyo/my_brand/utils"
	"net/http"
)

func MockUser(writer http.ResponseWriter, request *http.Request) {
	request.Header.Set("Content-Type", "application/json")
	mockedUser := model.User{Name: "Alain MUCYO", Email: "alainmucyo@gmail.com", Password: "password"}
	mockedUser.Prepare()

	user, err := mockedUser.Create()
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

func UserDetails(writer http.ResponseWriter, request *http.Request)  {
	request.Header.Set("Content-Type", "application/json")
	userID,err := utils.ExtractTokenID(request)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(writer).Encode("Unauthorized")
		return
	}
	user,err := model.User{}.FindById(userID)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(writer).Encode("Unauthorized")
		return
	}
	user.Password="hidden"
	response := resources.JsonResponse("User details!", user)
	json.NewEncoder(writer).Encode(response)
}