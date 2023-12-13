package simulatorFuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomicClassCreation(t *testing.T) {
	assert := assert.New(t)
	correctMap := map[string]string{
		"f": "A",
		"g": "A",
		"h": "A",
		"i": "A",
	}

	class := NewClass("A", nil, []string{"f", "g", "h", "i"})

	for k, v := range class.Methods {
		assert.Equal(correctMap[k], v)
	}

}

func TestInheritanceCreation(t *testing.T) {
	assert := assert.New(t)
	correctMapA := map[string]string{
		"f": "A",
		"g": "A",
		"h": "A",
		"i": "A",
	}
	correctMapB := map[string]string{
		"f": "B",
		"g": "A",
		"h": "A",
		"i": "B",
	}
	correctMapC := map[string]string{
		"f": "B",
		"g": "C",
		"h": "A",
		"i": "C",
	}

	classA := NewClass("A", nil, []string{"f", "g", "h", "i"})
	classB := NewClass("B", &classA, []string{"f", "i"})
	classC := NewClass("C", &classB, []string{"g", "i"})

	for k, v := range classA.Methods {
		assert.Equal(correctMapA[k], v)
	}

	for k, v := range classB.Methods {
		assert.Equal(correctMapB[k], v)
	}

	for k, v := range classC.Methods {
		assert.Equal(correctMapC[k], v)
	}
}
