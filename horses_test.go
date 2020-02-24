package pets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHorse(t *testing.T) {
	p := NewHorse("Hannibal", "Saint Bernard", 9)
	assert.Equal(t, "Horse", p.Species)
	assert.Equal(t, "Hannibal", p.Name)
	assert.Equal(t, "Saint Bernard", p.Breed)
	assert.Equal(t, 9, p.Age)
}

func TestRandomHorse(t *testing.T) {
	d1 := RandomCat()
	d2 := RandomCat()
	//log.Println(d1)
	//log.Println(d2)
	// Assert that the pets are not identical
	assert.NotEqual(t, d1, d2, "Two random horses seem to be identical.")
}
