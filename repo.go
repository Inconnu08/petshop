package petshop

import "fmt"

var currentId int

var pets Pets

// Give us some seed data
func init() {
	RepoCreatePet(Pet{Name: "Ben"})
	RepoCreatePet(Pet{Name: "Hen"})
	RepoCreatePet(Pet{Name: "Len"})
}

func RepoFindPet(id int) Pet {
	for _, p := range pets {
		if p.Id == id {
			return p
		}
	}
	// return empty pet if not found
	return Pet{}
}

func RepoCreatePet(p Pet) Pet {
	currentId += 1
	p.Id = currentId
	pets = append(pets, p)
	return p
}

func RepoDestroyPet(id int) error {
	for i, t := range pets {
		if t.Id == id {
			pets = append(pets[:i], pets[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Pet with id of %d to delete!", id)
}
