package models

import "github.com/kamva/mgm/v3"

type Product struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
	Category         string `json:"category" bson:"category"`
	Description      string `json:"description" bson:"description"`
	Image            string `json:"image" bson:"image"`
	Price            int    `json:"price" bson:"price"`
}
