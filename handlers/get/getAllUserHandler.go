package get

import (
	"encoding/json"
	"net/http"

	"github.com/jain-chetan/users-service/interfaces"
	"github.com/jain-chetan/users-service/model"
)

type GetHandler struct{}

func (g *GetHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {

	//Getting the userID from the headers
	userID := r.Header.Get("userID")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Call to the database query to get the User details
	userDetails, err := interfaces.DBClient.GetUserQuery(userID)
	if err != nil {
		response := ResponseMapper(400, "Bad Request")
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(userDetails)

}

func ResponseMapper(code int, message string) model.Response {
	var response model.Response
	response = model.Response{
		Code:    code,
		Message: message,
	}
	return response
}
