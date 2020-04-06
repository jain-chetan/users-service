package delete

import (
	"encoding/json"
	"net/http"

	"github.com/jain-chetan/users-service/interfaces"
	"github.com/jain-chetan/users-service/model"
)

type DeleteHandler struct{}

func (d *DeleteHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	//Getting the userID from the headers
	userID := r.Header.Get("userID")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	userExists := interfaces.DBClient.CheckUserExist(userID)

	if !userExists {
		response := ResponseMapper(400, "No records found")
		json.NewEncoder(w).Encode(response)
		return
	}

	//Call to database query for deleting a user
	err := interfaces.DBClient.DeleteUserQuery(userID)
	if err != nil {
		response := ResponseMapper(400, "Database error")
		json.NewEncoder(w).Encode(response)
	} else {
		response := ResponseMapper(200, "OK")
		json.NewEncoder(w).Encode(response)
	}

}

func ResponseMapper(code int, message string) model.Response {
	var response model.Response
	response = model.Response{
		Code:    code,
		Message: message,
	}
	return response
}
