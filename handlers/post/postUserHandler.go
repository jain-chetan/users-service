package post

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jain-chetan/users-service/interfaces"
	"github.com/jain-chetan/users-service/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct{}

func (p *PostHandler) PostUserHandler(w http.ResponseWriter, r *http.Request) {

	var user model.User

	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)

	if err != nil {
		response := ResponseMapper(400, "error in getting response")
		json.NewEncoder(w).Encode(response)
	}

	result, err := interfaces.DBClient.CreateUserQuery(user)
	if err != nil {
		response := ResponseMapper(400, "error inserting records")
		json.NewEncoder(w).Encode(response)
	}

	response := ResponseMapperCreate(200, "OK", result.ID)
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

func ResponseMapperCreate(code int, message string, id primitive.ObjectID) model.CreateResponse {
	var response model.CreateResponse
	response = model.CreateResponse{
		Code:    code,
		Message: message,
		ID:      id,
	}
	return response
}
