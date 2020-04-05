package put

import (
	"encoding/json"
	"net/http"

	"github.com/jain-chetan/users-service/interfaces"
	"github.com/jain-chetan/users-service/model"
)

type PutHandler struct{}

func (p *PutHandler) PutUserHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("userID")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Getting the user details of the ID passed
	user, err := interfaces.DBClient.GetUserQuery(userID)
	if err != nil {
		response := ResponseMapper(400, "error in getting data")
		json.NewEncoder(w).Encode(response)
		return
	}

	//Decoding the body from the Json structure
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response := ResponseMapper(400, "error in getting response")
		json.NewEncoder(w).Encode(response)
	}

	//call to database query to update the fields of user
	err = interfaces.DBClient.UpdateUserQuery(userID, user)

	if err != nil {
		response := ResponseMapper(400, "Request error")
		json.NewEncoder(w).Encode(response)
	}
	response := ResponseMapper(200, "OK")
	json.NewEncoder(w).Encode(response)

}

//ResponseMapper - function to send back the response
func ResponseMapper(code int, message string) model.Response {
	var response model.Response
	response = model.Response{
		Code:    code,
		Message: message,
	}
	return response
}
