package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User structure for User document
type User struct {
	UserID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName         string             `json:"firstName" bson:"firstName"`
	LastName          string             `json:"lastName" bson:"lastName"`
	Phone             string             `json:"phone" bson:"phone"`
	Email             string             `json:"email" bson:"email"`
	Address           Address            `json:"address" bson:"address"`
	EncryptedPassword string             `json:"encryptedPassword" bson:"encryptedPassword"`
	IsDeleted         bool               `json:"-" bson:"IsDeleted"`
	CreatedAt         time.Time          `json:"CreatedAt" bson:"CreatedAt"`
}

//Address structure for sub document inside User Structure
type Address struct {
	AddressLine1 string `json:"addressLine1" bson:"addressLine1"`
	City         string `json:"city" bson:"city"`
	State        string `json:"state" bson:"state"`
	Zipcode      string `json:"zipcode" bson:"zipcode"`
}

//DBConfig has information required to connect to DB
type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

//Response has the message and code
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//CreateResponse - response structure for create where ID is sent as response
type CreateResponse struct {
	ID      primitive.ObjectID `json:"userID"`
	Code    int                `json:"code"`
	Message string             `json:"message"`
}
