package config

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() {
	error := mgm.SetDefaultConfig(nil, "odm_test", options.Client().ApplyURI("mongodb+srv://hatwaarbeta:mongo7038@devcluster0.hdvnq.mongodb.net/"))
	if error != nil {
		log.Fatal(nil)
	}
}
