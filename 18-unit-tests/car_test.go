package car

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	a := Add(1,2)

	assert.Equal(t, 3, a)
}
