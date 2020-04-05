package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type Address struct {
	AddressLine1 string `json:"addressLine1" bson:"addressLine1"`
	City         string `json:"city" bson:"city"`
	State        string `json:"state" bson:"state"`
	Zipcode      string `json:"zipcode" bson:"zipcode"`
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CreateResponse struct {
	ID      primitive.ObjectID `json:"productID"`
	Code    int                `json:"code"`
	Message string             `json:"message"`
}
