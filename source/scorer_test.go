package source

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTermFrequency(t *testing.T){
	assert.Equal(t,2,termFrequency("light","light in iron_city_brewing_co-ic light"))
	assert.Equal(t,1,termFrequency("city","light in iron city brewing_co-ic light"))
}

func TestFieldNormalize(t *testing.T) {
	assert.Equal(t,1/math.Sqrt(4),fieldNormalization("hi this is example"))
}


