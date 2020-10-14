package middlewares

import (
	"encoding/json"
	"github.com/alainmucyo/my_brand/src/utils"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func IsAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := utils.TokenValid(request)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(errorResponse{Message: "Unauthorized"})
			return
		}
		next(writer, request)
	}
}
