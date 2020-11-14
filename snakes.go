package pets

import (
	"math/rand"
)

type Snake struct {
	Pet
}

func NewSnake(name string, breed string, age int) Snake {
	snake := Snake{}
	snake.Pet = NewPet(name, "Snake", breed, age)
	return snake
}

func RandomSnake() Snake {
	name := Names[rand.Intn(len(Names))]
	breed := SnakeBreeds[rand.Intn(len(SnakeBreeds))]
	age := rand.Intn(23) + 1
	return NewSnake(name, breed, age)
}

var SnakeBreeds = []string{"King Cobra",
	"American Copperhead",
	"Black Mamba",
	"Corn Snake",
	"Rattlesnake",
	"Boa Constrictor",
	"Eastern Coral Snake",
	"Black Rat Snake",
	"Burmese Python",
	"Ball Python",
	"Royal Python",
	"Reticulated Python",
	"Garter Snake",
	"Green Anaconda",
	"Water Moccasin Snake",
	"Green Tree Python" }
