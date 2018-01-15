package petshop

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"PetCreate",
		"POST",
		"/pets",
		PetCreate,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"PetsList",
		"GET",
		"/pets",
		GetPets,
	},
	Route{
		"PetShow",
		"GET",
		"/pets/{petId}",
		PetShow,
	},
}
