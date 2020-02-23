package pets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCat(t *testing.T) {
	p := NewCat("Bella", "American Shorthair", 5)
	assert.Equal(t, "Cat", p.Species)
	assert.Equal(t, "Bella", p.Name)
	assert.Equal(t, "American Shorthair", p.Breed)
	assert.Equal(t, 5, p.Age)
}

func TestRandomCat(t *testing.T) {
	c1 := RandomCat()
	c2 := RandomCat()
	//log.Println(c1)
	//log.Println(c2)
	// Assert that the pets are not identical
	assert.NotEqual(t, c1, c2, "Two random cats seem to be identical.")
}
