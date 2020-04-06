package post

import (
	"crypto/sha1"
	"encoding/hex"
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//Reading the data from the JSON body
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)

	if err != nil {
		response := ResponseMapper(400, "error in getting response")
		json.NewEncoder(w).Encode(response)
	}

	//Encrypting the password and storing it in database
	user.EncryptedPassword = Encrypt(user.EncryptedPassword)

	//Call to the database query to create a user
	result, err := interfaces.DBClient.CreateUserQuery(user)
	if err != nil {
		response := ResponseMapper(400, "error inserting records")
		json.NewEncoder(w).Encode(response)
	}

	response := ResponseMapperCreate(200, "OK", result.ID)
	json.NewEncoder(w).Encode(response)

}

//Function to encrypt the password
func Encrypt(pwd string) string {
	h := sha1.New()
	h.Write([]byte(pwd))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func ResponseMapper(code int, message string) model.Response {
	var response model.Response
	response = model.Response{
		Code:    code,
		Message: message,
	}
	return response
}

//Response mapper for create - sending the ID of the user created.
func ResponseMapperCreate(code int, message string, id primitive.ObjectID) model.CreateResponse {
	var response model.CreateResponse
	response = model.CreateResponse{
		Code:    code,
		Message: message,
		ID:      id,
	}
	return response
}
