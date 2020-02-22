package pets

type Pet struct {
	Name    string	`json:"name"`
	Species string	`json:"species"`
	Breed   string	`json:"breed"`
	Age     int		`json:"age"`
}

func NewPet(name string, species string, breed string, age int) *Pet {
	newPet :=  &Pet{Name:name, Species:species, Breed:breed, Age:age}
	return newPet
}
