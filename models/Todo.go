package models

import "github.com/kamva/mgm/v3"

type Todo struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
	Completed        bool   `json:"completed" bson:"completed"`
	Description      string `json:"description" bson:"description"`
}
