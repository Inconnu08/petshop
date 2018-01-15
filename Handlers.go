package petshop

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	//"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	text := "Hello Petshop!"
	if err := json.NewEncoder(w).Encode(text); err != nil {
		panic(err)
	}
}

func PetShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//fmt.Fprintln(w, r)
	petId, _ := strconv.Atoi(vars["petId"])
	for _, p := range pets {
		if p.Id-1 == petId {
			json.NewEncoder(w).Encode(pets[p.Id-1])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func GetPets(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(pets)
}

func PetCreate(w http.ResponseWriter, r *http.Request) {
	var pet Pet
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &pet); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	p := RepoCreatePet(pet)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}
