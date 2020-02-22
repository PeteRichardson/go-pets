package pets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPet(t *testing.T) {
	p := NewPet("Bella", "Cat", "American Shorthair", 5)
	assert.Equal(t, "Cat", p.Species)
	assert.Equal(t, "Bella", p.Name)
	assert.Equal(t, "American Shorthair", p.Breed)
	assert.Equal(t, 5, p.Age)
}