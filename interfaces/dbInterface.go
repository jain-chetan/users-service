package interfaces

import (
	"github.com/jain-chetan/users-service/model"
)

var DBClient DBInteractions

type DBInteractions interface {
	DBConnect(model.DBConfig) error
	CreateUserQuery(user model.User) (model.CreateResponse, error)
	UpdateUserQuery(userID string, user model.User) error
	DeleteUserQuery(string) error
	GetUserQuery(string) (model.User, error)
}
