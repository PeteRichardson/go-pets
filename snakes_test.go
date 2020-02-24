package pets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSnake(t *testing.T) {
	p := NewSnake("King", "Cobra", 7)
	assert.Equal(t, "Snake", p.Species)
	assert.Equal(t, "King", p.Name)
	assert.Equal(t, "Cobra", p.Breed)
	assert.Equal(t, 7, p.Age)
}

func TestRandomSnake(t *testing.T) {
	d1 := RandomCat()
	d2 := RandomCat()
	//log.Println(d1)
	//log.Println(d2)
	// Assert that the pets are not identical
	assert.NotEqual(t, d1, d2, "Two random snakes seem to be identical.")
}
