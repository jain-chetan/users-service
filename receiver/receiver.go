package receiver

import (
	delete "github.com/jain-chetan/users-service/handlers/delete"
	get "github.com/jain-chetan/users-service/handlers/get"
	post "github.com/jain-chetan/users-service/handlers/post"
	put "github.com/jain-chetan/users-service/handlers/put"
)

var Post post.PostHandler
var Get get.GetHandler
var Delete delete.DeleteHandler
var Put put.PutHandler
