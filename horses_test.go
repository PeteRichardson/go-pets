package pets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDog(t *testing.T) {
	p := NewDog("Hannibal", "Saint Bernard", 9)
	assert.Equal(t, "Dog", p.Species)
	assert.Equal(t, "Hannibal", p.Name)
	assert.Equal(t, "Saint Bernard", p.Breed)
	assert.Equal(t, 9, p.Age)
}

func TestRandomDog(t *testing.T) {
	d1 := RandomCat()
	d2 := RandomCat()
	//log.Println(d1)
	//log.Println(d2)
	// Assert that the pets are not identical
	assert.NotEqual(t, d1, d2, "Two random dogs seem to be identical.")
}
