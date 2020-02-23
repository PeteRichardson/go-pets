package pets

import (
	"math/rand"
)

type Cat struct {
	Pet
	HasHairball bool
}

func NewCat(name string, breed string, age int) *Cat {
	cat := Cat{}
	cat.Pet = *NewPet(name, "Cat", breed, age)
	cat.HasHairball = false
	return &cat
}

func RandomCat() *Cat {
	name := Names[rand.Intn(len(Names))]
	breed := CatBreeds[rand.Intn(len(CatBreeds))]
	age := rand.Intn(23) + 1
	return NewCat(name, breed, age)
}

var CatBreeds = []string{"Abyssinian", "American Bobtail",
	"American Curl", "American Shorthair", "American Wirehair", "Balinese",
	"Bengal Cats", "Birman", "Bombay", "British Shorthair", "Burmese",
	"Burmilla", "Chartreux", "Chinese Li Hua", "Colorpoint Shorthair",
	"Cornish Rex", "Cymric", "Devon Rex", "Egyptian Mau", "European Burmese",
	"Exotic", "Havana Brown", "Himalayan", "Japanese Bobtail",
	"Javanese", "Korat", "LaPerm", "Maine Coon", "Manx", "Nebelung",
	"Norwegian Forest", "Ocicat", "Oriental", "Persian", "Pixie-Bob",
	"Ragamuffin", "Ragdoll Cats", "Russian Blue", "Savannah", "Scottish Fold",
	"Selkirk Rex", "Siamese Cat", "Siberian", "Singapura",
	"Snowshoe", "Somali", "Sphynx", "Tonkinese", "Turkish Angora", "Turkish Van"}
