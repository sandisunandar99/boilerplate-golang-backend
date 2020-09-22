package models

import (
	"github.com/Kamva/mgm/v3"
)

type Users struct {
	mgm.DefaultModel `bson:",inline"`

	Fullname string `json:"Fullname" bson:"Fullname"`
	Username string `json:"Username" bson:"Username"`
}

type M map[string]interface{}

var internalError = M{"message": "Internal error"}
var userNotFound = M{"message": "User not found!"}
