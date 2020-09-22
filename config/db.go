package config

import (
	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// Setup mgm default config
	err := mgm.SetDefaultConfig(nil, "echo test db", options.Client().ApplyURI("mongodb://localhost:27017/golang_be"))
}
