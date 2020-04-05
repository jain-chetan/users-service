package put

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jain-chetan/users-service/interfaces"
	"github.com/jain-chetan/users-service/model"
)

type PutHandler struct{}

func (p *PutHandler) PutUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	userID := r.Header.Get("userID")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	if err != nil {
		response := ResponseMapper(400, "error in getting response")
		json.NewEncoder(w).Encode(response)
	}

	err = interfaces.DBClient.UpdateUserQuery(userID, user)

	if err != nil {
		response := ResponseMapper(400, "Request error")
		json.NewEncoder(w).Encode(response)
	}
	response := ResponseMapper(200, "OK")
	json.NewEncoder(w).Encode(response)

}

func ResponseMapper(code int, message string) model.Response {
	var response model.Response
	response = model.Response{
		Code:    code,
		Message: message,
	}
	return response
}
