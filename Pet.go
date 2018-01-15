package petshop

import "time"

type Pet struct {
	Id     int       `json:"id"`
	Name   string    `json:"name"`
	Animal string    `json:"animal"`
	Price  int       `json:"price"`
	Due    time.Time `json:"due"`
}

type Pets []Pet
